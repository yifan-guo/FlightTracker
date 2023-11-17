# Flight Tracker App

## Introduction
Service used to calcualte the start and final destinations for customers


## Format
The service listens on port 8080 for incoming requests and exposes a POST endpoint called `/calculate`

Request URL: 
`http://localhost:8080/calculate`

Description:
 - Accepts a request body containing the various legs of a customer's flight.
 - Returns a list containing the start and final destination of the flight

Request Headers:
```azure
Content-Type: application/json
```

Example Request Body
```azure
[
    {
        "start": "IND",
        "end": "EWR"
    },
    {
        "start": "SFO",
        "end": "ATL"
    },
    {
        "start": "GSO",
        "end": "IND"
    },
    {
        "start": "ATL",
        "end": "GSO"
    }
]

```

Response Status Codes
```
Success: 200
BadRequest: 400
InternalServerError: 500
```

Response Body (For above input)
```
[
	"SFO",
	"EWR"
]
```

![Screenshot 2023-11-17 at 11.24.45 AM.png](..%2F..%2FDesktop%2FScreenshot%202023-11-17%20at%2011.24.45%20AM.png)


## Optimizations (Future Enhancements)
There are three goals in building a microservice: scalability, availability, and latency.

### Scalability
To increase the scale, I would containerize the application and deploy on AWS ECS cluster for managed scaling.
As this is a small application, ECS is ideal for its ease of deployment and configuration. 

However, if the flight tracker service expands and needs to accept traffic from vendors spanning multiple cloud providers, I would opt for AWS EKS or Kubernetes for their cross-platform versatility.  

### Latency
To optimize latency, I would add a cache to reduce the response time for requests with more than 50+ legs. Caching allows requests to the graph construction and traversal and bring down the p95 response time by 60-70%

### Availability
I would deploy the microservice in multiple regions and have a Route53 (DNS load balancer) to distribute traffic across all regions. In each region, I would configure instances of the microservice across multiple availability zones sitting behind a network load balancer to distribute traffic among them.


