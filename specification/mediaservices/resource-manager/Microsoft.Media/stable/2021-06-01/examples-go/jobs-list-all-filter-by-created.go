package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/jobs-list-all-filter-by-created.json
func ExampleJobsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewJobsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		&armmediaservices.JobsClientListOptions{Filter: to.StringPtr("<filter>"),
			Orderby: to.StringPtr("<orderby>"),
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
