
<script>

    import { navMenuStore } from "./NavMenu.svelte.js"
    let { showBars } = $props();

</script>

<header class='p-4 border-b border-gray-100 flex flex-row justify-between items-center h-[76px]'>
    <div class='w-[150px]'>
        <img src='/img/logo.svg' alt='logo' >
    </div>
    {#if showBars}
        <button onclick={() => {navMenuStore.toggle()}} class='flex flex-row justify-between items-center'>
            <svg class="w-8 h-8" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                <path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 7h14M5 12h14M5 17h14"/>
            </svg>
        </button>
    {/if}
</header>