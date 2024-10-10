<script lang="ts">
    import { Header } from '$lib'
    import { goto } from '$app/navigation';

    function navigateToHealthTest() {
        goto('/healthTest');
  }
</script>

<Header title="BookLense"></Header>
<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<button on:click={navigateToHealthTest}>Go to Health Test</button>
