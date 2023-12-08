# OpenWeatherMap-Go-API-Server
Welcome to the OpenWeatherMap Go API Server, a simple yet powerful tool to fetch weather data using the OpenWeatherMap API. This lightweight Go application provides a RESTful API for querying weather information based on city names.
Features
OpenWeatherMap Integration: Seamlessly integrates with the OpenWeatherMap API to retrieve up-to-date weather data.

RESTful API: Offers a straightforward RESTful API endpoint for fetching weather information based on city names.

Configurable: Easily configure the OpenWeatherMap API key by loading the configuration from a JSON file.

Error Handling: Implements robust error handling for various scenarios, ensuring a reliable user experience.

Temperature Conversion: Automatically converts temperature from Kelvin to Celsius for convenience.

Getting Started
Clone this repository: git clone https://github.com/MohammedSalman999/openweathermap-go-api.git
Create an OpenWeatherMap API key and add it to the .apiConfig JSON file.
Build and run the server: go build && ./openweathermap-go-api
Access the API at http://localhost:8000/weather/{city} in your browser or through your favorite API client.
API Endpoints
Hello World: Visit http://localhost:8000/hello to receive a warm greeting from the Go server.

Weather Endpoint: Fetch weather data for a specific city by visiting http://localhost:8000/weather/{city}.

Configuration
Edit the .apiConfig JSON file to set your OpenWeatherMap API key:

json
Copy code
{
  "OpenWeatherApiKey": "your-api-key-here"
}
Dependencies
This project uses standard Go libraries and does not require any external dependencies.

Contributions
Contributions and feature requests are welcome! Feel free to open issues and pull requests to enhance this Go API server.

License
This project is licensed under the MIT License - see the LICENSE.md file for details.

Acknowledgments
Special thanks to the Go community for their excellent language and libraries.
