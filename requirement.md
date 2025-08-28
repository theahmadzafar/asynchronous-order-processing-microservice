Project: Asynchronous Order Processing Microservice

Objective
Build a Go microservice that receives OrderCreated events and processes them through a pipeline with the following steps:

Validation

Persistence

Notification

The application must use goroutines and channels to implement internal asynchronous processing, not just concurrent HTTP handling.

Requirements
1. Order Input
The service should expose an HTTP endpoint (or use a mocked message queue) to receive OrderCreated events.

Payload example:
{
  "order_id": "1234",
  "user_id": "5678",
  "items": ["item1", "item2"],
  "total": 49.99
}

2. Asynchronous Processing Pipeline
Implement an internal pipeline using goroutines and channels to simulate asynchronous processing.

Each stage (validation → persistence → notification) should run concurrently and communicate via channels.

3. Persistence Layer

Design it so the storage layer can be swapped with minimal changes.

4. Notification Step
Simulate sending an email or push notification after successful persistence (mocked logic is fine).

5. Concurrency Control
Use proper Go patterns for concurrency (goroutines, channels etc.)

6. Deployment
Pack the app in a multistage docker file. Stages: build, run