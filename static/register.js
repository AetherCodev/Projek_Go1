document.getElementById("formRegister").addEventListener("submit", function(e) {
    const password        = document.getElementById("password").value
    const confirmPassword = document.getElementById("confirmPassword").value // ← samakan id!
    const errorJs         = document.getElementById("errorJs")

    errorJs.style.display = "none"
    errorJs.innerText     = ""

    if (password === "") {
        e.preventDefault()
        errorJs.style.display = "block"
        errorJs.innerText     = "Password tidak boleh kosong!"
        return
    }

    if (password.length < 8) {
        e.preventDefault()
        errorJs.style.display = "block"
        errorJs.innerText     = "Password minimal 8 karakter!"
        return
    }

    if (password !== confirmPassword) {
        e.preventDefault()
        errorJs.style.display = "block"
        errorJs.innerText     = "Password tidak cocok!"
        return
    }
})