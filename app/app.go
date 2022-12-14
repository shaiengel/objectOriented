package app

import (
	"fmt"
	"ooo/model"
)

func Oop() {

	IKeyGenerator := model.NewKeyGenerator(nil, nil)
	IKeyGenerator.ProcessPeriodList(nil, nil)
	keyGenerator := IKeyGenerator.GetKeyGenerator()
	IKeyGenerator.Enhance()

	if (*keyGenerator.PeriodList)[0] == "yes" {
		fmt.Print("SUCCESS: upcasting works for pointers which didn't change\n")
	} else {
		fmt.Print("it doesn't work\n")
	}
	if keyGenerator.Misc == "vod" {
		fmt.Print("SUCCESS: upcast doesn't work for params")
	} else {
		fmt.Print("strange")
	}
}
