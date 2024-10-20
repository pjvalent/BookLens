<!-- src/routes/+layout.svelte -->
<script lang="ts">
    import Header from '$lib/components/Header.svelte';
    import { page } from '$app/stores';

    export let data;
    $: hideHeader = $page.url.pathname === '/welcome';
    
  </script>
  
  <style>
    .layout {
      display: flex;
      flex-direction: column;
      min-height: 100vh;
    }
  
    main {
      flex: 1;
    }
  </style>
  
  <div class="layout">
    {#if !hideHeader}
    <Header title="BookLense" />
    {/if}
    <main>
      <slot />
    </main>
  </div>