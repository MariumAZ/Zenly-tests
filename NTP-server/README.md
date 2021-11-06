# NTP-server 



Create a server (and demo client) using Go and gRPC that responds to time requests and fetches the time via SNTP or NTP on time.apple.com and time.microsoft.com.
The client can pass which NTP server it would like to query, and the response will be a time stamp to be properly printed to the console.