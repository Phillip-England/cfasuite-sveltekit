




<script>

    import { goto } from "$app/navigation";
    import Header from "$lib/components/nav/Header.svelte";
    import { Request } from "$lib/http/Request";
    import { onMount } from "svelte";

    onMount(async () => {
        let res = await Request.get('/api/auth');
        if (res.status != 200) {
            goto('/')
        }
    })

    let { children } = $props();
    
</script>
    
    
    
    
    <div>
        <Header showBars={true} />
        <main class='p-6'>
            {#if children}
                {@render children()}
            {/if}
        </main>
    </div>
    