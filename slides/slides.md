---
title: AuthN & AuthZ architecture for microservices AND monoliths
css: css/index.css
watch: true
revealOptions:
    transition: 'linear'
---

### **AuthN** & **AuthZ** architecture for **microservices** and **monoliths**

---

## About us

----

<!-- .slide: class="about" -->

- ![@ncarlier](img/ncarlier.jpg)
  - **Nicolas Carlier**
  - <i class="fa fa-twitter"></i> [@ncarlier](https://twitter.com/ncarlier)
  - <i class="fa fa-github"></i> [github.com/ncarlier](https://github.com/ncarlier)
- ![@mmarie](img/mmarie.jpg) <!-- .element: class="desaturate" -->
  - **Maximilien Marie**
  - <i class="fa fa-github"></i> [github.com/akraxx](https://github.com/akraxx)

---

## Auth<small>enticatio</small>N

> I'm sorry, I don't know who you are!

## Auth<small>ori</small>Z<small>ation</small>


> I'm sorry Dave, I'm afraid I can't do that.

---

## What can we expect from an AuthN service?

----

### Functional Requirements

- Login form
- Self registration
- Email verification
- Account management
- Password recovery
- Two factor authentication (OTP)
- Social logins
- ...

----

### Non Functional Requirements

- Highly available
- Uses state of the art industry standards
- Interoperable
- Auditable
- Administrable
- Secure

----

### And what do we mean by **secure**?

- What about **encryption**?
- What about **hashing**?
- What about **password policy**?
- What about **storage**?
- What about **used libraries** and **patch management**?

---

## Should you built that on your **own?**

----

> "Why not? Is not that hard! It's just copy paste. I already did that sooooo
> many times!"

The crazy developer <!-- .element: class="signature" -->

----

> "Come on, it's just glue. Their is tons of libraries to help you."

Spring is coming <!-- .element: class="signature" -->

----

> "You should not! We have a mighty homebrew PCI framework to do that."

The special advisor that don't code anymore <!-- .element: class="signature" -->

---

## You are **brave**! Here some DIY helpers

- Java Authentication and Authorization Service ([JAAS](http://docs.oracle.com/javase/7/docs/technotes/guides/security/jaas/JAASRefGuide.html))
- [Spring Security](https://projects.spring.io/spring-security/)
- [Spring Cloud Security](https://cloud.spring.io/spring-cloud-security/)
- [Apache Shiro](https://shiro.apache.org/)

---

## And what **if**...

We delegate **ALL** of this?

---

## To **who**?

> "What a question! To the CAS!!!"

RSSI <!-- .element: class="signature" -->

Ok... but which one? Which version? Can I use my own federation provider? Is it
scalable? ...

---

![Keycloak](img/keycloak-logo.png)

---

## Keycloak

- Open Source project supported by RedHat
- Since ~2013
- Built on top of JBoss
- Hosted on [GitHub](https://github.com/keycloak/keycloak)
- Active community
- Constant and regular feature, bugfix and security releases
- Good & comprehensive [documentation](http://www.keycloak.org/documentation.html)
- Enterprise offer: [Red Hat Single Sign-On](https://access.redhat.com/documentation/en/red-hat-single-sign-on/)

---

## Features

Single-Sign-On/Out, Self-Registration, Forgot Password, Verify User/Email, TOTP,
Verification (Work-)Flows, Customer Attributes, Custom Federation Provider,
SPIs, Social Logins, Custom Themes, Account Management, Management Console,
Impersonation, and [more](https://keycloak.gitbooks.io/documentation/server_admin/topics/overview/features.html).

---

## Implements standards

| | | |
| --- |
| SAML 2.0 | OASIS | 2005
| OAuth 2.0 | RFC 6749 | 2012
| OpenID Connect 1.0 | OpenID Foundation | 2014
| JWT | RFC 7519 | 2015

----

### OAuth2

Authorization, *NOT* Authentication!

> The OAuth 2.0 authorization framework enables a 3rd-party application to obtain limited access to an HTTP service.

<small>[IETF](https://tools.ietf.org/html/rfc6749)</small>

----

### OIDC

OpenID Connect - *NOT* OpenID

- Authentication layer on top of OAuth 2.0
- Verify the identity of an end-user
- Obtain basic profile information about the end-user
- RESTful HTTP API, using JSON as data format
- Allows clients of all types (web-based, mobile, JavaScript)


[OpenID Foundation](http://openid.net/connect)

----

<!-- .slide: class="jwt" -->

### JWT

JSON Web Token - pronounced "jot"

It’s a JSON object encoded and signed to be used like a token.

- Encapsulate all the data that the backend needs for authorization
- Payload contains reserved, public, and private [claims](https://www.iana.org/assignments/jwt/jwt.xhtml)
- Trust based on a simple payload signature verification:
  - H34D3R <!-- .element: class="jwt-header" -->
  - .
  - P4YL04D <!-- .element: class="jwt-payload" -->
  - .
  - 516N47UR3 <!-- .element: class="jwt-signature" -->

<small>[IETF](https://tools.ietf.org/html/rfc7519)</small>

---

## Provides many adapters

- OpenID Connect:
  - Java (JBoss, WildFly, Tomcat, Jetty, Spring Security/Cloud, ...)
  - JavaScript
  - Node.js (Connect, Express, ...)
  - C# (OWIN)
  - And many generic implementations for Python, Android, iOS, Apache, Proxy ...
- SAML:
  - Java (JBoss, WildFly, Tomcat, Jetty)
  - Apache HTTP Server (mod_auth_mellon)
  - And some generic implementations

---

<!-- .slide: data-background="img/demo.jpg" -->

## DEMO TIME!

> I can see no reason for it to fail

---

<!-- .slide: data-background="img/qa.jpg", class="conclusion" -->

## Are you ready to delegate?

# Thanks!


