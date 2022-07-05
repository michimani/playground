package main

import (
	"context"
	"fmt"
)

func main() {
	evidentry, err := NewEvidentlyClient(&NewClientInput{
		Context: context.Background(),
		Region:  "ap-northeast-1",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	out, err := evidentry.ListFeatures(context.Background(), "sample")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(len(out.Features))
	if out.NextToken != nil {
		fmt.Println(*out.NextToken)
	}
}
