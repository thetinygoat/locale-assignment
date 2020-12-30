# Locale Assignment
## Dependencies
 - Redis - Redis must be installed and running as the task queue depends on redis.
 - PostgreSQL - PostgreSQL is the persistance layer of choice to store data.

## Code Structure
The code structre is based upon very popular [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

- Entities are the blueprint of the data (model).
- Service is the layer that abstracts away business logic.
- Handler handles connections, in this case we have only HTTP handlers but the functionality can be extended to support other protocols easily.
- Repository is the layer that abstracts away the persistance layer. We can easily switch our persistance layer from PostgreSQL to something like MongoDB without breaking other layers.