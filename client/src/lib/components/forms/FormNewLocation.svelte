

<script>

    import { BaseFormStore } from "./store.svelte.js";
    import { Request } from "$lib/http/Request.js";
    import FormShell from "./FormShell.svelte";
    import FormInput from "./bits/FormInput.svelte";
    import FormSubmit from "./bits/FormSubmit.svelte";

    let form = new BaseFormStore();
    let name = $state("")
    let number = $state("")

    form.setOnSubmit(async () => {
        form.hideErr()
        let res = await Request.post('/api/cfa', {name: name, number: number})
        let json = await res.json()
        console.log(json)
        if (res.status !== 201) {
            form.setErr(json.message)
            return;
        }
    })


</script>

<FormShell id='form-new-location' onSubmit={async (e) => {await form.onSubmit(e)}} title={"New CFA Location"} imgSrc={"/img/cfa.jpg"} errText={form.errText} errClass={form.errClass}>    
    <FormInput onInput={(e) => {name = e.target.value; form.hideErr()}} inputType={'text'} value={name} name="name" label="Name" />
    <FormInput onInput={(e) => {number = e.target.value; form.hideErr()}} inputType={'text'} value={number} name="number" label="Number" />
    <FormSubmit text="Create" />
</FormShell>

