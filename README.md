# Quick go api - Matt P.

### ✅ **Go-Kit Hello World API for quick start and learning**

This repository provides a simple "Hello World" API built using Go-Kit. 

Version 1.0.0 
> Basic Hello World API with Go-Kit framework. This is for me to get used to syntax and follow a guide while learning how to structure the api. Checkout the `goa-1` branch to see this. 

version 2.0.0
> Create a simplistic CMS similar to Prismic. I'm expecting this will be building a bit of knowledge on structuring a go API. Then will eventually include me trying to convert this out to work as either a ecs or a lambda. Checkout the `goa-2` branch to see this.

***
## How to run:

1. Clone the repository: 
   ```bash
   git clone
    ```

2. Navigate to the project directory:
    ```bash
    cd go-api
    ```

3. Run the application:
    ```bash
    go run main.go    
    ```
    or use the provided script:
    ```bash
    ./scripts/go-tools.sh start
    ```

4. Curl the API:
    Open your web browser or use curl to access the API endpoint:
    ```bash
    curl http://localhost:8080/form?id=registration
    ```
    You should see a JSON response:
    ```json
    {
      "form_id": "registration",
      "elements": [
        {"type": "text", "label": "Name", "name": "name", "required": false},
        {"type": "email", "label": "Email", "name": "email", "required": false}
      ]
    }
    ```
    Theres also a few other forms to curl. Check the script file for more info.