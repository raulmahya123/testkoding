document.getElementById('course-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    // Get the form data
    const formData = {
        title: document.getElementById('title').value,
        slug: document.getElementById('slug').value,
        description: document.getElementById('description').value,
        privacy: document.getElementById('privacy').value,
        start_at: document.getElementById('start_at').value,
        end_at: document.getElementById('end_at').value,
        image: document.getElementById('image').value,
        certificate: document.getElementById('certificate').checked,
        level: document.getElementById('level').value,
        price: parseFloat(document.getElementById('price').value), // Convert to float
        status_enum: document.getElementById('status_enum').value
    };

    // Get JWT token from localStorage
    const token = localStorage.getItem('jwt_token');

    // Send a POST request to the API endpoint with JWT token
    fetch('http://127.0.0.1:3000/api/createcourse', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(formData)
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Failed to create course');
        }
    })
    .then(data => {
        console.log('Course created successfully:', data);
        alert('Course created successfully');
        // Redirect or perform other actions as needed
    })
    .catch(error => {
        console.error('Error creating course:', error);
        alert('Failed to create course. Please try again.');
    });
});

document.getElementById('search-form').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission

    // Get JWT token from localStorage
    const token = localStorage.getItem('jwt_token');

    // Send a GET request to the API endpoint with JWT token
    fetch('http://127.0.0.1:3000/api/getcourse', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Failed to fetch courses');
        }
    })
    .then(data => {
        // Handle successful response here
        console.log('Courses fetched successfully:', data);
        // Perform actions with fetched courses
    })
    .catch(error => {
        // Handle error here
        console.error('Error fetching courses:', error);
        alert('Failed to fetch courses. Please try again.');
    });
});

