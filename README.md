# Player Manager

Application to manage player and his items.

## How to run

Start with cloning the repository

```
git clone https://github.com/Vrangz/player-manager.git
```

You can use already prepared scripts in /script directory. To run the application call

```
./scripts/run.sh
```

and if you want to stop the containers then execute

```
./scripts/stop.sh
```

## How to use

The application doesn't allow to create a user, but there's one already created `krzysztofszulcjr`.

The server will be accessible on localhost at 8080 port. Check out [the swagger](http://localhost:8080/api/v1/swagger) to learn more about the API.

### Examples 

Get player information
```
curl "http://localhost:8080/api/v1/players/krzysztofszulcjr"
```

Get player items
```
curl "http://localhost:8080/api/v1/players/krzysztofszulcjr/items"
```
