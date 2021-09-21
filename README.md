# KeyStat

This repo provides a solution to finding both streaks and 'key stats' within data in Golang. 
An example of a key stat is "Manchester United have won 5 out of their last 7 games", with the 
data supplied being binary win (or not) values and supplying the minimum and maximum number of recent 
games to focus on, as well as what a "meaningful cutoff" for this stat to be deemed key is; ie if 
Man U had won 2 of 10 (20%) and the supplied cutoff was 50% this would not trigger as below the 
threshold.
 
## Docs

*TODO*


## Usage

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
	so.Streak() // => 3.0

	kso := NewKeyStatObject(5, 10, 0.6)
	for _, f := range samples {
		kso.Add(f)
	}
	kso.KeyStat() // => true, 7, 10
}
```

## Contributing

Currently developing a release process so this will be fleshed out with greater detail in the future, 
currently we are accepting pull requests for minor fixes, etc:

* Small bug fixes
* Typos
* Documentation or comments

Feel free to open issues to discuss new features.

## License

This repository is Copyright (c) 2021 Sam Watson. All rights reserved.
It is licensed under the MIT license. Please see the LICENSE file for applicable license terms.
