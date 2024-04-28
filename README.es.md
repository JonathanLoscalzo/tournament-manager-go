# Administrador Torneos

## Architecture flow

<!-- https://github.com/sklinkert/go-ddd/blob/master/internal/application/interfaces/product_service.go -->

application: 
    - use cases interactions with repositories and external services (all injected). Dependency with domain
    
domain: 
    - entities that interact

infrastructure: 
    - implementation of repositories
    - database entities
    - gateways or mappers between domain and database entities

core: 
    interfaces
        - usecases
        - repository
        - service

flow:
cmd -> build dependencies (external services, repositories, usecases, api) -> application -> domain 

### Fútbol
Hay ciertos términos que son importantes tener en cuenta, algunos pueden ser almacenados y otros calculados

## Términos Clave

- Competitions 
    - Season
        - Fixture
            - Matches
                - LineUps
            - Groups
        - Standings
            

- Stats
    - Top Scorers
    - Top Assists
    - Players Stats
    - Teams Stats
    - Fouls
    - Head2Head's
    ...

- Transfers (no en principio)

- Odds/Apuestas (1x2, BTS, O/U, AH) (no en principio)

Otros Deportes: Handball, Basquet, Voley, Tenis, Paddle

## Referencias

-  https://clickn.run/es/blog/tournament-system-types
    - eliminación
        - simple
        - en cuenta casos pares e impares (uno queda flotante)
        ...
    - liga
    - grupos(liga) + eliminación(copa)


- ORMS
    - gorm no tiene table inheritance fácil (para MatchEvent)
    - https://bun.uptrace.dev/guide/models.html#embedding-structs