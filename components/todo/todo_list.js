const layout__on_click = (e) => {
    e.preventDefault();
    input = document.getElementById("layout__input")
    form = document.getElementById("layout__form")
    fetch("/todo", {method: "post", body: new FormData(form)})
    .then((response) => {
        return response.json();
    })
    .then((data) => {
        if (data.error) {
            popup(data.error, "error");
        } else {
            popup("New item created", "info")
            document.getElementById("layout__todos")
            .insertAdjacentHTML("beforeend", data.node);
        }
    });
    input.value = ""
};

const layout__del = (uid) => {
    fetch("/todo/" + uid, {method: "delete"})
    .then(() => {
        document.getElementById("layout " + uid).remove();
    })
}
