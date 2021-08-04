
document.addEventListener("DOMContentLoaded", () => {
    const elems = document.querySelectorAll(".dropdown-trigger");
    const instances = M.Dropdown.init(elems, {
        alignment: "right",
        inDuration: 300,
    });
});
