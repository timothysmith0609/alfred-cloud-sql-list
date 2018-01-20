package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

var (
	project = flag.String("project", "", "target GCP project")
)

func main() {
	flag.Parse()
	if *project == "" {
		log.Fatal("You must specify a project")
	}

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, sqladmin.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	sqlAdminService, err := sqladmin.New(c)
	if err != nil {
		log.Fatal(err)
	}

	results, err := sqlAdminService.Instances.List(*project).Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range results.Items {
		fmt.Printf("%s %s\n", i.Name, i.InstanceType)
	}

}
