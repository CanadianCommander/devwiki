// little drag "library" for dragging elements horizontally. I'm getting the sense that it will
// be a careful balance between using JS raw and switching over to Angular Elements. htmx certainly
// doesn't provide the same level of interactivity as Angular.


const dragListeners = new Map();
const dragEndListeners = new Map();

/**
 * Start dragging an element horizontally by its id
 * @param targetId - id of the element to drag
 * @param event - [optional] event that triggered the drag
 */
window.startDrag = function startDrag(targetId, event = null) {
    event?.preventDefault();

    dragListeners.set(targetId, (event) => dragNav(event, targetId, 175, window.innerWidth - 400));
    dragEndListeners.set(targetId, (event) => endDrag(event, targetId));

    window.addEventListener("mousemove", dragListeners.get(targetId));
    window.addEventListener("mouseup", dragEndListeners.get(targetId));
}

function endDrag(event, targetId) {
    window.removeEventListener("mousemove", dragListeners.get(targetId));
    window.removeEventListener("mouseup", dragEndListeners.get(targetId));

    dragEndListeners.delete(targetId);
    dragListeners.delete(targetId);
}

/**
 * Drag an element horizontally by its id
 * @param event - event that triggered the drag
 * @param targetId - id of the element to drag
 * @param min - minimum width of the element
 * @param max - maximum width of the element
 */
function dragNav(event, targetId, min, max) {
    const leftNav = document.getElementById(targetId);
    leftNav.style.width = `${Math.min(Math.max(event.clientX + 2, min), max)}px`;
}