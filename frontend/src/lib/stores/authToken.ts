import { writable } from 'svelte/store';

function createAuthTokenStore() {
  const { subscribe, set } = writable<string | null>(null);

  return {
    subscribe,
    set: (token: string | null) => {
      set(token);
      if (typeof window !== 'undefined') {
        if (token) {
          localStorage.setItem('authToken', token);
        } else {
          localStorage.removeItem('authToken');
        }
      }
    },
    initialize: () => {
      if (typeof window !== 'undefined') {
        const token = localStorage.getItem('authToken');
        if (token) {
          set(token);
        }
      }
    },
    clear: () => {
      set(null);
      if (typeof window !== 'undefined') {
        localStorage.removeItem('authToken');
      }
    },
  };
}

export const authToken = createAuthTokenStore();
