package armprivatedns_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/PrivateZoneListInResourceGroup.json
func ExamplePrivateZonesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armprivatedns.NewPrivateZonesClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("resourceGroup1",
		&armprivatedns.PrivateZonesClientListByResourceGroupOptions{Top: nil})
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
