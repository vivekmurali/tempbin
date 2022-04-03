# TEMPB.IN

[tempb.in](https://tempb.in) is a temporary file storage.

The application uses the following:

- Go
- Postgresql
- HTML/CSS/JS
- Prometheus

Development of this project is complete unless there are some alarming bugs.

# Flow:

On submission of the form, a unique string is created. This string is set as the filename and will be used as the path of the URL that is returned.

A database record is created with the unique string and the actual file name. The file is then stored in a directory called bucket.

A worker scans the database every minute to find files that have _expired_. The files that have expired are deleted and the database record deleted.

# Usage

`git clone https://github.com/vivekmurali/tempbin.git`

Create a database in postgresql _tempbin_ is used here

`psql -d tempbin < db.sql`

`make`

# Deployment

The application is deployed on a VM and the application is run by [systemd](https://systemd.io/). To visualize the prometheus metrics, I have used [grafana](https://grafana.com/).

# Metrics:

A prometheus client counts the number of http requests, the responses, the number of files in the bucket and the duration of the responses.

# Schema

| ID (PK) | NAME | UPLOAD_TIME | URL  | IS_PROTECTED | PASSWORD | IS_LIMIT | LIMIT |
| ------- | ---- | ----------- | ---- | ------------ | -------- | -------- | ----- |
| serial  | text | datetime    | text | bool         | text     | bool     | int   |

# Contributing

All issues are welcome, there's a higher chance of it being fixed if the issues is accompanied with a PR.

# Contact

Vivek Murali - [@vivekmurali2k](https://twitter.com/Vivekmurali2k) - [vivekmurali2k@gmail.com](mailto:vivekmurali2k@gmail.com)

# Acknowledgements

Thanks to:

* [Pico css](https://picocss.com/)
* [pgx](https://github.com/jackc/pgx)
* [go-chi](https://go-chi.io/)
* [gocron](https://github.com/go-co-op/gocron)
* [godotenv](https://github.com/joho/godotenv)

# License

MIT
