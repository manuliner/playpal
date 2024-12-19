

import { defineStore } from 'pinia';

interface UserPayloadInterface {
  username: string;
  password: string;
}

export interface User {
  id: number;
  username: string;
  email: string;
  firstName: string;
  lastName: string;
  gender: string;
  image: string;
  accessToken: string;
  refreshToken: string;
}


export const useAuthStore = defineStore('auth', {
  state: () => ({
    authenticated: false,
    loading: false,
    loggedInUser: {}
  }),
  actions: {
    initializeAuth() {

      const user = useCookie('user');

      if (user.value) {
        this.authenticated = true;
        this.loggedInUser = JSON.parse(user.value || '{}');
      } else {
        this.authenticated = false;
        this.loggedInUser = {};
      }
    },
    async authenticateUser({ username, password }: UserPayloadInterface) {
    },

    logUserOut() {

    },
  },
});