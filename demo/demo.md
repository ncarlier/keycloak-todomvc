Demo Time !

- https://jsonformatter.curiousconcept.com/
- https://jwt.io/

Main goal of this demo is to show you how it's simple to fully delegate the AuthN and AuthZ part to keycloak.
We choosed something very simple, it's a todo webapp where you can add some tasks to complete.
There are 3 parts, the backend in Java (built with Sprint Boot), the Webapp in VueJS and a commandline in Golang.

## Quickstart

Let's start everything :

```bash
make deploy
```

This command will build every part of the application : back, webapp and cli images.
As you can see it's also building and launching keycloak, we will use it just after.

Let's see our wondeful application !!
```bash
http://devbox/app/

curl -s http://devbox/api/todos | jq .

docker run --rm -ti --add-host="devbox:172.17.0.1" ncarlier/keycloak-todomvc-cli
todomvc ls
```

## Secure the back

It's great but... Not very secured!
First thing to do : secure the back

```bash
> todomvc-api/conf.env
SECURITY_BASIC_ENABLED=true
MANAGEMENT_SECURITY_ENABLED=true
```

```bash
make deploy
```

Launch the curl again :
```bash
curl -s http://devbox/api/todos | jq .
```
Our API is now secured and you have to be logged in to make a REST request
The only thing we did in our app is to extend the security adapter provided by keycloak !
There is a lot of available adapters, but of course, it's possible to do it manually because keycloak is based on OIDC standard.

Go to the app now :
http://devbox/app/
Open the browser console to inspect network requests, you will get the same 401 error as before.

## Integrate keycloak on our webapp

```bash
> todomvc-app/index.html
<script src="http://devbox/auth/js/keycloak.js"></script> # Uncomment line 50
```

```bash
make deploy
```

Go back to http://devbox/app/
Make sure you still have you browser console active to spy network requests (XHR).
You are automatically redirected to keycloak ! 
Enter todo/todo, and put a new password for the todo user.
Click on login, you are now logged on the todo app!

In the console you will see the 2 following requests :
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

As you can see you have 2 tokens : access and refresh.
The access token is the one used to call the API.
Refresh token is used to call the keycloak token endpoint to generate a new access token. That's why the refresh token has a longer lifetime.

The second request is used to get the todos :
GET http://devbox/api/todos

Look at the headers of the request, you will see that the authorization header has been automatically added. 

If we inspect this token in http://jwt.io, we can see that we have a lot of data in the payload, some of them :
Issuer (iss) : App who generated the token
Aud (audience) : Who the token has been generated for
exp (Expiration) : When the token will expire, by default 300s (5m)

## CLI

It's time to see how we can integrate keycloak and our the command line

```bash
todomvc login -u todo -p todo
```
Unable to get offline token : 400 Bad Request : {"error":"unauthorized_client","error_description":"UNKNOWN_CLIENT: Client was not identified by any client authenticator"}

Error is very explicit, there is not client in keycloak for the command line...
Go to http://devbox/auth/, click on 'Administration Console' with admin/admin as credentials.

Go to clients > create
Client id : todo-cli
Click on save

Go back to the container an type :

```bash
todomvc login -u todo -p todo
todomvc ls
```

In fact, the login step just store a file "creds.json" containing the response from keycloak :

```bash
cat /root/.todomvc/creds.json | jq .
```

The difference between this response and the one for the webapp is the expiration date of the refresh token.
In our example, for every http call we will use this refresh token to get a brand new access token!

## Authorization

We saw the authentication part, but not so much on authorization ! Let's see a little example of how we can use keycloak roles in our applications.

Go to the page http://devbox/api/
Try to explore /api/todos. The response should be 200 with the todo list.
No explore /api/env... The response should be :

```json
{
  "timestamp": 1495543258125,
  "status": 403,
  "error": "Forbidden",
  "message": "Access is denied. User must have one of the these roles: ACTUATOR",
  "path": "/api/env"
}
```

Go to http://devbox/auth/, click on 'Administration Console' with admin/admin as credentials.

Create the role 'ROLE_ACTUATOR'
Add the role to the user 'todo'

Open a new window in private navigation (or clear storage from devbox)
Go to : http://devbox/api/ and explore /api/env.
You should see all system properties !

Open http://devbox/app/ and open the browser console to capture the access token. Copy it and paste it to jwt.io. 
The role actuator is now in the list of roles !