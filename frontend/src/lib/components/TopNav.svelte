<script lang="ts">
    import { onMount, onDestroy } from 'svelte';
    import { goto } from '$app/navigation';
    import { browser } from '$app/environment';
  
    let showMenuDropdown = false;
    let showProfileDropdown = false;
    let searchQuery = '';
  
    let menuDropdownEl: HTMLElement;
    let profileDropdownEl: HTMLElement;
  
    function toggleMenuDropdown(event: Event) {
      showMenuDropdown = !showMenuDropdown;
      if (showMenuDropdown) {
        showProfileDropdown = false;
      }
      event.stopPropagation();
    }
  
    function toggleProfileDropdown(event: Event) {
      showProfileDropdown = !showProfileDropdown;
      if (showProfileDropdown) {
        showMenuDropdown = false;
      }
      event.stopPropagation();
    }
  
    function logout() {
      goto('/login');
    }
  
    function handleSearch(event: KeyboardEvent) {
      if (event.key === 'Enter') {
        // Implement search functionality
        console.log('Searching for:', searchQuery);
      }
    }
  
    function handleClickOutside(event: MouseEvent) {
      if (showMenuDropdown && menuDropdownEl && !menuDropdownEl.contains(event.target as Node)) {
        showMenuDropdown = false;
      }
      if (showProfileDropdown && profileDropdownEl && !profileDropdownEl.contains(event.target as Node)) {
        showProfileDropdown = false;
      }
    }
  
    onMount(() => {
      document.addEventListener('click', handleClickOutside);
    });
  
    onDestroy(() => {
      if (browser) {
        document.removeEventListener('click', handleClickOutside);
      }
    });
  </script>
  
  
  <style>
    .top-nav {
      display: flex;
      align-items: center;
      justify-content: space-between;
      background-color: #3b3b3b;
      color: #ffffff;
      padding: 1rem;
      width: 100%;
      box-sizing: border-box;
      margin: 0;
    }
  
    .nav-left,
    .nav-right {
      display: flex;
      align-items: center;
    }
  
    .nav-left > * {
      margin-right: 1rem;
    }
  
    .nav-right > * {
      margin-left: 1rem;
    }
  
    .dropdown {
      position: relative;
    }
  
    .dropdown-content {
      display: none;
      position: absolute;
      background-color: #ffffff;
      min-width: 160px;
      box-sizing: border-box;
      box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
      top: 100%;
      z-index: 1;
    }
  
    /* Menu dropdown positioning */
    .nav-left .dropdown-content {
      left: 0;
    }
  
    /* Profile dropdown positioning */
    .nav-right .dropdown-content {
      right: 0;
    }
  
    .dropdown-content a {
      color: #1b4d3e;
      padding: 12px 16px;
      text-decoration: none;
      display: block;
    }
  
    .dropdown-content a:hover {
      background-color: #ddd;
    }
  
    .show {
      display: block;
    }
  
    /* Optional: Style buttons */
    button {
      background: none;
      border: none;
      color: inherit;
      font: inherit;
      cursor: pointer;
    }
  </style>
  
  <nav class="top-nav">
    <div class="nav-left">
      <!-- Menu Dropdown -->
      <div class="dropdown" bind:this={menuDropdownEl}>
        <button
          on:click|stopPropagation={toggleMenuDropdown}
          aria-haspopup="true"
          aria-expanded={showMenuDropdown}
        >
          Menu
        </button>
        <div class="dropdown-content" class:show={showMenuDropdown}>
          <a href="/categories">Categories</a>
          <a href="/bestsellers">Best Sellers</a>
          <!-- Add more menu items -->
        </div>
      </div>
  
      <!-- Search Bar -->
      <input
        type="text"
        placeholder="Search books..."
        bind:value={searchQuery}
        on:keydown={handleSearch}
      />
    </div>
  
    <div class="nav-right">
      <!-- Profile Dropdown -->
      <div class="dropdown" bind:this={profileDropdownEl}>
        <button
          on:click|stopPropagation={toggleProfileDropdown}
          aria-haspopup="true"
          aria-expanded={showProfileDropdown}
        >
          Profile
        </button>
        <div class="dropdown-content" class:show={showProfileDropdown}>
          <a href="/profile">My Profile</a>
          <a href="#" on:click|preventDefault={logout}>Logout</a>
        </div>
      </div>
    </div>
  </nav>
  