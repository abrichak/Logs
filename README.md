# Overview
Metrics service

## Usage
1. Copy .env.dist to .env and set the environment variables, if necessary.
2. Run your application using the command in the terminal:

    `docker-compose up`
3. Browse to {HOST}:{EXPOSE_PORT}/swagger/index.html. You will see Swagger 2.0 API documents.
4. Using the API documentation, make requests to save log messages with different users' IPs.
5. Browse to {HOST}:{EXPOSE_PORT_METRICS}/metrics. You will see the number of unique users' IPs.
