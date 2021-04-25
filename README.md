# Self-hosted HIBP password hash check only

This is an example of self-hosted pwned password list API obeying the HiBP k-anonymity setting.

## Download the passwords SHA-1 by count file

Download the SHA-1 7zip archive from HiBP website. The most up to date file can be downloaded from https://haveibeenpwned.com/Passwords. The compressed file is 12.5GB, 26GB decompressed. The file is a text file with each line in the following format:

```
SHA-SUM-41-CHARS-LONG:INT-COUNT
```

To decompress the file:

```sh
sudo apt-get install p7zip-full
p7zip -d pwned-passwords-sha1-ordered-by-count-v7.7z
```

## Build the Docker image:

```
make docker-build
```

## Start the example Docker compose environment

```sh
cd compose/
docker-compose -f compose.yml up
```
