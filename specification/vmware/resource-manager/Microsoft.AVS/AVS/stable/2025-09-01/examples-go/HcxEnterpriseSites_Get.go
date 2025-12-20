package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/HcxEnterpriseSites_Get.json
func ExampleHcxEnterpriseSitesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHcxEnterpriseSitesClient().Get(ctx, "group1", "cloud1", "site1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.HcxEnterpriseSitesClientGetResponse{
	// 	HcxEnterpriseSite: &armavs.HcxEnterpriseSite{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/hcxEnterpriseSites/site1"),
	// 		Name: to.Ptr("site1"),
	// 		Properties: &armavs.HcxEnterpriseSiteProperties{
	// 			ActivationKey: to.Ptr("0276EF1A9A1749A5A362BF73EA9F8D0D"),
	// 			Status: to.Ptr(armavs.HcxEnterpriseSiteStatusAvailable),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/hcxEnterpriseSites"),
	// 	},
	// }
}
