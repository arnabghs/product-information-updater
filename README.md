## About

---
## ADR

[1. Requirements](https://github.com/arnabghs/product-information-updater/blob/main/ADR/1.%20requirements/requirements.md)
<br>
[2. Capacity Estimation](https://github.com/arnabghs/product-information-updater/blob/main/ADR/2.%20capacityEstimation/estimation.md)
<br>
[3. API Design](https://github.com/arnabghs/product-information-updater/blob/main/ADR/3.%20APIDesign/api.md)
<br>
[4. High Level Design]()
<br>
[5. Database Selection](https://github.com/arnabghs/product-information-updater/blob/main/ADR/5.%20databaseSelection/dbSelection.md)
<br>
[6. Data Modelling](https://github.com/arnabghs/product-information-updater/blob/main/ADR/6.%20dataModelling/dataModelling.md)
<br>
[7. Deployment Strategy](https://github.com/arnabghs/product-information-updater/blob/main/ADR/7.%20deploymentStrategy/deployment.md)
<br>
[8. Release Strategy](https://github.com/arnabghs/product-information-updater/blob/main/ADR/8.%20releaseStrategy/release.md)
<br>


---

## Local Infra  Setup

### Bring up Mongo and AWS queues in local
```
make local_setup_up
```

### Setup AWS config
```
aws configure set aws_access_key_id "dummy" --profile test-profile
aws configure set aws_secret_access_key "dummy" --profile test-profile
aws configure set region "eu-central-1" --profile test-profile
aws configure set output "table" --profile test-profile
```

### Create SNS queue
```
aws --endpoint-url=http://localhost:4566 sns create-topic --name order-creation-events --region eu-central-1 --profile test-profile --output table | cat
```

### Create SQS queue
```
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name dummy-queue --profile test-profile --region eu-central-1 --output table | cat
```

### Get SQS ARN
```
aws sqs get-queue-attributes --queue-url  http://sqs.eu-central-1.localhost.localstack.cloud:4566/000000000000/dummy-queue --attribute-names QueueArn --endpoint-url=http://localhost:4566 --profile test-profile --region eu-central-1 --output table | cat
```

### SQS Subscribing SNS Topic
```
aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn   arn:aws:sns:eu-central-1:000000000000:order-creation-events --profile test-profile  --protocol sqs --notification-endpoint arn:aws:sqs:eu-central-1:000000000000:dummy-queue --output table | cat
```
----------------------------------------------
### make sure to add .env file in your local

```
PORT=8080
MONGO_URI=mongodb://localhost:27017
MONGO_DB=local
MONGO_COLLECTION=product
MONGO_USERNAME=root
MONGO_PASSWORD=password123
MONGO_AUTH_SOURCE=admin

AWS_REGION=eu-central-1
AWS_SNS_TOPIC_ARN=arn:aws:sns:eu-central-1:000000000000:order-creation-events
AWS_ACCESS_KEY_ID=dummy
AWS_SECRET_ACCESS_KEY=dummy
SNS_ENDPOINT=http://localhost:4566
```

### Run App
```
make run
```

### Hit endpoint
```
curl --location 'http://localhost:8080/api/v1/products/pid123' \
--header 'Content-Type: application/json' \
--data '{
    "id" : "001",
    "message" : "hello World !!!"
}'
```

### Expected Response
```
{
    "status": "success"
}
```

### mongo DB will get updated
```
[
  {
    "_id": {"$oid": "670e0c0959d5c7278312fde5"},
    "$set": {
      "value": {
        "id": "001",
        "message": "hello World !!!",
        "productID": "pid123"
      }
    }
  }
]
```

### SQS will receive the event from SNS
```
#Read from SQS command

aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url http://localhost:4566/000000000000/dummy-queue --profile test-profile --region eu-central-1 --output json | cat
```

```
#Response

{
    "Messages": [
        {
            "MessageId": "d5ebb586-a7e2-4a4a-b6d8-bea63e29d1ce",
            "ReceiptHandle": "MzU4ZTU0MDctMWU2MC00NzhiLWJkMWItNzNjZjE5NzQwNTVkIGFybjphd3M6c3FzOmV1LWNlbnRyYWwtMTowMDAwMDAwMDAwMDA6ZHVtbXktcXVldWUgZDVlYmI1ODYtYTdlMi00YTRhLWI2ZDgtYmVhNjNlMjlkMWNlIDE3Mjg5NzM4NTMuMjQwNzQ3NQ==",
            "MD5OfBody": "a411d4af2ccbdac8ea71eb28447711af",
            "Body": "{\"Type\": \"Notification\", \"MessageId\": \"4af50809-7373-40eb-97dc-8adb35c87a23\", \"TopicArn\": \"arn:aws:sns:eu-central-1:000000000000:order-creation-events\", \"Message\": \"ID: 001, Message: hello World !!!, ProductID: pid123\", \"Timestamp\": \"2024-10-15T06:30:33.574Z\", \"UnsubscribeURL\": \"http://localhost.localstack.cloud:4566/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:eu-central-1:000000000000:order-creation-events:ef657e91-16be-426e-833b-bd7e2eae4856\", \"SignatureVersion\": \"1\", \"Signature\": \"XWoOo9b+6Lvvo5BCiZbrNfjc0H9NyvS3e+RpO/7dvM1RioZOgCI5IxMHpOAGOhqDmYTktrqTFCi8XgYnarReUHR+G5kDEGVlqxfTgJ34Sp28wVfrOKOywErd9DloaGtPK+T3ik9rl8wj4i9whz82CTOKKe76o1wutpvU/i1Mi/oEcxkMeQukmxPJm1ikF08Gr9MVMaDU4tcNvm49sAN+9y+zsZbUuuPMhtLIADRatZfpyl/OERT0PqMpY7jp618QycSgxzi1dsZ4p54TP4U/luMMoM/TVmwZknEj/iGJOkDAbxN7HTFsSH8PO9Fg/pzOw91hAs0UdqoyUMgSAzcn7w==\", \"SigningCertURL\": \"http://localhost.localstack.cloud:4566/_aws/sns/SimpleNotificationService-6c6f63616c737461636b69736e696365.pem\"}"
        }
    ]
}
```

----------------------------------------------
### Helper commands

### List ARNs
```
aws --endpoint-url=http://localhost:4566 sns list-topics --region eu-central-1 --profile test-profile --output table | cat
```

send events to SNS topic
```
aws sns publish --endpoint-url=http://localhost:4566 --topic-arn arn:aws:sns:eu-central-1:000000000000:order-creation-events --message "Hello World" --profile test-profile --region eu-central-1 --output json | cat
```

receive events in SQS
```
aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url http://localhost:4566/000000000000/dummy-queue --profile test-profile --region eu-central-1 --output json | cat
```

delete events in SQS
```
aws sqs delete-message --endpoint-url=http://localhost:4566 --queue-url http://localhost:4566/000000000000/dummy-queue --profile test-profile --region eu-central-1  --receipt-handle <message-handle>
```

send events to SQS queue
```
aws --endpoint-url=http://localhost:4566 sqs send-message  --queue-url http://localhost:4566/000000000000/dummy-queue --profile test-profile --region eu-central-1  --message-body '{
          "event_id": "7456c8ee-949d-4100-a0c6-6ae8e581ae15",
          "event_time": "2019-11-26T16:00:47Z",
          "data": {
            "test": 83411
        }
      }' | cat
```
