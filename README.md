Implements some simple string manipulation and string search problem solutions using stack:

* Golang API (Gin)
* Elasticsearch v7

## Problem 1

A string is said to be a special string if either of two conditions is met:

1. All of the characters are the same, e.g. aaa.
2. All characters except the middle one are the same, e.g. aadaa.

A special substring is any substring of a string which meets one of those criteria. Given a string, determine how many special substrings can be formed from it.

### Solution 1

Expose API endpoint:

POST /special/count

which consumes data like:

```json
{
    "data": {
        "attributes": {
            "text": "aaaaaabbaaa"
        }
    }
}
```

and responds with the count of special substrings as defined above.

### Use

1. Clone repo
2. From project root run `docker-compose up`
3. Run tests `docker exec -it app go test ./...`
4. Make special strings request count:

```
curl -X POST \
http://localhost:8080/special/count \
-H 'Content-Type: application/json' \
-d '{
    "data": {
        "attributes": {
            "text": "alskdfjalsiddjfoirjkdddddddddddddddddddjijiiirrrrrrrrrrrrrrrrrrrrlllllllllllllllllllldkaldjifaddjddddddddddddddddddddjddddddddddddddddkkkkalsdlsksjlsajdlkfjaldskjfalsdkiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnssssssssssssssss"
        }
    }
}'
```

Should return the count of special substrings 6,033.
