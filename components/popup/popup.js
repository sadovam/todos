
const popup = (text, style) => {
    const root = document.getElementById("popup");
    const message = document.createElement("h4");
    message.innerHTML = text;
    message.className = "popup__mess popup__mess_" + style;
    root.appendChild(message);
    setTimeout(() => message.classList.add("popup__mess_visible"), 100);
    setTimeout(() => message.classList.remove("popup__mess_visible"), 2000);
    setTimeout(() => message.remove(), 2500);
}
