package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	// this will show an in-place diff
	runDiff(false)
	// this will show a complete object replacement
	runDiff(true)
}

func runDiff(deployDependency bool) {
	// difference in properties is exactly 3, unless deployDependency (or any other property!)
	// is changed

	// src -> src / src2 (change 1)
	// dst -> dst / dst2 (change 2)
	// dstVersion -> v1 / v2 (change 3)
	// optionally changing DeployDependency based on input value

	val := []*Parent{
		{
			Id: "app",
			Slice: []Child{
				{
					Src:        "src",
					Dst:        "dst",
					DstVersion: "v1",
					Routes: []string{
						"test1",
						"test2",
					},
					DeployDependency: deployDependency,
				},
			},
		},
	}

	compareVal := []*Parent{
		{
			Id: "app",
			Slice: []Child{
				{
					Src:        "src2",
					Dst:        "dst2",
					DstVersion: "v2",
					Routes: []string{
						"test1",
						"test2",
					},
					DeployDependency: false,
				},
			},
		},
	}

	fmt.Println(cmp.Diff(val, compareVal))
}
