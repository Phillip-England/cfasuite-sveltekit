

<script>
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { navMenuStore } from "./NavMenu.svelte.js";

    let { href, text } = $props()
    let isActive = $page.url.pathname == href

</script>

{#if isActive}

    <li class='transition duration-200 border rounded border-gray-100 flex cursor-pointer bg-black text-white'>
        <button onclick={() => {
            navMenuStore.toggle()
            goto(href)
        }} class='p-4 text-sm w-full flex'>{text}</button>
    </li>

{:else}

    <li class='transition duration-200 border rounded border-gray-100 flex cursor-pointer hover:text-white'>
        <button onclick={() => {
            navMenuStore.toggle()
            goto(href)
        }} class='p-4 text-sm w-full flex'>{text}</button>
    </li>

{/if}
