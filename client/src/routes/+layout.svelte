




<script>

let { children } = $props();

</script>




<div>
    {#if children}
        {@render children()}
    {/if}
</div>
