# DEMO

Main goal of this demo is to show you how it is simple to delegate the
AuthN and AuthZ part of your application to [Keycloak](http://www.keycloak.org/).

We chose a very simple sample app to secure: It's the well known
[Todo MVC App](http://todomvc.com/).

The application consists of:

- a **backend** built in Java with Sprint Boot
- a **webapp** built in Javascript with VueJS
- and a **CLI** built in Golang

## Quickstart

Let's start everything :

```bash
make build deploy
```

This command will build and deploy every services of the application thanks to
Docker compose.

During the deployment, you can see the configuration phase of Keycloak.

Let's see our wonderful application:

With the browser: [http://devbox/app/](http://devbox/app)

With the shell: `curl -s http://devbox/api/todos | jq .`

With the CLI: `make cli`

Everything should work but without any authentication. Let's secure this!

## Secure the back

We have to activate Spring Security:

```bash
vi todomvc-api/conf.env
# Uncomment the following lines:
# SECURITY_BASIC_ENABLED=true
# MANAGEMENT_SECURITY_ENABLED=true
```

Then redeploy the API:

```bash
make deploy
```

Launch cURL again:

```bash
curl -s http://devbox/api/todos | jq .
```

The API is now secured and you have to be logged in to make a REST request.
The only thing we did in our app is to extend the security adapter provided by
Keycloak. There is many other adapters. If you want you can also manually
validate the token using a very simple Servlet Filter (thanks to OIDC and JWT).

If we try to test our [Webapp](http://devbox/app/) we will get the same `401`
error. Let's fix that.

## Integrate Keycloak into our Webapp

We have to import the Javascript adapter:

```bash
vi todomvc-app/index.html
# Uncomment the following line:
# <script src="http://devbox/auth/js/keycloak.js"></script>
```

Then redeploy the APP:

```bash
make build service=app
make deploy
```

Go back to the [Webapp](http://devbox/app/)

Make sure you still have you browser console active to spy network requests.
As you can see you are now automatically redirected to Keycloak login page.

Enter todo/todo, and update the todo user password as asked.

Click on login and you should be logged in the Todo App.

What happened? In the console you should see following requests:

```
POST http://devbox/auth/realms/todomvc/protocol/openid-connect/token

Response (use https://jsonformatter.curiousconcept.com/ to format the response) :
{
   "access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJFa1JxNUN...",
   "expires_in":300,
   "refresh_expires_in":1800,
   "refresh_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJFa1JxNUN...",
   "token_type":"bearer",
   "id_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJFa1JxNUN6TDI...",
   "not-before-policy":0,
   "session_state":"f2c44dc9-b560-43e7-9b67-b0a742cef625"
}
```

As you can see you have some tokens:
- The access token
- The Refresh token
- And the ID token

The access token is the one used to call the API. This token has a short TTL.
Once expired, the refresh token is used to call the Keycloak token endpoint in
order to get a new access token. This token has a longer TTl and can be used
only once.
The ID token contains informations regarding the user.

You can decode those tokens using [http://jwt.io](http://jwt.io). You will see a
lot of attributes in the payload like:

- sub (subject): Who the token has been generated for
- exp (expiration): When the token will expire (by default 300s)
- ...

The second request is used to retrieve the todos:

```
GET http://devbox/api/todos
```

Look at the headers of the request, you will see that the authorization header
has been automatically added.

## CLI

It's time to see how we can integrate Keycloak with our CLI.
Let's try to login.


```bash
make cli
todomvc login -u todo -p todo
```

We should be unable to get offline token. Keycloak response:

```
400 Bad Request
{
  "error":"unauthorized_client",
  "error_description":"UNKNOWN_CLIENT: Client was not identified by any client authenticator"
}
```

The error is very explicit: there is not configured client in Keycloak for the
CLI.
You have to declare a client for the CLI. In order to do that, go to the
[Administration Console](http://devbox/auth) with admin/admin as credentials.

Create a new client with the following client id: `todo-cli`.

Save and go back to the CLI:

```bash
todomvc login -u todo -p todo
todomvc ls
```

Behind the scene the login command store a file `creds.json` containing the
response from Keycloak:

```bash
cat /root/.todomvc/creds.json | jq .
```

In this specific implementation the Access Token will not be used. For every
call, the refresh token will be used to retrieve a brand new Access Token.
The refresh token has no TTL. It's an offline token.

You can manage this offline token in the admin console:
click on *users* search for the *todo* user, click on consent tab and revoke
offline token.

Now if you try to use again the CLI:

```bash
todomvc ls
```

You will get a bad request and you have to login again.

## Quick glance at Authorization

We saw the authentication part, but not so much on authorization.
Let's see a little example of how we can use Keycloak roles in our applications.

Open the [API Browser](http://devbox/api/).
Try to fetch `/api/todos`. The response should be `200` with the todo list.

Now, fetch `/api/env`. The response should be :

```json
{
  "timestamp": 1495543258125,
  "status": 403,
  "error": "Forbidden",
  "message": "Access is denied. User must have one of the these roles: ACTUATOR",
  "path": "/api/env"
}
```

This endpoint is a Spring Actuator endpoint. And Spring protect this endpoint by
requiring the ACTUATOR role. We have to give this role to our user.

Open again the [Admin Console](http://devbox/auth/), and create the role:
`ROLE_ACTUATOR` (The 'ROLE_' prefix is required by Spring Security).
Then affect this role to the 'todo' user.

Open a new window in private navigation (or clear cookies and local storage)
Open the [API Browser](http://devbox/api/) and fetch again `/api/env`.

You should see all system properties.

If you extract and decode the Access Token thanks to the development console.
You will see this role inside the payload.


Thanks.
