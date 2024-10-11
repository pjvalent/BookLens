<script lang="ts">
    import { Header } from '$lib'
    import { goto } from '$app/navigation';

    function navigateToHealthTest() {
        goto('/healthTest');
  }

    function navigateToSignup() {
      goto('/signup');
    }
</script>

<style>
  .home-page {
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 2rem;
  }

  h1 {
    color: #1b4d3e; /* Dark green */
    margin-bottom: 1rem;
  }

  p {
    margin-bottom: 2rem;
  }

  button {
    background-color: #ffffff; /* Dark green */
    color: #1b4d3e; /* White text */
    border: none;
    border-radius: 4px;
    padding: 0.75rem 1.5rem;
    font-size: 1rem;
    cursor: pointer;
    margin: 0.5rem;
  }

  button:hover {
    background-color: #166a52;
  }
</style>

<div class="home-page">
  <h1>Welcome to BookLense</h1>
  <p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
  <div>
    <button on:click={navigateToHealthTest}>Go to Health Test</button>
    <button on:click={navigateToSignup}>Sign Up</button>
  </div>
</div>
