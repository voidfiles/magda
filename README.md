# Magda

A platform for emmersing one self in the power of beauty, chance, and juxtaposition.

# What?

There are so many beautiful things in this world. Art, a quote, a picture from your loved ones. As we pass through life we are enriched when we encounter these things. But, so often they are hidden from our eye, or they are arranged so staticially that we miss out on all the new thigns being produced everyday.

Wouldn't it be great if you could ermerse yourself in these touch points. For them to slowly evolve around you everyday. From your house, to your phone, to the screens around you.

Magda is my attempt to build a platform for that experience.

# How?

Great question. I don't know exactly.

## Authentication 

User accounts are managed via firebase. But, firebase has relativley minimal user properties. So things like a user roles will need to be modeled in the datastore.

Admin auth will be done based on firebase auth SDK

*How to get your service credentials*

1. Vist https://firebase.google.com/docs/admin/setup/
2. Download your credential to `./service_account_key.json` this is ignored by .gitignore

*Firebase*

Right now this is going to use [firebase authentication](https://firebase.google.com/docs/auth/).

I don't want to key off this to much. The application should be built with authenticaiton and authorization, but I want to keep it simple to begin with. Users should be able to see almost everythign except for like the user table. And only admins should be able to write anything, except maybe the user table. A user should be able to create their own user account.

### JWTs 

Firebase can create a JWT, that will be sent to the API as the authorization method.

### User Model

This will be stored in firebase, like /users/{uid}

Right now it will contain a role field.

### Roles

*Anonymous*

They can be viewered

*User*

They can maybe someday actually create things

*Admin*

Can modify almost anything in the system


## Contributors

**Prerequists**: Right now, you need to have working go environment.

Then you can run: `make setup` to install required tools.

To regenerate the support tools run: `make generate`

You can run the webapp via `make run`

## The Plan

- [X] Authorization of go Service
    - [X] Setup a middleware
    - [X] Understand how logging is going to work accross middlewares
        - [X] This is going to require stacking middlewares togher
- [] Setup authorization
    - [] Create a user model
    - [] Only allow admins right now to make any changes
    - [] Require authentication for all reads, use anonymous for folks who haven't signed up
- [] Create the data model in GraphQL
    - [x] Some initial model work
    - [] Entrys
    - [] Entitys
    - [] Collections
- [] Create an API
- [] Create mockups of some of the screens
- [] Create a storybook for the components
- [] Figure out the Authentication story
    - [] Roles
        - [] Viewer (possibly anonymous), can look at collections
        - [] User (Can add things to the system that belong to them)
        - [] Admin (Can do anything )


# Examples of data

## Art

Art is often found browsing the web, or while reading feeds. Based on some signal from a user, the page can be broken down into a set of the following bits of data. Note there is very little data here. It will require human curation at a later date to transform this raw entry into a fully detailed one.

Raw

Shape

```
source:
  page: Website
  image: SourceFile
```

```yaml
source:
  page:
    url: https://la-beaute--de-pandore.tumblr.com/post/190039691507/loretta-lux
    title: "La Beauté de Pandore — Loretta Lux"
    description: ""
  image:
    found_url: https://66.media.tumblr.com/e33bd2baf3bd7fcc98c8044c951315b3/1052e56105e7b734-99/s1280x1920/4d8bee72d5b3ef070d600bd39187bdf00e3eaab7.jpg
    dimensions:
      width: 800
      height: 585
```

Full

```yaml
source:
  page:
    url: https://la-beaute--de-pandore.tumblr.com/post/190039691507/loretta-lux
    title: "La Beauté de Pandore — Loretta Lux"
    description: ""
  site:
    url: https://la-beaute--de-pandore.tumblr.com/
    title: "La Beauté de Pandore"
    description: ""
item:
  type: "image"
  found_url: https://66.media.tumblr.com/e33bd2baf3bd7fcc98c8044c951315b3/1052e56105e7b734-99/s1280x1920/4d8bee72d5b3ef070d600bd39187bdf00e3eaab7.jpg
  dimensions:
    width: 800
    height: 585
creators:
  - name: "Loretta Lux"
    image: ""
    description: ""
    urls:
      - url: http://lorettalux.de/
        type: homepage
      - url: http://www.artnet.com/artists/loretta-lux/
        type: artnet
      - url: https://en.wikipedia.org/wiki/Loretta_Lux
        type: wikipedia
      - url: https://www.artsy.net/artist/loretta-lux
        type: artsy
```

## Quote

Quotes unlike art would require detailed entry in the first place. It's difficult to automatically curate a quote.

```yaml
source:
  page:
    url: https://www.cs.yale.edu/homes/perlis-alan/quotes.html
    title: "Perlisisms - "Epigrams in Programming" by Alan J. Perlis"
    description: ""
item:
  type: "quote"
  content:
    text: "One man's constant is another man's variable."
creators:
  - name: "Alan J. Perlis"
    description: ""
    urls:
      - url: https://en.wikipedia.org/wiki/Alan_Perlis
        type: wikipedia
```