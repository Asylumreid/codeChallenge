
![Broadcast Transaction Service - Page 1](https://github.com/Asylumreid/codeChallenge/assets/114651163/df112ae4-6db0-4b66-b81d-3d0c21080056)

Based on the Diagram

We have the relevant key components below:
API Gateway:
This component is responsible for receiving and validating incoming POST /transaction/broadcast requests. This includes basic parameter validation like the message type and the data format in the request. Once validation is completed, it will route the request payload to the Transaction Processor. 

Transaction Processor:
This component is responsible for the verification and signing of the transaction. The processor first verifies the message type against a predefined list of options such as add_weight in the example. Once verified, it will sign the transaction data, which constructs a transaction object containing the signed data. It then passes the signed transaction to the Broadcaster to be validated by nodes.
Broadcaster:
This component is responsible for interacting with the blockchain to broadcast the transaction and handle responses from the blockchain. The Broadcaster established a connection to the blockchain node using the RPC endpoint. It sends the transaction object to the node via the RPC interface. It handles the Blockchain node’s response based on:

Success (HTTP 200):
-	Retrieves the transaction receipt containing details such as the transaction block number, gas fee used, etc.
-	Update the transaction status in the Database manager to “Success” and store the receipt.
-	Returns HTTP 200 to the API Gateway indicating success.
Failure (HTTP 4XX – 5XX):
-	Since it’s 95% within 20-30 seconds, implement a Linear Backoff Retries with an increased delay of 20, 30, 40 seconds and a maximum of three retries.
-	After maximum retries are reached, it updates the transaction status in the Database manager from pending to “failed” and records the error details. 
-	Returns the specific HTTP error code (HTTP 400 for invalid transaction, HTTP 500 for internal server error) to the API Gateway.
Database Storage Manager: 
This component is responsible for storing the transaction details for persistent data storage. The information stored would include the Transaction ID (uid), message type, data, signed transaction status, broadcast status and the blockchain receipt. There will also be methods provided to the User interface to retrieve the transaction details.

User Interface: 
This component is responsible for retrieving transaction details from the Database Storage Manager. GET request methods will be implemented to communicate and retrieve transaction details from the Database Storage Manager. There will also be implementation for filtering and sorting the transaction details based on criteria. The user interface will also provide a method whereby it  allows the admin the option to retry a failed broadcast at any point in time.

