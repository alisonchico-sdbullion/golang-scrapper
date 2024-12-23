# Quotes Scraper App

This project is a simple Go web application that scrapes quotes from [Quotes to Scrape](https://quotes.toscrape.com) and serves them in JSON format through an HTTP endpoint.

## Features
- Scrapes up to 100 quotes from the website.
- Returns quotes in JSON format via the `/quotes` endpoint.

---

## Prerequisites

Ensure you have the following installed:
- [Go (1.20 or later)](https://go.dev/dl/)
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/) (optional, for containerized builds)

---

## Local Setup Instructions

### Step 1: Clone the Repository
```bash
git clone git@github.com:alisonchico-sdbullion/golang-scrapper.git
cd golang-scrapper
```

### Step 2: Install Dependencies
Go modules are used to manage dependencies.
```bash
go mod tidy
```

### Step 3: Build the Application
Run the following command to build the application:
```bash
go build -o main .
```
This will generate an executable named `main` in the project directory.

### Step 4: Run the Application
Execute the application:
```bash
./main
```
The server will start on `http://localhost:8080`. You should see the following log message:
```
2024/12/23 16:37:35 Server is running on http://localhost:8080
```

### Step 5: Access the API
Visit the following endpoint in your browser or use `curl`:
```bash
curl http://localhost:8080/quotes
```
The response will be a JSON array of quotes:
```json
[
  {
    "text": "The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking.",
    "author": "Albert Einstein"
  },
  {
    "text": "It is our choices, Harry, that show what we truly are, far more than our abilities.",
    "author": "J.K. Rowling"
  }
]
```

---

## Running with Docker

### Step 1: Build the Docker Image
```bash
docker build -t quotes-scraper .
```

### Step 2: Run the Docker Container
```bash
docker run -p 8080:8080 quotes-scraper
```

### Step 3: Access the API
Visit the endpoint in your browser or use `curl`:
```bash
curl http://localhost:8080/quotes
```

## Running with Pipeline

This automation use 3 major automations to deploy a Golang app Container.
 - Docker for containerization (detailed on Containerization.md) 
 - Terraform for infrastructure provision (detailed on Infrastructure.md)
 - Github action workflow for ci/cd provision (detailed on Pipeline.md)

As the purporse of this code is to apply for a exam i detailed additional roadmap that would enhance the automation of the code but was not applied due to time limit.


- If is the first usage you need to follow some specific steps before you trigger the pipeline you need to create a service account on aws, download the security credentials and add on github secrets and variables on settings page of your repo, you need to add this secrets
  - AWS_ACCESS_KEY_ID
  - AWS_REGION
  - AWS_SECRET_ACCESS_KEY    
- After that you will need to apply ecr, s3 bucket and dynamodb tables mannualy, as this is necessary for the enablement of terraform and so the pipeline be able to push the build artefact, to do that you need to install terraform on your computer (latest version) and apply the following targets:
  - terraform init
  - terraform apply -target=module.s3_tf-state.aws_s3_bucket.this[0]
  - terraform apply -target=aws_ecr_repository.solo
  - terraform apply -target=aws_dynamodb_table.terraform_locks
- Uncomment the code on backend.tf  
- By end execute terraform init --migrate-state that will move your state to the bucket, following this approachs you can trigger the pipeline.

---

## Notes
- Ensure you have internet access while running the scraper, as it fetches quotes from the web.
- If the API returns an empty response, inspect the logs for error messages.

---

## Troubleshooting

### Common Issues

1. **No Quotes Returned**
   - Verify the website `https://quotes.toscrape.com` is accessible.
   - Check the logs for scraping errors.

2. **Port Already in Use**
   - Change the port in the code (`main.go`) to an available port.

3. **Docker Networking Issues**
   - Ensure Docker has internet access and restart your Docker daemon if necessary.

---

## Contributing
Feel free to submit issues or contribute to the project by creating a pull request.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

