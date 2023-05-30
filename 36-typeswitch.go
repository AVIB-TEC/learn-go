package main

import "fmt"

type cloud interface {
	launch() string
}

type aws struct {
	computeSvcName string
}

type azure struct {
	computeSvcName string
}

func (cloud aws) launch() string {
	return fmt.Sprintf("%s launching instance...", cloud.computeSvcName)
}

func (cloud azure) launch() string {
	return fmt.Sprintf("%s launching virtual machine...", cloud.computeSvcName)
}

func compute(cloud interface{}) {
	switch cloudplatform := cloud.(type) {
	case aws:
		aws := cloud.(aws)                                         //casting
		fmt.Printf("AWS: %s -> %s\n", aws, cloudplatform.launch()) //polymorphism
		break
	case azure:
		azure := cloud.(azure)                                         //casting
		fmt.Printf("Azure: %s -> %s\n", azure, cloudplatform.launch()) //polymorphism
		break
	}
}

func main() {
	var clouds []cloud = []cloud{
		aws{"ec2"}, azure{"vm"},
	}

	for _, cloud := range clouds {
		compute(cloud)
	}

}
