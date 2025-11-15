package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/AccountCapabilityHost/get.json
func ExampleAccountCapabilityHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountCapabilityHostsClient().Get(ctx, "test-rg", "account-1", "capabilityHostName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapabilityHost = armcognitiveservices.CapabilityHost{
	// 	Name: to.Ptr("capabilityHostName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/capabilityHosts"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/account-1/capabilityHosts/capabilityHostName"),
	// 	Properties: &armcognitiveservices.CapabilityHostProperties{
	// 		Description: to.Ptr("string"),
	// 		Tags: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		CustomerSubnet: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroups/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubne"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.CapabilityHostProvisioningStateSucceeded),
	// 	},
	// }
}
