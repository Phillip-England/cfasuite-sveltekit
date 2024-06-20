

<script>
    import { Request } from "$lib/http/Request.js";
    import { navMenuStore } from "./NavMenu.svelte.js"
    import NavMenuItem from "./NavMenuItem.svelte";
    

    async function logout() {
        let res = await Request.post("/api/auth/logout", {})
        if (res.status != 200) {
            console.error('failed to logout, redirecting to login page...')
            window.location.href = '/'
        }
        window.location.href = '/'
    }

</script>

{#if navMenuStore.isOpen}
    <nav class='absolute top-0 left-0 flex flex-col w-[66%] bg-white border-r border-gray-100 h-full z-30'>
        <div class='border-b border-gray-100 h-[76px] p-4'>
            <div class='w-[150px]'>
                <img src='/img/logo.svg' alt='logo' >
            </div>
        </div>
        <ul class='flex flex-col gap-1 p-1'>
            <NavMenuItem href='/admin' text='Admin Panel'/>
            <li class='transition duration-200 border rounded border-gray-100 flex cursor-pointer hover:text-white hover:bg-primary'>
                <button onclick={async () => {await logout()}} class='p-4 text-sm w-full flex'>Logout</button>
            </li>
        </ul>
    </nav>
    <button onclick={() => {navMenuStore.toggle()}} class='bg-black opacity-50 z-20 h-full w-full absolute top-0 left-0'></button>
{/if}