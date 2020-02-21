import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import About from "../views/About.vue";
import store from "../store";
import firebase from "firebase/app";
import "firebase/auth";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/about",
    name: "About",
    component: About,
    meta: { requiresAuth: true }
  },
  {
    path: "/signout",
    name: "SignOut",
    beforeEnter: (_to, _from, next) => {
      firebase.auth().signOut();
      next("/");
    }
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
    props: route => ({ redirect: route.query.redirect }),
    beforeEnter: (_to, _from, next) => {
      console.log("store state", store.getters.currentUser);
      if (store.getters.currentUser) {
        next("/");
      } else {
        next();
      }
    }
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, _from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.currentUser) {
      console.log("continuing to", to);
      next();
    } else {
      next({
        path: "/login",
        query: { redirect: to.fullPath }
      });
    }
  } else {
    next();
  }
});

export default router;
