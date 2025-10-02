const form = document.forms.namedItem("url-form")

form.addEventListener("submit", (event) => {
    event.preventDefault()
    event.stopPropagation()
})

form.addEventListener("change", () => {
    const url = form.elements.namedItem("url").value
    const limits = Number.parseInt(form.elements.namedItem("limits").value, 10)
    const results = form.elements.namedItem("results")
    const u = new URL("/feed", location.href)
    u.searchParams.set("url", url)
    if (Number.isSafeInteger(limits) && limits > 0) {
        u.searchParams.set("limits", String(limits))
    }
    results.value = u.toString()
})

form.addEventListener("click", (event) => {
    if (event.target.id !== "results") return
    event.target.select()
})