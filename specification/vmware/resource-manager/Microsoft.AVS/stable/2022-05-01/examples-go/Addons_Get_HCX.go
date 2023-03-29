package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/Addons_Get_HCX.json
func ExampleAddonsClient_Get_addonsGetHcx() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAddonsClient().Get(ctx, "group1", "cloud1", "hcx", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Addon = armavs.Addon{
	// 	Name: to.Ptr("hcx"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/addons"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/addons/hcx"),
	// 	Properties: &armavs.AddonHcxProperties{
	// 		AddonType: to.Ptr(armavs.AddonTypeHCX),
	// 		ProvisioningState: to.Ptr(armavs.AddonProvisioningStateSucceeded),
	// 		Offer: to.Ptr("VMware MaaS Cloud Provider (Enterprise)"),
	// 	},
	// }
}
