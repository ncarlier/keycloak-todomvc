// Fork to work with restdb.io NoSql backend
// Full spec-compliant TodoMVC with localStorage persistence
// and hash-based routing in ~246 effective lines of JavaScript.

var keycloak = Keycloak({ url: 'http://devbox/auth', realm: 'todomvc', clientId: 'todo-webapp' });

class RestApi {
  constructor (endpoint) {
    this.endpoint = endpoint;
  }

  /**
   * Build query string from object.
   * @param {object} query Query object
   * @return {string} Query string
   */
  _buildQueryString (query) {
    if (query) {
      const params = Object.keys(query).reduce((acc, key) => {
        if (query[key]) {
          acc.push(
            encodeURIComponent(key) + '=' + encodeURIComponent(query[key])
          )
        }
        return acc;
      }, [])
      return params.length ? '?' + params.join('&') : '';
    } else {
      return '';
    }
  }

  /**
   * Reslove full URL.
   * @param {string} url Base URL
   * @param {object} query Query object
   * @return {string} full URL
   */
  _resolveUrl (url, query) {
    return url + this._buildQueryString(query);
  }

  /**
   * Fetch method.
   * @param {string} url URL
   * @param {object} params Fetch parameters
   * @return {promise} Request result.
   */
  fetch (url, params) {
    params = Object.assign({
      method: 'get',
      headers: {
        Accept: 'application/json'
      },
      credentials: 'include'
    }, params);
    const {method, body, headers, query, credentials} = params;
    if (method === 'post' || method === 'put' || method === 'patch') {
      headers['Content-Type'] = 'application/json';
    }

    let authz = new Promise(function (resolve, reject) {
      keycloak.updateToken(30).success(() => {
        return resolve(keycloak.token);
      }, (err) => {
        // Fatal error from keycloak server
        console.error('Fatal error when updating the token', err);
        return reject(err);
      })
    });

    const _url = this._resolveUrl(url, query);
    // console.log('Fetcing URL:', _url)
    return authz.then((token) => {
      headers['Authorization'] = `Bearer ${token}`;
      return fetch(_url, {method, body, headers, credentials});
    })
    .then(response => {
      if (response.status === 204 || response.status === 205) {
        return Promise.resolve();
      } else if (response.status >= 200 && response.status < 300) {
        return response.json();
      } else {
        return response.json().then((err) => Promise.reject(err));
      }
    });
  }
  list (params) {
    const {page, size, sort} = params || {};
    return this.fetch(this.endpoint, {
      query: {page, size, sort}
    });
  }

  create (object) {
    return this.fetch(this.endpoint, {
      method: 'post',
      body: JSON.stringify(object)
    });
  }

  update (obj, update) {
    return this.fetch(obj._links.self.href, {
      method: 'put',
      body: JSON.stringify(update)
    });
  }

  remove (obj) {
    return this.fetch(obj._links.self.href, {
      method: 'delete'
    });
  }
}

const todoApi = new RestApi('/api/todos');

var todoStorage = {
  fetch: function() {
    return todoApi.list().then(data => {
      var todos = data._embedded && data._embedded.todos ? data._embedded.todos : [];
      todos.forEach(function(todo, index) {
        todo.id = index;
      });
      return Promise.resolve(todos);
    })
  },
  save: function(todo) {
    console.log("saving", todo);
    return todoApi.create(todo).then(result => {
      console.log("Created", result);
      return Promise.resolve(result);
    });
  },
  update: function(todo) {
    console.log("updating", todo);
    const {title, completed} = todo
    return todoApi.update(todo, {title, completed}).then(result => {
      console.log("Updated", result);
      return Promise.resolve(result);
    });
  },
  updateAll: function(todos) {
    console.log("updating", todos);
    var saveOne = function(todo, callback) {
      const {title, completed} = todo
      todoApi.update(todo, {title, completed}).then(result => {
        console.log("Updated", result);
        callback(null, result);
      })
    }

    var funcs = [];
    todos.forEach(function(todo) {
      funcs.push(async.apply(saveOne, todo));
    });
    async.parallel(funcs, function(error, result) {
      console.log("updateAll", error, result)
    });
  },
  delete: function(todo) {
    console.log("deleting", todo);
    return todoApi.remove(todo).then(result => {
      console.log("Deleted", result);
      return Promise.resolve(todo);
    });
  }
}

// visibility filters
var filters = {
  all: function(todos) {
    return todos
  },
  active: function(todos) {
    return todos.filter(function(todo) {
      return !todo.completed
    })
  },
  completed: function(todos) {
    return todos.filter(function(todo) {
      return todo.completed
    })
  }
}

// app Vue instance
var app = new Vue({
  // app initial state
data: {
  todos: [],
  newTodo: '',
  editedTodo: null,
  visibility: 'all'
},

  // watch todos change for localStorage persistence
watch: {
  todos: {
    handler: function(todos) {
      console.log("Something happend to ", todos.length)
        //todoStorage.save(todos)
    },
    deep: true
  }
},

  // computed properties
  // http://vuejs.org/guide/computed.html
computed: {
  filteredTodos: function() {
    return filters[this.visibility](this.todos)
  },
  remaining: function() {
    return filters.active(this.todos).length
  },
  allDone: {
    get: function() {
      return this.remaining === 0
    },
    set: function(value) {
      this.todos.forEach(function(todo) {
        todo.completed = value
      });
      todoStorage.updateAll(this.todos);
    }
  }
},

filters: {
  pluralize: function(n) {
    return n === 1 ? 'item' : 'items'
  }
},
created: function() {
  keycloak.init({ onLoad: 'login-required' })
    .success((authenticated) => {
      console.log(authenticated ? 'authenticated' : 'not authenticated');
      console.log("Ready")
      this.getTodoFromDb();
    }).error(err => console.error(err));
},
  // methods that implement data logic.
  // note there's no DOM manipulation here at all.
methods: {

  getTodoFromDb: function() {
    console.log("Getting todo list...");
    todoStorage.fetch().then((tododata) => {
      console.log("Get data ", tododata);
      this.todos = tododata;
    });
  },
  addTodo: function() {
    var value = this.newTodo && this.newTodo.trim()
    if (!value) {
      return
    }
    var newtodo = {
      title: value,
      completed: false
    };
    todoStorage.save(newtodo).then(t => {
      this.todos.push(t);
      this.newTodo = '';
    });
  },

  removeTodo: function(todo) {
    todoStorage.delete(todo).then(t => {
      this.todos.splice(this.todos.indexOf(todo), 1)
    });
  },

  setState: function(todo) {
    todo.completed = !todo.completed;
    todoStorage.update(todo);
  },

  editTodo: function(todo) {
    // todoStorage.update(todo);
    this.beforeEditCache = todo.title
    this.editedTodo = todo
  },

  doneEdit: function(todo) {
    console.log("doneEdit", todo)
    if (!this.editedTodo) {
      return
    }
    this.editedTodo = null
    todo.title = todo.title.trim()
    if (!todo.title) {
      this.removeTodo(todo)
    } else {
      todoStorage.update(todo);
    }
  },

  cancelEdit: function(todo) {
    this.editedTodo = null
    todo.title = this.beforeEditCache
  },

  removeCompleted: function() {
    this.todos = filters.active(this.todos)
  }
},

  // a custom directive to wait for the DOM to be updated
  // before focusing on the input field.
  // http://vuejs.org/guide/custom-directive.html
directives: {
  'todo-focus': function(el, value) {
    if (value) {
      el.focus()
    }
  }
}
})

// handle routing
function onHashChange() {
  var visibility = window.location.hash.replace(/#\/?/, '')
  if (filters[visibility]) {
    app.visibility = visibility
  } else {
    // window.location.hash = ''
    app.visibility = 'all'
  }
}

window.addEventListener('hashchange', onHashChange)
onHashChange()

app.$mount('.todoapp');


