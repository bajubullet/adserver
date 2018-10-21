# adserver
Ad server for serving Ads


## Usage
- Clone the repo into gopath.
- go install
- $GOPATH/bin/adserver <csv_file>
- Re-execute the last command when you get a new file.

### What if the .csv file is very big (billions of entries) - how would your application perform?
If csv file is very big then it will take a lot of time to start the server and the memory usage will be high but once its up it should work fine.

### How would your application perform in peak periods (millions of request per minute)?
Right now it should work fine with thousands of requests but for millions we'll have to setup some caching and load balancing.

### Every new file is immutable, that is, you should erase and write the whole storage.
Right now the program takes care of it since the data stays in memory itself and since we are re-executing the command everytime we get a new file it will reset the memory.


## Extras
We can improve the program by separating the API server and file parsing and using an external cache or database.

So that API has 100 % uptime and whenever a new file comes we can execute the parsing binary with the filename to populate the DB/cache.

Ignored this approach for now because it will add a lot of things to the setup process. But I have kept the functions seprate so its easy to do it when we decide to do so.
