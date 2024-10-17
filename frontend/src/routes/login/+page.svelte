<script lang="ts">
  import SignupForm from '$lib/components/LoginForm.svelte';
  import { goto } from '$app/navigation';
  import { authToken } from '$lib/stores/authToken';

  let error: string | null = null;

  async function handleLogin(event: CustomEvent) {
    const { email, password } = event.detail;
    error = null;

    try {
      const response = await fetch('http://localhost:8080/v1/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: email,
          password: password,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to log in');
      }

      // Parse the response to get the token
      const data = await response.json();
      const token = data.token;

      // Store the token
      authToken.set(token);

  
      goto('/welcome');
    } catch (err) {
      error = (err as Error).message;
    }
  }
</script>





<style>
  .login-page-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: 1rem;
  }
</style>
  
  <div class="login-page-container">
    <SignupForm {error} on:submit={handleLogin} />
  </div>