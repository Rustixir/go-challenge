Here's a well-written document based on your description:


---

Service Design and Implementation Overview

Overview

This document outlines the architectural and implementation choices for the service, highlighting the tools, libraries, and design patterns used. The goal is to build a robust, maintainable, and efficient service with support for multiple storage backends and localization.


---

Design Approach

1. Clean Architecture

Reasoning:

Clean Architecture separates concerns and ensures that core business logic is independent of frameworks, libraries, and data sources.

This improves testability, scalability, and maintainability.

Dependencies flow inward, so the application’s core is insulated from external changes.


Implementation:

The service is divided into layers:

Entities (Core Business Models): Define high-level abstractions.

Use Cases (Business Logic): Contain application-specific business rules.

Interface Adapters (Handlers, Repositories): Convert data between formats suitable for the use cases and the external systems.

Frameworks & Drivers: Handle interactions with external systems like HTTP or databases.





---

2. REST API with Echo HTTP Framework

Reasoning:

REST APIs are stateless and widely adopted for web services, offering simplicity, scalability, and interoperability.

The Echo HTTP library is chosen for its lightweight, high-performance nature and built-in support for middleware, routing, and error handling.


Implementation:

RESTful routes are implemented using Echo to handle client interactions.

A custom middleware for detecting the Accept-Language HTTP header is used for localization.




---

3. Repository Pattern

Reasoning:

Decouples the business logic from data access logic.

Allows switching between different data sources (e.g., SQLite and Redis) by swapping repository implementations.


Implementation:

An interface defines the repository contract.

Two concrete implementations:

SQLite-based Repository (using Ent ORM).

Redis-based Repository (using go-redis).


Switching between implementations is controlled via configuration, promoting flexibility.




---

Tools and Libraries

1. Ent ORM for SQLite

Reasoning:

Type-Safe: Eliminates runtime errors by leveraging Go’s type system.

Performance: Provides better performance compared to alternatives like GORM.

Code Generation: Automatically generates schema, models, and helper methods.

Fluent API: Offers a clean and expressive syntax for querying.


Usage:

Ent is used for managing user segments in SQLite with type-safe queries.

Schema is defined in code, enabling easy migrations and schema changes.




---

2. go-redis for Redis Integration

Reasoning:

A robust and performant Redis driver for Go.

Provides support for Redis features like sorted sets, pipelines, and transactions.


Usage:

User segments are stored in Redis sorted sets (ZAdd), with timestamps as scores.

This allows efficient querying of users within a specific timeframe (e.g., the past 14 days).

A Cleanup Worker is implemented to periodically remove expired users from Redis.




---

3. Accept-Language Header for Localization

Reasoning:

Detecting user language preferences based on the Accept-Language header enables better user experiences by serving localized content or error messages.


Usage:

A middleware extracts the preferred language (en for English or fa for Persian) and makes it available throughout the application.

A custom error handler maps business logic errors to localized messages via a Localization Component.




---

4. Custom Error Handler

Reasoning:

Provides a centralized mechanism for handling errors and returning user-friendly messages.


Usage:

The error handler maps internal business layer errors to specific user-facing messages, formatted in the user’s preferred language.

This ensures consistency and proper localization of error responses.




---

Cleanup Worker

SQLite Implementation

Logic:

Deletes rows from the database where created_at is older than 14 days.


Testing:

Unit tests validate the logic by inserting mock data and verifying the expected state after cleanup.



Redis Implementation

Logic:

Uses the ZRemRangeByScore command to remove members from sorted sets where the score (timestamp) is older than 14 days.


Testing:

Mock data is populated in Redis, and the cleanup process is validated using assertions.




---

Testing

1. Unit Tests

Separate unit tests are written for both the SQLite and Redis repository implementations.

Tests verify the functionality of:

Create: Ensures data is stored correctly.

Count: Verifies querying logic for user counts within the last 14 days.

Cleanup: Validates the removal of expired users.



2. Integration Tests

End-to-end tests ensure that the service works seamlessly with both Redis and SQLite backends.



---

Makefile

Commands

Run the Application:

run

Installs required dependencies (e.g., build-essential) and sets CGO_ENABLED=1 for SQLite support.


Run Tests:

test

Executes all unit and integration tests.


Generate Ent Schema:

ent_gen_schema

Creates new schema files for Ent ORM.


Generate Ent Code:

ent_gen

Regenerates Ent models and helpers after schema changes.




---

Conclusion

The service is built using clean architecture principles, ensuring high maintainability and testability. By leveraging SQLite with Ent ORM and Redis with go-redis, the service offers flexibility in data storage. Localization and error handling enhance user experience, while automated tests and a Makefile streamline development and deployment.

