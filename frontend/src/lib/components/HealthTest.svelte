<!-- src/lib/components/HealthCheck.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
  
    let message: string = '';
    let loading: boolean = true;
    let error: string | null = null;
  
    onMount(async () => {
      try {
        const response = await fetch('http://localhost:8080/v1/healthz');
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        message = data.Message; // Extract the 'Message' field from the response
      } catch (err) {
        error = (err as Error).message;
      } finally {
        loading = false;
      }
    });
  </script>
  
  <style>
    .status-message {
      font-size: 1.2rem;
      font-weight: bold;
    }
    .loading {
      color: gray;
    }
    .error {
      color: red;
    }
  </style>
  
  {#if loading}
    <p class="loading">Checking server status...</p>
  {:else if error}
    <p class="error">Error: {error}</p>
  {:else}
    <p class="status-message">{message}</p>
  {/if}
  