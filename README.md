# Go Hackernews Scraper

Go Hackernews retrieves the top posts on hacker news by scraping them from the main page.
The application can take an argument of -posts to specify n number of posts.

0 < n <=100


## Installation
This is assuming that docker is installed.

To build the image run the following command
```
Run docker build -t hackernews . 
```
## Usage
The application can be executed by running the following command (assuming the docker image has been built)

```
docker run hackernews --posts n  //n being a positive integer between 1-100
```
This should then output the posts to the terminal:
```
[{"title":"Is this Paypal experience customary?","uri":"https://gist.github.com/stratosgear/6923623321adacd368e965ff193fe2ce","author":"stratosgear","points":109,"comments":45,"rank":1}]
```

## Testing
The application is tested as part of the pre build process with Docker

## Assumptions
* A post is not included if the uri is invalid or if the rank is not >=0 due to feeling it didn't make sense if the rank was incorrect also the same for the uri
* I found that there is a /best page for Hackernews but whenever I tried to go onto another page I received an accessed denied error so I chose to stuck with the main page

## Improvements
* The processFeedItem function isn't tested due to time constraints but would be looked at later.
* Potentially caching the information or storing it somewhere, however this depends how often HN updates its page so might not be useful
* There is also an official HN client so using that might be better than scraping, but for the purpose of the assignment the current implementation is fine.

## Issues
* I've noticed on the odd occasion that HN can block requests and I don't get anything back, so maybe something to retry if failed would fix this.

##Choices
* I have used dep to manage dependencies so that they are easily managed within the project.
* For scraping I chose to use Colly because it provides a pretty clean API into creating a crawler, and it seemed pretty fast. This also uses GoQuery which is useful when traversing the DOM and using jQuery like functions. I used this for the next function to get the next row.
* Additionally I used mockgen because it is useful when mocking interfaces, for example the client didn't really need to test the crawling part so by mocking it I am only  testing the clients functionality
