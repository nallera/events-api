# Events API
### Implementation
The solution is implemented in Golang 1.22, using hexagonal architecture. There's a configuration _yaml_ file, a Makefile with different targets, and a SQL Database file (_events.db_). I decided to call the provider **Provider X** just for the sake of the exercise.  

The only exposed endpoint is under the _reader_ handler, and it's mainly composed of five steps:

- Input validation
- Attempt to get current events from the provider endpoint
- Update the database with the current events if any (insert or update)
- Query the database for events in the required time range
- Build the response

I decided to use a SQL database to store the events because querying is a good resource to be able to filter events based on a date range. I chose to use SQLite as it is a good option for local projects. The app loads the _events.db_ database, or creates it if it doesn't exist, and creates the table _events_. The table has a double primary key, based on the event and the base event IDs, as sometimes two different events have the same event ID even though they are associated to different base events.

### Execution
You should first download the project dependencies by either:
- Run in terminal (from project root):
```
go mod tidy
```
- Using the _tidy_ target of the **Makefile**, by running in terminal (from project root):
```
make tidy
```

Then you can run the app by either:

- Run in terminal (from project root):
```
go run cmd/app/main.go
```
- Using the _run_ target of the **Makefile**, by running in terminal (from project root):
```
make run
```

You can run the tests by either:
- Run in terminal (from project root):
```
go test -v ./...
```
- Using the _test_ target of the **Makefile**, by running in terminal (from project root):
```
make test
```

### Using the app
The _search_ endpoint curl is:
```
curl --location 'http://localhost:8080/search?starts_at=2021-01-01T17:32:28Z&ends_at=2021-12-01T17:32:28Z'
```
And a successful _200_ status code response example:
```
{
    "data": {
        "events": [
            {
                "id": "B291-291",
                "title": "Camela en concierto",
                "start_date": "2021-06-30",
                "start_time": "21:00:00",
                "end_date": "2021-06-30",
                "end_time": "21:30:00",
                "min_price": 15,
                "max_price": 30
            },
            {
                "id": "B1591-1642",
                "title": "Los Morancos",
                "start_date": "2021-07-31",
                "start_time": "20:00:00",
                "end_date": "2021-07-31",
                "end_time": "21:00:00",
                "min_price": 65,
                "max_price": 75
            }
        ]
    },
    "error": null
}
```
 