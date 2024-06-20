

<script>
    import FormErr from "./bits/FormErr.svelte";

    
    let { children, onSubmit, title, imgSrc, errText, errClass, id, } = $props()
    
    let hiddenFileInput = $state();
    let image = $state(imgSrc)

    const onImageUpload = (e) => {
        errClass = "invisible"
        const file = hiddenFileInput.files[0];
        const fileName = file.name;
        const ext = fileName.split('.').pop();
        if (ext != "jpg" && ext != "jpeg" && ext != "png") {
            errClass = 'visible'
            errText = "png/jpg only"
            return;
        }
        const reader = new FileReader();
        reader.onload = (e) => {
            image = e.target.result;
        }
        reader.readAsDataURL(file);
    }

</script>





<form id={id} onsubmit={async (e) => {await onSubmit(e)}} class='flex flex-col gap-12'>
    
    {#if image != ""}
        <div class='flex flex-row justify-between gap-8'>
            <h2 class='text-2xl font-semibold'>{title}</h2>
            <div class='flex flex-col gap-4'>
                <div class='w-[150px] h-[150px] flex items-center'>
                    <img src="{image}" alt='cfa location' class='rounded-full' />
                </div>
                <button onclick={() => {hiddenFileInput.click()}} type='button' class='border transition duration-200 hover:bg-black hover:text-white px-2 py-1 rounded bg-white border border-gray-500 text-xs'>Upload</button>
            </div>
            <input bind:this={hiddenFileInput} type='file' class='hidden' oninput={() => {onImageUpload()}} />
        </div>   
    {:else}
        <h2 class='text-2xl font-semibold'>{title}</h2>
    {/if}

    {#if children}
        <FormErr errText={errText} errClass={errClass} />
        {@render children()}
    {/if}
</form>