# URL Shortener

A basic URL shortener service in Go that creates short URLs and redirects users to the original URLs. It supports generating short URLs with MD5 hashing and stores them in memory. To use, send a POST request to `/shorten` with the URL in JSON format to get a shortened URL, or access `/redirect/{id}` to redirect to the original URL. The server runs on port 3000 and includes CORS support for easy frontend integration.
