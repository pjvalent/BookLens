<!-- src/lib/components/SignupForm.svelte -->
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { goto } from '$app/navigation'
  
    let firstName: string = '';
    let lastName: string = '';
    let email: string = '';
    let password: string = '';
  
    export let error: string | null = null;
  
    const dispatch = createEventDispatcher();
  
    function handleSubmit(event: Event) {
      event.preventDefault();
      dispatch('submit', { firstName, lastName, email, password });
    }

    function goToLogin() {
      goto('/login');
    }
  </script>
  
  <style>
    .signup-form-container {
      background-color: #ffffff;
      color: #1b4d3e;
      padding: 2rem;
      border-radius: 8px;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      width: 100%;
      max-width: 400px;
    }
  
    h1 {
      text-align: center;
      margin-bottom: 1.5rem;
    }
  
    form div {
      margin-bottom: 1rem;
    }
  
    label {
      display: block;
      margin-bottom: 0.5rem;
      font-weight: bold;
    }
  
    input {
      width: 100%;
      padding: 0.5rem;
      border: 1px solid #ccc;
      border-radius: 4px;
    }
  
    input:focus {
      outline: none;
      border-color: #1b4d3e;
    }
  
    input:hover {
      border-color: #1b4d3e;
    }
  
    button {
      width: 100%;
      padding: 0.75rem;
      background-color: #1b4d3e;
      color: #ffffff;
      border: none;
      border-radius: 4px;
      font-size: 1rem;
      cursor: pointer;
    }
  
    button:hover {
      background-color: #166a52;
    }
  
    button:focus,
    button:active {
      outline: none;
      background-color: #145a45;
    }
  
    .error-message {
      color: #ff4d4d;
      margin-bottom: 1rem;
      text-align: center;
    }

    .login-button {
    width: auto;
    padding: 0.5rem 1rem;
    background-color: #1b4d3e;
    color: #ffffff;
    border: none;
    border-radius: 4px;
    font-size: 0.875rem;
    cursor: pointer;
    display: block;
    margin: 1rem auto 0;
    text-align: center;
  }
  
  .login-button:hover {
    background-color: #166a52;
  }

  </style>
  
  <div class="signup-form-container">
    <h1>Sign Up</h1>
  
    {#if error}
      <p class="error-message">{error}</p>
    {/if}
  
    <form on:submit={handleSubmit}>
      <div>
        <label for="firstName">First Name:</label>
        <input
          id="firstName"
          name="firstName"
          type="text"
          bind:value={firstName}
          required
        />
      </div>
  
      <div>
        <label for="lastName">Last Name:</label>
        <input
          id="lastName"
          name="lastName"
          type="text"
          bind:value={lastName}
          required
        />
      </div>
  
      <div>
        <label for="email">Email:</label>
        <input
          id="email"
          name="email"
          type="email"
          bind:value={email}
          required
        />
      </div>

      <div>
        <label for="password">Password:</label>
        <input
          id="password"
          name="password"
          type="password"
          bind:value={password}
          required
          />
      </div>
  
      <button type="submit">Sign Up</button>
      <button type="button" class="login-button" on:click={goToLogin}>Already have an account? Log In</button>
    </form>
  </div>
  