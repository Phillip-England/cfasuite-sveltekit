

<script>

    import { BaseFormStore } from "./store.svelte.js";
    import { Request } from "$lib/http/Request.js";
    import { goto } from "$app/navigation";
    import FormShell from "./FormShell.svelte";
    import FormInput from "./bits/FormInput.svelte";
    import FormSubmit from "./bits/FormSubmit.svelte";

    export const loginForm = new BaseFormStore();
    let email = $state("")
    let password = $state("")

    loginForm.setOnSubmit(async () => {
        loginForm.hideErr()
        let res = await Request.post('/api/auth/login', {email, password})
        let json = await res.json()
        if (res.status !== 200) {
            loginForm.setErr(json.message)
            return;
        }
        goto("/admin")
    })

</script>

<FormShell id='form-login' onSubmit={loginForm.onSubmit} title={"Login"} errText={loginForm.errText} errClass={loginForm.errClass} imgSrc="">
    <FormInput inputType={'text'} value={email} name="email" label="Email" onInput={(e) => {email = e.target.value; loginForm.hideErr()}} />
    <FormInput inputType={'password'} value={password} name="password" label="Password" onInput={(e) => {password = e.target.value; loginForm.hideErr('')}} />
    <FormSubmit text="Login" />
</FormShell>