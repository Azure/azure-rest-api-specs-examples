package armdomainservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListOuContainers.json
func ExampleOuContainerClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewOuContainerClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("OuContainerResourceGroup",
		"OuContainer.com",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
