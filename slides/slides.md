---
title: AuthN & AuthZ architecture for microservices AND monoliths
css: css/index.css
watch: true
revealOptions:
    transition: 'linear'
---

### AuthN & AuthZ architecture for microservices AND monoliths

---

## About us

----

<!-- .slide: class="about" -->

- ![@ncarlier](img/ncarlier.jpg)
  - Nicolas Carlier
  - <i class="fa fa-twitter"></i> [@ncarlier](https://twitter.com/ncarlier)
  - <i class="fa fa-github"></i> [github.com/ncarlier](https://github.com/ncarlier)
- ![@mmarie](img/mmarie.jpg)
  - Maximilien Marie

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
- Account management
- Password recovery
- Email verification
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

## Should we built that on your own?

----

> "Why not? Is not that hard! It's just copy paste. I already did that sooooo
> many times!"

*The crazy developer*

----

> "Come on, it's just glue. Their is tons of libraries to help you."

*Spring is comming*

----

> "You should not! We have a mighty homebrew PCI framework to do that."

*The special advisor that don't code anymore*

---

## You are brave! Here some DIY helpers

- JAAS
- Spring Security
- Spring Cloud Security
- Apache Shiro

---

## And what if...

We delegate ALL of this?

---

## To who?

> RSSI: What a question: To the CAS!

> You: Which one? Which version? Can I use my own federation provider? Is it
> scalable? ...

> RSSI: ... Did I already told you about PCI --the-- SS

---

![Keycloak](img/keycloak-logo.png)

---

## Keycloak

- JBoss
- since ~2013
- Open Source Software
- hosted at GitHub
- active Community
- constant and regular feature- and bugfix-releases
- good & comprehensive documentation

---

## Features

- Single-Sign-On/Out
- Self-Registration
- Forgot Password
- Verify User/Email
- TOTP
- Verification (Work-)Flows
- Customer Attributes
- Custom Federation Provider
- SPIs
- Social Logins
- Custom Themes
- Account Management
- Management Console
- Impersonation
- and more_

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

Authentication layer on top of OAuth 2.0

- verify the identity of an end-user
- obtain basic profile information about the end-user
- RESTful HTTP API, using JSON as data format
- allows clients of all types  
  (web-based, mobile, JavaScript)


[OpenID Foundation](http://openid.net/connect)

----

### JWT

TODO

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

---

<!-- .slide: data-background="img/qa.jpg" -->

## Thanks! <!-- .element: class="black" -->


