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

The name of a method's receiver should be a reflection of its identity; 
often a one or two letter abbreviation of its type suffices (such as "c" or "cl" for "Client"). 
Don't use generic names such as "me", "this" or "self"

Reference: 
https://github.com/golang/go/wiki/CodeReviewComments#receiver-names

