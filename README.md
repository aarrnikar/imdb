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
