
    document.getElementById('post-button').onclick = function() {
    document.getElementById('post-modal').style.display = 'block';
}
    document.getElementById('close-modal').onclick = function() {
    document.getElementById('post-modal').style.display = 'none';
}
    window.onclick = function(event) {
    if (event.target == document.getElementById('post-modal')) {
    document.getElementById('post-modal').style.display = 'none';
}
}
