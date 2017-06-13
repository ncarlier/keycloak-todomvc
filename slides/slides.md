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

Note:

In order to have the same vocabulary, here a quick definition:
AuthN is about verify and confirm the identity claimed by a subject.
AutnZ is about verify access rights of a known user to resources.

---

## What can we expect from an AuthN service?

Note:

So if we want to build or use an AuthN service, what can we expect?

----

### Functional Requirements

- Login form <!-- .element: class="fragment" -->
- Self registration <!-- .element: class="fragment" -->
- Email verification <!-- .element: class="fragment" -->
- Account management <!-- .element: class="fragment" -->
- Password recovery <!-- .element: class="fragment" -->
- Two factor authentication (OTP) <!-- .element: class="fragment" -->
- Social logins <!-- .element: class="fragment" -->
- ... <!-- .element: class="fragment" -->

Note:

First regarding the functional requirements:
Obviously we want to expose a login form or a access point to this service. You
may want to customize the look and feel regarding your customer or your needs.
You may want to allow your users the ability to self register. This feature
often needs other mechanisms like e-mail verification.
You may want to provide account management for your customer and also a self
management for your users.
If your authentication is password based, you certainly want a lost password
feature.
Again if you are password based, you way want to harden the security by adding
other mechanisms like One Time Password.
Something quite usual nowadays, you may want to delegate this authentication to
a social provider (such as Twitter, Facebook, Google, ...).

----

### Non Functional Requirements

- Highly available <!-- .element: class="fragment" -->
- Uses state of the art industry standards <!-- .element: class="fragment" -->
- Interoperable <!-- .element: class="fragment" -->
- Auditable <!-- .element: class="fragment" -->
- Administrable <!-- .element: class="fragment" -->
- Secure <!-- .element: class="fragment" -->

Note:

Now, regarding the non-functional requirements:

Such a system is critic. You want high availability for this service.
This service should also use the state of the art industry standards.
You way want some interoperability with other services. Maybe the authentication
service of your customer or maybe your customer user referential.
Any service related to security aspects must be auditable. The ability to trace
every access, every actions. Especially if you are in a PCI context.
Such a complex service should provide administrations features and user
interfaces.
And last but not the least! The service have to be secure!

----

### And what we mean by **secure**?

- <!-- .element: class="fragment" --> What about **encryption**?
- <!-- .element: class="fragment" --> What about **hashing**?
- <!-- .element: class="fragment" --> What about **password policy**?
- <!-- .element: class="fragment" --> What about **storage**?
- <!-- .element: class="fragment" --> What about **used libraries** and **patch management**?

Note:

And what do we meant by secure?
Encryption: What are the algorithms used by your service to cypher channels and
data? Are these algorithms still secure?
Hashing: Same question. SHA1 is considerate as deprecated.
What is your password policy? Can you tolerate a weak password?
How do you store sensitives informations like personal data and credentials?
Which libraries are you using? Is there any published security issue in it?

You should be able to answer those questions.

---

## Should you built that on your **own?**

Note:

So considering all these requirements. Should we built that by ourself?

----

> "Why not? Is not that hard! It's just copy paste. I already did that sooooo
> many times!"

The crazy developer <!-- .element: class="signature" -->

----

> "Come on, it's just glue. There is tons of libraries to help you."

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

Note:

Ok you are brave and this is great. In this case here some DIY helpers to help
you in that quest.

---

## And what **if**...

We delegate **ALL** of this?

Note:

But if like me you are not so brave, what if we delegate all of this to someone
else?

---

## To **who**?

> "What a question! To the CAS!!!"

RSSI <!-- .element: class="signature" -->

Ok... but which one? Which version? Can I use my own federation provider? Is it
scalable? ...
<!-- .element: class="fragment" -->

---

![Keycloak](img/keycloak-logo.png)

Note:

Let us present you a more simple and relevant solution: Keycloak

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

Note:

Let's have a very quick glance at those standards...

----

### OAuth2

Authorization, *NOT* Authentication!

> The OAuth 2.0 authorization framework enables a 3rd-party application to obtain limited access to an HTTP service.

<small>[IETF](https://tools.ietf.org/html/rfc6749)</small>

----

### OIDC

OpenID Connect - *NOT* OpenID

- Authentication layer on top of OAuth 2.0 <!-- .element: class="fragment" -->
- Verify the identity of an end-user <!-- .element: class="fragment" -->
- Obtain basic profile information about the end-user <!-- .element: class="fragment" -->
- Standardized RESTful API, using JSON as data format <!-- .element: class="fragment" -->
- Support many types of clients (web-based, mobile, JavaScript) <!-- .element: class="fragment" -->


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

Note:

Keycloak comes with a lot of adapters in order to simplify its integration.

---

<!-- .slide: data-background="img/demo.jpg" -->

## DEMO TIME!

> I can see no reason for it to fail

Note:

It is over for the theory. Demo time!

---

<!-- .slide: data-background="img/qa.jpg", class="conclusion" -->

## Are you ready to delegate?

# Thanks!


