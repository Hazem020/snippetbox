# Snippetbox

Snippetbox is a full-stack web application built using Go. It allows users to create, view, and manage text snippets, similar to GitHub Gists. The application includes features such as user authentication, protected endpoints, RESTful routing, and integration with a MongoDB database.

## Features

- **Create Snippets:** Users can create text snippets with a title, content, and expiration time (1 day, 7 days, or 365 days).
- **View Snippets:** Users can view individual snippets by navigating to a unique URL.
- **Latest Snippets:** The home page displays the latest snippets.
- **User Authentication:** Signup and login functionality with form validation.
- **JSON API:** A simple JSON API to fetch the latest snippets.
- **Go Templates:** Dynamic server-side HTML rendering using Go templates.
- **Security Features:** The app uses CSRF protection, form validation, and secure session management.

## Routes

- **GET /**: Home page displaying the latest snippets.
- **GET /snippet/view/:id**: View a specific snippet by its ID.
- **GET /snippet/create**: Display the form to create a new snippet.
- **POST /snippet/create**: Submit the form to create a new snippet.
- **GET /snippet/latest**: Retrieve the latest snippets in JSON format.
- **GET /user/signup**: Display the user signup form.
- **POST /user/signup**: Submit the user signup form.
- **GET /user/login**: Display the user login form.
- **POST /user/login**: Authenticate and login the user.
- **POST /user/logout**: Logout the current user.
- **GET /static/*filepath**: Serve static files like CSS and JavaScript.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/snippetbox.git
   cd snippetbox
