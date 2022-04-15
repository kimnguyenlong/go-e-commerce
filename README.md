# Introductions
#### This is a backend project written in Go which provides Rest APIs of a e-commerce system. This project uses microservices architecture and includes 5 services: customer, provider, product, cart and order. Each service has some apis listed bellow: 
- Customer: register/login
- Provider: register/login
- Product: create, get, update, delete products and product's reviews
- Cart: set/get customer's cart
- Order: create, get customer's order
#### Used in this project:
- Gin Framework ([https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin))
- gRPC (for communnicating between services)
- MongoDB (for storing data)
- Redis (for caching)
# How to use
After cloning this repo, you can run the system via the following ways:
#### With Docker
Run cmd `docker-compose build` and then `docker-compose up`, then the services will be available on:
- Customer: localhost:9000
- Provider: localhost:9001
- Product: localhost:9002
- Cart: localhost:9003
- Order: localhost:9004
