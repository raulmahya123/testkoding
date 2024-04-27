document.getElementById('login-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    // Get the form data
    const formData = {
        username: document.getElementById('username').value,
        password: document.getElementById('password').value
    };

    // Send a POST request to the API endpoint
    fetch('http://127.0.0.1:3000/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Login failed');
        }
    })
    .then(data => {
        // Handle successful login response here
        console.log('Login successful:', data);
        
        // Save JWT token to localStorage
        localStorage.setItem('jwt_token', data.response.token);

        if (data.response.role === 'admin') {
            // Redirect to dashboard for admin
            window.location.href = '../dashboard/dark-uibank-dashboard-concept/dist/index.html';
        } else {
            // Redirect to index.html for regular users
            window.location.href = 'index.html';
        }
    })
    .catch(error => {
        // Handle error here
        console.error('Login error:', error);
        alert('Login failed. Please check your username and password.');
    });
});
