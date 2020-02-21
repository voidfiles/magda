# Magda

A platform for emmersing one self in the power of beauty, chance, and juxtaposition.

# What?

There are so many beautiful things in this world. Art, a quote, a picture from your loved ones. As we pass through life we are enriched when we encounter these things. But, so often they are hidden from our eye, or they are arranged so staticially that we miss out on all the new thigns being produced everyday.

Wouldn't it be great if you could emerser you self in those touch points. For them to slowly evolve around you everyday. From your house, to your phone, to the screens around you.

Magda is my attempt to build a platform for that experience.

# How?

Great question. I don't know exactly.

## Authentication 

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

- [] Create the data model in GraphQL
    - [] Entrys
    - [] Entitys
    - [] Collections
- [] Setup authorization
    - [] Create a user model
    - [] Only allow admins right now to make any changes
    - [] Require authentication for all reads, use anonymous for folks who haven't signed up
- [] Create an API
- [] Create mockups of some of the screens
- [] Create a storybook for the components
- [] Figure out the Authentication story
    - [] Roles
        - [] Viewer (possibly anonymous), can look at collections
        - [] User (Can add things to the system that belong to them)
        - [] Admin (Can do anything )
