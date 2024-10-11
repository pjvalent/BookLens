<!-- src/routes/signup/+page.svelte -->
<script lang="ts">
  import SignupForm from '$lib/components/SignupForm.svelte';
  import { goto } from '$app/navigation';

  let error: string | null = null;

  async function handleSignup(event: CustomEvent) {
    const { firstName, lastName, email } = event.detail;
    error = null;

    try {
      const response = await fetch('http://localhost:8080/v1/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          email: email,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to sign up');
      }

      // On successful signup, redirect the user
      goto('/welcome');

    } catch (err) {
      error = (err as Error).message;
    }
  }
</script>

<style>
  .signup-page-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: 1rem;
  }
</style>

<div class="signup-page-container">
  <SignupForm {error} on:submit={handleSignup} />
</div>