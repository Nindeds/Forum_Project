document.addEventListener('DOMContentLoaded', function() {
    var modal = document.getElementById("post-modal");
    var btn = document.getElementById("create-post-button");
    var span = document.getElementsByClassName("close")[0];

    btn.onclick = function() {
        modal.style.display = "block";
    }

    span.onclick = function() {
        modal.style.display = "none";
    }

    window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }

    document.getElementById('post-form').addEventListener('submit', function(event) {
        event.preventDefault();
        var title = document.getElementById('post-title').value;
        var message = document.getElementById('post-message').value;
        var image = document.getElementById('post-image').files[0];
        
        if (title && message) {
            var postList = document.getElementById('post-list');
            var postId = Date.now(); 
            var newPost = document.createElement('li');
            newPost.className = 'post-summary';
            newPost.innerHTML = `
                <div class="post-title">${title}</div>
                <div class="post-content">${message.substring(0, 100)}...</div>
                <div class="post-date">${new Date().toLocaleString()}</div>
            `;
            newPost.setAttribute('data-id', postId);
            newPost.addEventListener('click', function() {
    
                window.location.href = `postDetail.html?id=${postId}`;
            });
            postList.appendChild(newPost);

           
            var postData = { title, message, date: new Date().toLocaleString(), image };
            localStorage.setItem(postId, JSON.stringify(postData));
            
            document.getElementById('post-title').value = '';
            document.getElementById('post-message').value = '';
            document.getElementById('post-image').value = '';

            modal.style.display = "none";
        } else {
            alert("Please fill out the title and message fields.");
        }
    });
});
