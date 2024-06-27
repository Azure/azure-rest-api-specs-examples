package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/HcxEnterpriseSites_CreateOrUpdate.json
func ExampleHcxEnterpriseSitesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHcxEnterpriseSitesClient().CreateOrUpdate(ctx, "group1", "cloud1", "site1", armavs.HcxEnterpriseSite{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HcxEnterpriseSite = armavs.HcxEnterpriseSite{
	// 	Name: to.Ptr("site1"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/hcxEnterpriseSites"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/hcxEnterpriseSites/site1"),
	// 	Properties: &armavs.HcxEnterpriseSiteProperties{
	// 		ActivationKey: to.Ptr("0276EF1A9A1749A5A362BF73EA9F8D0D"),
	// 		Status: to.Ptr(armavs.HcxEnterpriseSiteStatusAvailable),
	// 	},
	// }
}
