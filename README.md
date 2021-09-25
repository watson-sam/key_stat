# KeyStat

This repo provides a solution to finding a number of 'key stats' within data in Golang.

# Getting Started

## Installing

To start using KeyStat, install Go and run `go get`:

```sh
$ go get -u github.com/watson-sam/key_stat
```

This will retrieve the library.

# Documentation
For each function offered I will try and flesh out a bit both what I am trying to 
achieve with the code as well as give an example use case; these are entirely in 
as sporting context as that is the original point of this library although they 
should be fairly universal in terms of what they achieve and hopefully the sports 
context will help make the aims easier to understand.

### Streaks 
_Definition_

For each supplied points how many have been consecutively 1; supplied with either a 1 
or 0 'streaks' can be found, for each additional 1 supplied the streak value increases
by 1 and is reset to 0 when the supplied value is 0.

If the data supplied is the outcome of a football game with a 1 representing if a team has
won and 0 a team has not won, with a streak value of 5 one could say "Team X has won 5 on 
the bounce".

_Usage_
```go
package main

import "github.com/watson-sam/key_stat"

func main() {
	var samples = [10]float64{
		1, 1, 1, 1, 0, 0, 0, 1, 1, 1,
	}

	so := NewStreakObject()
	for _, f := range samples {
		so.Add(f)
	}
	so.Value() // => 3.0
}
```

### Historic 
_Definition_

The historic method is used to capture the extremities of the data supplied, specifically the 
historic maximum or minimum values given, this method also includes the functionality to compare
a given value with the current min or max with an outcome of 0 meaning supplied value is not 
equal or more extreme (large or smaller given the what stat was requested when setting up), 1 means 
supplied value is equal to the most extreme and 2 means that supplied value is the most extreme.

If the data supplied are the historic scores of a football team, if the current stored value is 5 and 
today they scored 6, when querying the Compare method we would have 2 returned and one could say "Team X 
has scored 6 goals, this is the most in their entire history".

_Usage_
```go
package main

import "github.com/watson-sam/key_stat"

func main() {
	var samples = [10]float64{
		10, 10, 6, 8, 7, 7, 7, 10, 6, 20,
	}

	ho := NewHistoricObject("min")
	for _, f := range samples {
		ho.Add(f)
	}
	ho.Compare(5) // => 2
}
```
### Key Stat 
_Definition_
A key stat is aiming to find 'meaningful' stats about supplied data; the method looks over all points 
in the supplied window (starting with the maximum set with maxWindow and decreasing to a limit of 
minWindow if no meaning is found), what a meaningful stat is given with the cutoff which is the minimum 
percentage of points in the window that need to be 1. I appreciate this is frankly a terrible definition 
so hopefully the clouds will clear and the path to meaning will be found with the example.

An example of a key stat could be "Team X have won 5 out of their last 7 games", with the 
data supplied being binary win (a supplied value of 1) or not (a supplied value of 0) values. With a cutoff of 
0.7 and a min and max windows of 5 and 10 respectively.
 
_Usage_
```go
package main

import "github.com/watson-sam/key_stat"

func main() {
	var samples = [10]float64{
		1, 1, 1, 1, 0, 0, 0, 1, 1, 1,
	}

	kso := NewKeyStatObject(5, 10, 0.6)
	for _, f := range samples {
		kso.Add(f)
	}
	kso.KeyStat() // => true, 7, 10
}
```
# Contributing

Currently developing a release process so this will be fleshed out with greater detail in the future, 
currently we are accepting pull requests for minor fixes, etc:

* Small bug fixes
* Typos
* Documentation or comments

Feel free to open issues to discuss new features.

# License

This repository is Copyright (c) 2021 Sam Watson. All rights reserved.
It is licensed under the MIT license. Please see the LICENSE file for applicable license terms.
