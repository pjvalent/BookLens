<!-- src/routes/+layout.svelte -->
<script lang="ts">
    import Header from '$lib/components/Header.svelte';
    import { onMount } from 'svelte';

    export let data;
    
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
    <Header title="BookLense" />
    <main>
      <slot />
    </main>
  </div>