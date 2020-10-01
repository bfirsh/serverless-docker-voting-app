# Serverless Docker Example Voting App

This is a serverless app built with Docker. Read more in the [Serverless Docker repository](https://github.com/bfirsh/serverless-docker).

## Architecture

It consists of a simple entrypoint server that listens for HTTP requests. All of the other functionality of the app is run on-demand as Docker containers for each HTTP request:

 - **vote**: The voting web app, as a CGI container that serves a single HTTP request.
 - **record-vote-task**: A container which processes a vote in the background, run by the vote app.
 - **result**: The result web app, as a CGI container.

## Running

Run in this directory:

    $ make

The app will be running at http://localhost/vote/ and http://localhost/result/ click links to continue..
