



export class BaseFormStore {

    errText = $state('invisible')
    errClass = $state('invisible')
    
    onSubmit = async (e) => {
        console.error("no onsubmit method set")
    }

    setErr(errString) {
        this.errText = errString;
        this.errClass = "visible";
    }

    hideErr() {
        this.errClass = "invisible";
    }

    setOnSubmit(onSubmitAsyncCallback) {
        this.onSubmit = onSubmitAsyncCallback;
    }


}
