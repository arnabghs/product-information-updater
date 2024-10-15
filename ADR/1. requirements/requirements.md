**Functional:**

- update product information promptly through events in a number of downstream systems that is deployed in different cloud providers in different regions


- simple visualisation/dashboard about product information updates [Good to have]



**Non-Functional:**

| Scalability | The platform should be auto-scalable based on compute capacity or processing backlog |
| --- | --- |
| Consistency | Ensure **eventual data consistency**, even if not instantly but the product information should get updated eventually across all regions |
| Availability | Our system should be highly available so that theres not delay in price updates and don’t affect the business forecasts |
| Resiliency | Minimize downtime |
| Low Latency | Since our system is **globally distributed**, we should take into considerations how to best improve the latency.
| Cost Efficiency | The design should be cost effective |