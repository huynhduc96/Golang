# Step 7 Identifying and resolving bottlenecks

## 1- Problem with people with many followers:

- when a person with many followers posts, it may take a long time to push data to the followers' newsfeeds, and cause the system to be congested

==> so we will divide the user group over 1 million followers separately

and not process the newsfeed update for the stopped person

- The newsfeed will be a merge between the cached newsfeed + posts of the hot users that the person is following

- Ensure that the web server and app server can scale horizontally.

- Index fields like created_at, visible for the fastest query.

- Take advantage of requests to redis to be as multiple get as possible to optimize the network
