# RabbitMQ integration in Go

## Setup

**Run Docker Compose:**
   ```bash
   docker-compose up --build -d
   ```

**Wait for Initialization:**
   Wait for approximately 1-2 minutes for the services to initialize.

## Sending Messages

Make HTTP requests to the API endpoint with different text in the `msg` query parameter:

```bash
curl \
    --request GET \
    --url 'http://localhost:3000/send?msg=test'
```

## Accessing RabbitMQ GUI

1. Visit [http://localhost:15672/](http://localhost:15672/) in your browser.

2. Login to RabbitMQ GUI:
   - **Username:** guest
   - **Password:** guest

3. Explore RabbitMQ User Interface:
   - Navigate the RabbitMQ user interface to observe message queues and exchanges.

## Stopping the Environment

To stop the environment and clean up resources, run:

```bash
docker-compose down
```