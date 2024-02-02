# Brief about my CRUD Chain
Welcome to my CRUD Chain for Problem 5. I decided to create a use case on how this chain might be implemented in the real world. Thus creating a Events Platform inspired CRUD chain. This chain allows creators to create event posts. Those interested in those events can be able to express their interest and the chain will tally the total number of interests for each event. Let go into detail how this is done. 

## Create
Creating an event post by:

```
./crudd tx crud create-event-post '(title)' '(eventType)' '(dateTime)' '(venue)' '(description)' '(interestCount)' --from (User) --chain-id crud
```

## Read
### Specific Post:

```
./crudd q crud show-post 0
```

### List of Posts:

```
./crudd q crud list-post
```

### Further Filtering of Post by eventType:

```
./crudd q crud list-post
```

## Update

```
./crudd tx crud create-update-post '(title)' '(eventType)' '(dateTime)' '(venue)' '(description)' '(interestCount)' (postId) --from (User) --chain-id crud
```

## Delete
```
./crudd tx crud delete-post (postId) --from (User) --chain-id crud 
```

## Additional Feature (Express Interest)
Express interest to event post

```
./crudd tx crud express-interest 0 --from (User) --chain-id crud
```
