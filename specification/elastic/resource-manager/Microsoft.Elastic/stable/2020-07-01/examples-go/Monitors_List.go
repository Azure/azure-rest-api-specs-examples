package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2020-07-01/examples/Monitors_List.json
func ExampleMonitorsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armelastic.NewMonitorsClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("ElasticMonitorResource.ID: %s\n", *v.ID)
		}
	}
}
