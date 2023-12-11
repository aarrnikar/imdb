## imdb
- Analyze data from the IMDB dataset by locally setting up the publicly avaialble data from https://developer.imdb.com/non-commercial-datasets/
- Information courtesy of IMDb (https://www.imdb.com). Used with permission.

### Setting up the environment

#### Database
- clone the repository
- I have used Mysql as the database, any other sql database should be fine as long as those parts are tweaked to match the database
- `--local-infile` option is disabled by default to avoid file imports into the db for security reasons. Since its a dev database we can override the option in the mysql config `my.cnf` by adding the following block

```
[mysqld]
local_infile=1
```

```
chmod +x download.sh
./download.sh -> place all files in dataset in the dataset folder

chmod +x load.sh
./load.sh > load.sql

mysql -u {username} -p {dbname} < create.sql
mysql --local-infile=1 -u {username} -p {dbname} < load.sql
```

#### APIs
- Use `/top/movies` for retrieving top movies and `/top/tv-shows` for top TV shows.
- The `limit` and `numVotes` parameters are optional and can be included in the query string.

### Running the application

- API server
```
From Source:
cd backend
go build
./backend  # runs on port 8000 by default can specify --port flag to change port

# TODO: Dockerize the application
```

#### Interacting with the API server
- cli
```
cd cli
go build
./cli top-movies -l 15 -n 1500 --base-url http://localhost:8000
./cli top-tv-shows -l 15 -n 1500 --base-url http://localhost:8000
or 
./cli top-movies --limit 15 --num-votes 1500 --base-url http://localhost:8000
./cli top-tv-shows --limit 15 --num-votes 1500 --base-url http://localhost:8000
```
