# event-producer
event-prducer service publishes events to a kafka topic. There is a POST API exposed which takes in a request body and sends the message to kafka. 
That message is then consumed by event-consumer service to determine the amount to be paid for the fuel.

## POST API details
URL : /api/v1/event
Request body: {
  "fuellid": true,
  "city": "bangalore,
}

## Environment Variables required
  Variable Name: BROKER
  Variable Description: Required to get the broker url of kafka
  
  Variable Name: TOPIC
  Variable Description: Required to get the kafka topic name
  
  Variable Name: USERNAME
  Variable Description: Required to get the username for basic authentication
  
  Variable Name: PASSWORD
  Variable Description: Required to get the password for basic authentication
  
## Installation
1. Clone the repository in your GOPATH
2. Install kafka, start the zookeeper, and start the broker.
3. Execute the below commands :
    NOTE: replace the parameters present in <> with actual values.
    i.  set BROKER=<broker_url>
    ii. set TOPIC=<topic_name>
    iii.set USERNAME=<any_username>
    iv. set PASSWORD=<any_password>
4. Navigate to the repository's directory in your local and run the command
    go run main.go
  The application would start running at port number 8091.
