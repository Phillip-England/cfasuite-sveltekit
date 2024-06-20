

<script>

    let { inputType, value, name, label, onInput } = $props()

</script>



<div class='form-input flex flex-col gap-1 text-sm'>
    <label for={name}>{label}</label>
    <input value={value} oninput={async (e) => {await onInput(e)}} class='border rounded p-1 transition duration-200 focus:outline-none border-gray-100 focus:border-gray-500' type={inputType} id={name} name={name} autocomplete="">
</div>


