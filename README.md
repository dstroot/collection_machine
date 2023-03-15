# Collection Machine

This program is designed to drive the collections process for unfunded tax returns.  

Processing outline
------------------
* get UNFUNDED_ITEMS
* process the records
    * determine next process step
    * perform the next process step
    * Log any exceptions

Testing
-------
    $ go test -v
    $ go test -bench=.
    $ go test -cover

Compiling
---------
```
    $ go build collection_machine.go
	$ go build && ./collection_machine && rm -rf collection_machine
```

(Will produce an executable `collection_machine`)

Running
-------

    $ ./collection_machine

Installing
---------

    $ go install collection_machine.go

then run `collection_machine`.  

Docker
------

All you have to do is run `./docker_build.sh` and boom!

To run the docker image locally:

    `$ docker run --rm -it -p 8000:8000 dstroot/collection_machine`


TODO
----
[X] Add transactions for database update/insert activity
[X] Switch configs to .env file 
[X] Vendor all machines
[X] Need to create a means to create the bcrypt hash once due to the computational load
     - Need a place to store it - add URL column to COLLECTION_ITEM
     - Need to select it in as part of the records to process
     - Need to identify if it is step one and create the URL for the first time, or just use the existing one
     - Can remove it from the email queue?
[X] Change SQL to only pull back items that are ready for processing by looking at timestamp
[X] Need a way to "defer" an autodebit per Intuit's request.  Deferring is conceptually simple in that we just need a deferral step in the process definition that defines the wait period. The challenges are:
	- We don't know at what point in the process (after the first email? after the second?) when the person is deferred so we don't exactly know where to jump back to. We can obviously figure it out by ordering the activities by timestamp but that is additional processing.  
	- We don't know exactly when the next autodebit cycle is in relation to when the deferral takes place. We know that we like autodebits to happen on Fridays (payday).
	- If we move deferral processing to a higher number in the process definition can rejects still be handled as they are today (at lower processing numbers)? Yes, this has been tested and verified.  
    - Today, the total elapsed cycle time takes 24 days from start on a Tuesday to autodebit hitting bank account on a Friday. If we create a deferral period of 21 days, and someone is deferred immediatly after the first email, they actually may get autodebited *earlier* than if they just go through the normal cycle.
	[X] Added steps 40-42 to process 10
	[X] Create new deferral email
	[X] Verify that rejects can go back in as they do today because they are at a lower process number
[X] Add processing for Taxslayer
	[X] Add new 30 and 40 processes to the process type table for Taxslayer
	[X] Add new process definition steps for processes 30 and 40 for Taxslayer
	[X] Change code to recognize multiple autodebit and non-autodebit (direct mail) cycles.  
