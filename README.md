# This repo is simply provide some golang programming guilde

## Best practice in Go

### Naming

####A good name is:
* Consistent (easy to guess)
* Short (easy to type)
* Accurate (easy to understand)

The greater the distance between a name's declaration and its uses, 
the longer the name should be

https://talks.golang.org/2014/names.slide#1


#### Receiver Name

**Don't use generic names such as "me", "this" or "self"**

* The name of a method's receiver should be a reflection of its identity; 
* Often a one or two letter abbreviation of its type suffices (such as "c" or "cl" for "Client"). 
* The name need not be as descriptive as that of a method argument, as its role is obvious and serves no documentary purpose.
* Be consistent, too: if you call the receiver "c" in one method, don't call it "cl" in another.

Reference: 
https://github.com/golang/go/wiki/CodeReviewComments#receiver-names

