### Option 1

If we want to be **cost effective**, we can go for a combination of **Canary + Rolling Deployment** Strategy.
This mitigates the risk greatly as we can monitor our system with actual production traffic, and based on feedbacks, roll out gradually.
It also ensures High availability and less easier rollbacks only to a set of services.
****

### Option 2

However if budget is not a constraint, I believe a combination of **Canary Deployment** and **Blue-Green Deployment** would be an effective release strategy.

We can reap all the benefits of **canary strategy** along with enhanced safety of a mirrored environment.
If our application is a mission-critical application, and even minimal downtime is unacceptable, this combination works well.


If any issue is found during rollout in Green env(new), instead of gradual rollbacks all users can be routed to Blue env(Old) immediately.

Also for configuration or Infra changes, Canary strategy of testing is not very useful. We would want to update the infra of the entire Green env before testing for performance and sanity.

-----

Whatever Deployment Strategy we may choose, its always wise to combine it with **Feature Toggle strategy** for

- Continuous Delivery
- Trunk Based Developement
- Finer Control on feature release
- A/B testing etc