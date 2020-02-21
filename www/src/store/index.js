import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export const mutations = {
  SAVE_USER: (state, user) => {
    state.auth.initialized = true;
    console.log("Initialzing auth");
    if (user) {
      state.auth.authenticated = true;
      state.auth.user = user;
    } else {
      state.auth.authenticated = true;
      state.auth.user = user;
      state.auth.token = null;
    }
  },
  SAVE_TOKEN: (state, token) => {
    state.auth.token = token;
  }
};

export default new Vuex.Store({
  state: {
    auth: {
      initialized: false,
      authenticated: false,
      user: null,
      token: null
    }
  },
  mutations: mutations,
  getters: {
    currentUser: state => {
      return state.auth.user;
    },
    authInitialzied: state => {
      return state.auth.initialized;
    }
  },
  actions: {
    saveUser: ({ commit }, user) => {
      commit("SAVE_USER", user);
      if (user) {
        user
          .getIdToken(true)
          .then(token => {
            commit("SAVE_TOKEN", token);
          })
          .catch(console.log);
      }
    },
    saveToken: ({ commit }, token) => {
      commit("SAVE_TOKEN", token);
    }
  },
  modules: {}
});
