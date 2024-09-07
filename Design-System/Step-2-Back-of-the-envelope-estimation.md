# Step 2 Back-of-the-envelope estimation

## 1 - Target

User Number:

- User: 10 Millions user
- Active user: 3 Millions user/day and 10.000 user/second

Follow:

- Each user follow : 200 user
- Special User : 100.000 user

Post:

- Total 10 Millions User, but example only 3 Millions user post per day
- Each User will post 3 posts per day = 30 Millions Posts per day.

===> **350 QPS, max 1500 QPS**

- Each user will spent 5 hour per day to view post ~ 200 post
- 200 post ~ 10 request / hour.
- 3 Millions user \* 10 request = **30 Millions request per hour.**

===> **8.3k QPS, max 40k QPS**
