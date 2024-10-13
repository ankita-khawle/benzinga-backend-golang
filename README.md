# benzinga-backend-golang

Summary:

Endpoints:

1. GET /healthz: Health check API returning HTTP 200 OK.
2. POST /log: Accepts JSON payloads, deserializes, and stores them in-memory.

--------------------------

Environment Variables:

1. BATCH_SIZE: Number of logs before sending to an external endpoint.
2. BATCH_INTERVAL: Time in seconds between sending batches.
3. POST_ENDPOINT: URL endpoint to send the batch of data.

--------------------------

Batch Processing:

1. Send data when either the batch size is reached or the batch interval has passed.
2. Clear in-memory cache after sending.
3. Retry sending data up to 3 times if the post fails, with a 2 Second delay between retries.
4. Log each request, and provide structured, leveled logging for tracking and debugging.

--------------------------

Docker and Git:

1. Deployed on public git repository 
2. dockerfile for docker related configurations