# Rikit

Rikit (pronounced *rick-it*) is an API testing CLI for quick and configurable HTTP calls.

Current Version: 1.0.0
Current Release: None

## Set Up

Must have `.rikit.json` file in home directory

the json file should look like:

```
{
  "sites": {
    "google": {
      "url": "http://google.com"
    },
    "digitalocean" : {
      "url": "https://api.digitalocean.com",
      "auth": "Bearer $TOKEN"
    }
  }
}
```

## Usage

### GET

Basic GET request that looks up google in rikit.json file
`rikit get google`

GET request with path flag
`rikit get -p /v2/actions digitalocean`

GET request the enables auth header with auth flag
`rikit get -a -p /v2/actions digitalocean`


## Future Release: MVP

- [ ] Config set up
- [ ] Refactoring of Code
- [x] Making basic GET requests with paths and auths
