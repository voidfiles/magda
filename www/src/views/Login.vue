<template>
  <section id="firebaseui-auth-container"></section>
</template>

<script>
import firebase from "firebase/app";
import * as firebaseui from "firebaseui";
import "firebaseui/dist/firebaseui.css";

export default {
  name: "Login",
  props: ["redirect"],
  mounted: function () {
    let ui = firebaseui.auth.AuthUI.getInstance();
    if (!ui) {
      ui = new firebaseui.auth.AuthUI(firebase.auth());
    }
    let vm = this;
    var uiConfig = {
      signInSuccessUrl: this.redirect,
      credentialHelper: firebaseui.auth.CredentialHelper.NONE,
      signInOptions: [
        firebase.auth.EmailAuthProvider.PROVIDER_ID,
        firebaseui.auth.AnonymousAuthProvider.PROVIDER_ID
      ]
    };

    vm.$nextTick(() => {
      ui.start("#firebaseui-auth-container", uiConfig);
    });
  }
};
</script>
