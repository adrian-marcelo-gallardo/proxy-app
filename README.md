# Proxy-app

### RUN THE APP AND SEND TEST REQUESTS

In one terminal run:

```bash
go run main.go
```

In another terminal run:

```bash
curl http://localhost:8080/ping -H domain:alpha
```

Values for `domain` header are:
- `alpha` (LOW priority)
- `beta` (MEDIUM priority)
- `omega` (MEDIUM priority)
- `delta` (HIGH priority)

After every request you should get in the response body the sorted list of domains by priority

**NOTE:** if an invalid domain is provided you'll get a response error `invalid-domain`

### RUN UNIT TESTING

In one terminal run:

```bash
go test main_test.go
go test github.com/adrian-marcelo-gallardo/proxy-app/api/models
```