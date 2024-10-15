## Daily/Monthly Active Users :

Depends on the pricing managers


## Throughput:

Assuming we make 100,000 product updates each day,

**write request =** 0.1 million/day

Assuming we check 20% of the product information updates

**read request =** 20,000 per day

## Storage:
Assuming each request is 2 KB, hence

total storage needed in a day = 2KB * 0.1 million = 0.2 million KB = 200MB/day

total storage needed in 10 years = 200 MB * 365 * 10 = **730 GB**

## Cache Memory

Assuming required Cache to be 1% of total daily storage,
200 MB * 0.01 = 2 MB/day

## Bandwidth

**Data flow into our system(Ingress)**

Incoming data in a day = 200 MB

Incoming data per second = 200 MB / (24* 60*60) = **2.5 KB/sec**

<br />

**Data flow out of our system(Egress)**

Outgoing data in a day = 20,000 read req * 2 KB (each request)

Outgoing data per second = (20,000 * 2) / (24* 60*60) = **0.5 KB /sec**