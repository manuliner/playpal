

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
    async signInOrUpWithEmail( email: UserPayloadInterface) {
      this.loading = true;
      
      try {
        const response = await fetch('https://api.example.com/auth/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(email),
        });

        const data = await response.json();

        if (response.ok) {
          this.authenticated = true;
          this.loggedInUser = data.user;
          useCookie('user', JSON.stringify(data.user));
        } else {
          console.error(data);
        }
      } catch (error) {
        console.error(error);
      }

      this.loading = false;
    },

    logUserOut() {

    },
  },
});