package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/AccountCapabilityHost/list.json
func ExampleAccountCapabilityHostsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountCapabilityHostsClient().NewListPager("test-rg", "account-1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CapabilityHostResourceArmPaginatedResult = armcognitiveservices.CapabilityHostResourceArmPaginatedResult{
		// 	Value: []*armcognitiveservices.CapabilityHost{
		// 		{
		// 			Name: to.Ptr("capabilityHostName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/accounts/capabilityHosts"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/account-1/capabilityHosts/capabilityHostName"),
		// 			Properties: &armcognitiveservices.CapabilityHostProperties{
		// 				Description: to.Ptr("string"),
		// 				Tags: map[string]*string{
		// 					"string": to.Ptr("string"),
		// 				},
		// 				CustomerSubnet: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroups/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubne"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.CapabilityHostProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
