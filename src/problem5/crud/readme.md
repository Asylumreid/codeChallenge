# Implementing Consensus Breaking

My approach to this is by trying to hard fork my chain by modifying some versions of my App.go files which would allow one chain to run the old version of the consensus protocol and state while the other chain runs the new version, incompatible with the previous one. With this theory, I tried to modify the version however, was not able to get a clear consensus-breaking introduction into my chain. Would still need more time in research to be able to fully understand the tender mint consensus to implement another method to introduce this change.



# Brief about my CRUD Chain
Welcome to my CRUD Chain for Problem 5. I decided to create a use case on how this chain might be implemented in the real world. Thus creating an Events Platform-inspired CRUD chain. This chain allows creators to create event posts. Those interested in those events can be able to express their interest and the chain will tally the total number of interests for each event. Let's go into detail how this is done. 

## Create
Creating an event post by:

```
./crudd tx crud create-event-post '(title)' '(eventType)' '(dateTime)' '(venue)' '(description)' '(interestCount)' --from (User) --chain-id crud
```
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/e130cbce-6c61-4642-affa-463b74ae70d7)

## Read
### Specific Post:

```
./crudd q crud show-post 0
```
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/debaea52-c152-46ad-9c4d-0487189b3aa6)

### List of Posts:

```
./crudd q crud list-post
```
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/be94b2f1-c147-41d1-91e5-1d357f5271f7)

### Further Filtering of Post by eventType:

```
./crudd q crud list-post
```
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/8f134a90-4fb1-4f3f-8103-b569cb0eb78c)

## Update

```
./crudd tx crud create-update-post '(title)' '(eventType)' '(dateTime)' '(venue)' '(description)' '(interestCount)' (postId) --from (User) --chain-id crud
```
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/5a9e15cb-3cac-42e3-93b4-177b748b124d)
## Delete
```
./crudd tx crud delete-post (postId) --from (User) --chain-id crud 
```
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/f9e15130-e9a2-47e9-bb12-55da8b6d73ab)
## Additional Feature (Express Interest)
Express interest to event post
```
./crudd tx crud express-interest 0 --from (User) --chain-id crud
```
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/d9f2ab00-48a9-4496-8fcc-6d63731c4aaf)
![image](https://github.com/Asylumreid/codeChallenge/assets/114651163/9f2d9bc2-2d4e-428a-a50d-f330a2030b09)
