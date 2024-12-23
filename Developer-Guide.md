## Notes

How to link a module in the project:

1. In your module SOME_MOD/go.mod, add the following:

```
require (
  github.com/hasanaburayyan/raid-bot/SOME_MOD v0.0.0
)

replace github.com/hasanaburayyan/raid-bot/SOME_MOD => /path/to/SOME_MOD
```

## Run Project

```
docker-compose up
```
