class NavMenuStore {

    isOpen = $state(false)

    toggle() {
        this.isOpen = !this.isOpen
    }
    
}

export const navMenuStore = new NavMenuStore()