# Self-hosted HIBP password hash check only

This is an example of self-hosted HiBP pwned password list API with the k-anonymity setting.

This code serves primarily as an example for setting up [self-hosted HiBP for Ory Kratos](https://github.com/ory/kratos/pull/1009#issuecomment-826372061) but there is nothing preventing it from running in a completely local environment where the access to the public HiBP API would be prevented but password checks are required.

## Supported databases

- PostgreSQL 13 is the only database this has been tested with and PostgreSQL is the only database supported

## Download the passwords SHA-1 ordered by count file

Download the SHA-1 7zip archive from HiBP website. The most up to date file can be downloaded from https://haveibeenpwned.com/Passwords. The compressed file is 12.5GB, 26GB decompressed. The file is a text file with each line in the following format:

```
SHA-SUM-41-CHARS-LONG:INT-COUNT
```

On Ubuntu (requires ~40GB free space in `/tmp`):

```sh
sudo apt-get install p7zip-full
cd /tmp
wget https://downloads.pwnedpasswords.com/passwords/pwned-passwords-sha1-ordered-by-count-v7.7z
p7zip -d pwned-passwords-sha1-ordered-by-count-v7.7z
```

## Build the Docker image

```
make docker-build
```

A `localhost/hibp:latest` Docker image will be created.

## Start the example Docker compose environment

```sh
cd examples/compose/
docker-compose -f compose.yml up
```

## Create the table

By following this readme, the Docker compose setup creates a network called `compose_hibpexample`. Hence, in another terminal:

```sh
docker run --rm \
    --net=compose_hibpexample \
    localhost/hibp:latest \
    migrate --dsn=postgres://hibp:hibp@postgres:5432/hibp?sslmode=disable
```

This command should exit without any output. No output means it executed okay.

### Import the data

This will take some time, there are over 613 million lines in the V7 file. Here, I'm using the `/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt` file on the host and `/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt` in the container:

```sh
docker run --rm \
    -net=compose_hibpexample \
    -v=/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt:/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt \
    localhost/hibp:latest \
    data-import --dsn=postgres://hibp:hibp@postgres:5432/hibp?sslmode=disable \
      --password-file=/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt
```

For testing, you can import X first lines using the `--first=X` flag, line this:

```sh
docker run --rm \
    --net=compose_hibpexample \
    -v=/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt:/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt \
    localhost/hibp:latest \
    data-import --dsn=postgres://hibp:hibp@postgres:5432/hibp?sslmode=disable \
      --password-file=/tmp/pwned-passwords-sha1-ordered-by-count-v7.txt \
      --first=10000
```

## Test

If you import the data with `--first=X` flag, pick a hash from first X lines of the file, for example, for verion 7:

```sh
less /tmp/pwned-passwords-sha1-ordered-by-count-v7.txt
```

gives this in the first line:

```
7C4A8D09CA3762AF61E59520943DC26494F8941B:24230577
```

The `:prefix` in the `GET /range/:prefix` call is the first 5 characters of the hash, to query, execute:

```sh
curl http://localhost:15000/range/7C4A8
```

## Setting up behind reverse proxy with TLS

TODO

## License

Apache-2.0 License
