document.getElementById('reset-password-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    // Get the form data
    const formData = {
        username: document.getElementById('username').value,
        password: document.getElementById('password').value
    };

    // Send a POST request to reset the password
    fetch('http://127.0.0.1:3000/api/resetpassword/' + formData.username, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ password: formData.password })
    })
    .then(response => {
        if (response.ok) {
            console.log('Password reset successful');
            alert('Password reset successful');
            window.location.href = 'appointment.html';
        } else {
            throw new Error('Failed to reset password');
        }
    })
    .catch(error => {
        console.error('Error resetting password:', error);
        alert('Failed to reset password. Please try again.');
    });
});
