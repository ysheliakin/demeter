# Demeter: Local Food Donation Platform

![alt text](wwwroot/demeter.png "Demeter App Logo")

>In ancient Greek religion and mythology, Demeter (/dɪˈmiːtər/) is the Olympian goddess of the harvest and agriculture, presiding over crops, grains, food, and the fertility of the earth.
>Source: https://en.wikipedia.org/wiki/Demeter

The website aims to connect individuals and organizations willing to donate food with individuals in need and organizations such as shelters and food banks. The idea is to redistribute excess food to minimize waste and ensure surplus food reaches those in need rather than being discarded.

## Development

1. Install Go: https://go.dev/doc/install  
  (optional) Install [Air](https://github.com/cosmtrek/air?tab=readme-ov-file#installation) to track code changes on the fly. Might need to add $GOPATH/bin in your PATH somewhat like [so](https://stackoverflow.com/questions/70098792/go-install-do-i-need-to-manually-update-my-path) or [so](https://github.com/golang/go/issues/18583).

2. Install [goose](https://github.com/pressly/goose) for database migration management.

2. If using Air, run in the project directory:

  ```bash
  export DB="<db connection string>"; 
  export IMAGEKIT_API_KEY="<private image hosting key>"; 
  air
  ```

  - Or just run (need to restart the server any time you make code changes):

  ```bash
  export DB="<db connection string>"; 
  go run *.go
  ```

3. Enjoy the website at localhost:42069. 

4. Pushes into `main` branch also result in deployments to [`demeter.adaptable.app`](https://demeter.adaptable.app/).

### Database Migrations

To create an empty SQL migration file (run from project root directory):

```bash
goose -dir ./db/migrations create [migration-name] sql
```

To migrate the database (this will make changes to the cloud database instance):

```bash
goose -dir ./db/migrations postgres $DB up

goose -dir ./db/migrations postgres $DB down
```

### Database Queries

To generate Go code based on scheme and queries, go to `./db` directory, add/update queries, then run (still in `./db`):

```bash
sqlc generate
```
