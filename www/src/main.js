import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import firebase from "firebase/app";

Vue.config.productionTip = false;

const firebaseConfig = {
  apiKey: "AIzaSyCqALgbt_4i2E4qJun7LDIo0VVGqbne88o",
  authDomain: "magda-4f6b9.firebaseapp.com",
  databaseURL: "https://magda-4f6b9.firebaseio.com",
  projectId: "magda-4f6b9",
  storageBucket: "magda-4f6b9.appspot.com",
  messagingSenderId: "704879928566",
  appId: "1:704879928566:web:6edc70c8d13b01ee55b8d7",
  measurementId: "G-YTDSJRVKFP"
};

firebase.initializeApp(firebaseConfig);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
