# SimpleRatingsApp

Keep in mind i am still learning Go and all those packages. So i followed the approach of Vic -- https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j

I have adjusted everthing I needed for my Rating App version. I tried to code it myself and not copy pasting the code snippets.

Also there was a lot of bugfixing to it, to get everything running. I think the biggest issue with the example above was, that the PostgreSQL database is case sensitive and the fields had to be placed into "" e.g. "ID" UUID DEFAULT uuid_generate_v4 ().
But this step was the most important for me, since i had to really understand my code and why those error where occurring.

All in all i had a lot of fun coding this small project, but i think i will not stick to Fiber in following projects.

# Get the App running localy

Before you start, you can change the .env and Makefile files to namings you would like for your project.

Getting everything up and running:
make docker.run // this will run all the commands from Makefile 

# Testing with Postman
If no error occured, you are ready to go for testing with Postman (Postman collection is also present here)
Keep in Mind, that for the private routes e.g. CreateRating you will need to get a jwt token with the Get token request.
Simply copy the token and paste it into "Authorization" -> "Bearer Token" or add a Header.

To look into your DB you could also use a Tool like PgAdmin4 (this is the one i am using).

# Don't forget to delete everthing
After you have finished playing around with this simple app, do not forget to delete the container and images that are still running on docker.

# Have fun testing :)