This repository demonstrates code design patterns and anti-patterns in practice. 
The project structure demonstrates a convenient organization of layers and components.

### Patterns:
[Repository registry](https://github.com/vadiminshakov/dddgo/blob/main/repository/reporegistry.go#L30)

[Transactional method](https://github.com/vadiminshakov/dddgo/blob/main/repository/reporegistry.go#L55)

[Aggregate](https://github.com/vadiminshakov/dddgo/blob/main/core/domain/aggregates/basket.go)

[Service](https://github.com/vadiminshakov/dddgo/blob/main/core/services/basketsvc.go)

### Antipatterns:
[Anemic domain model](https://github.com/vadiminshakov/dddgo/blob/main/core/services/antipatterns/anemicBasket.go)

Project is under development.