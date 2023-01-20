## Notes_app is just a warm up that will show us how we can scale applications to next stage

## Functional requirement

    # User can signup on the platform
    # User can login on the platform
    # User can search through availables blogs posted by users
    # User while signup , choose their interested labels.
    # each blog is attached to a label.
    # User can have a recommended section where they'll get blogs according to their interest
    # User can update their profile
    # user can follow other blog creators
    # Users can Unfollow

## Non - Functional Requirement

    # High Availability
    # System should be able to server requests
    # we can have low consistency as little latency between update on the blogs update is ok

## Steps to run the project

- Make sure you have docker installed
  `check by using command docker --version`

- Make sure you have IP whitelisted for database access

- Build the docker image
  `docker build -t notes_app .`

- Run the app
  `docker run -p 9200:9200 notes_app`
