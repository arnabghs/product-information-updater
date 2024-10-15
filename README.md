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


--------------------------------------
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