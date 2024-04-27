document.getElementById('register-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    // Get the form data
    const formData = {
        username: document.getElementById('username').value,
        password: document.getElementById('password').value,
        image: document.getElementById('image').value,
        email: document.getElementById('email').value,
        role: document.getElementById('role').value // Role set to "user"
    };
    console.log(formData)

    // Send a POST request to register the user
    fetch('http://127.0.0.1:3000/api/registerr', {
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
            throw new Error('Failed to register');
        }
    })
    .then(data => {
        console.log('Registration successful:', data);
        
        // Display success message to the user
        alert('Registration successful! Welcome, ' + data.data.username + '!');

        // Redirect the user to another page (e.g., login page)
        window.location.href = 'appointment.html';
    })
    .catch(error => {
        console.error('Error registering:', error);
        alert('Failed to register. Please try again.');
    });
});
