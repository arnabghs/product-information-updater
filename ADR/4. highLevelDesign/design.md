## Option 1: Cloud Native Infrastructure

Pros
1. Implement our system in a cloud platform helps leverage the autoscaling and
     fault tolerance behaviours clouds provides by default


2. Easy and quick to setup


3. No additional maintenance effort

Cons
1. Costly Setup for cloud services


2. Higer Latency ( instead of directly On-prem to Downstream connection, one extra hop in between)

Notes
1. AWS is chosen simply because I am comfortable in AWS terminologies compared to AZURE or GCP.
   <br> However its better to go with the cloud provider we are already using extensively for easier maintainance, cost effectiveness and low latency


2. For authentication, we can attach a authentication/authorization lambda with API Gateway. we can also leverage AWS cognito service


3. The SQS makes sure that the message will not get lost and retry can happen.
   If direct connection to lambda was there, and processing failed message would get lost. But Queue makes sure to retry and then put in DLQ

<br>
Diagram 1: Using VPN for multi/hybrid cloud connection
<br>
 <img width="1350" alt="1" src="https://github.com/user-attachments/assets/cb01d138-fa2e-4bc0-bd65-f7c1e3fe3d93">
<br>
#### We Can replace Private VPN Tunnels with HTTPs calls to reduce cost and additional VPN setup
However it'll increase latency and will be less secured
<br>

Diagram 2: Using public Internet for multi/hybrid cloud connection
<br>
<img width="1161" alt="2" src="https://github.com/user-attachments/assets/55f1d58d-834c-4842-a8be-e0917c8c3678">




--------------


## Option 2: On-premise Infrastructure
I was not entirely sure what is the scope of source system or on-prem data centre , hence created this design.

Pros
1. Less Costly since we don't need to pay for cloud services


2. Lower Latency, On-prem server is directly connected to downstream services

Cons

1. More maintenance needed for the entire infra


2. Carefully setup required to make the system scalable and fault tolerant

<img width="1277" alt="3" src="https://github.com/user-attachments/assets/924a457b-2a95-41a1-8a73-05a61690dd37">
