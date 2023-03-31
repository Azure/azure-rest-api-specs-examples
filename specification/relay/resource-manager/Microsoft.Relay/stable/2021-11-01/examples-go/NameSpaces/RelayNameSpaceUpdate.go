package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/NameSpaces/RelayNameSpaceUpdate.json
func ExampleNamespacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespacesClient().Update(ctx, "RG-eg", "example-RelayRelayNamespace-01", armrelay.UpdateParameters{
		Tags: map[string]*string{
			"tag3": to.Ptr("value3"),
			"tag4": to.Ptr("value4"),
			"tag5": to.Ptr("value5"),
			"tag6": to.Ptr("value6"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Namespace = armrelay.Namespace{
	// 	Name: to.Ptr("example-RelayRelayNamespace-01"),
	// 	Type: to.Ptr("Microsoft.Relay/Namespaces"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/RG-eg/providers/Microsoft.Relay/namespaces/example-RelayRelayNamespace-01"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag3": to.Ptr("value3"),
	// 		"tag4": to.Ptr("value4"),
	// 		"tag5": to.Ptr("value5"),
	// 		"tag6": to.Ptr("value6"),
	// 	},
	// 	Properties: &armrelay.NamespaceProperties{
	// 		MetricID: to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff:example-Relayrelaynamespace-01"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
