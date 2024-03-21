package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01a71545e82bb98b8137d3038150c436d46a59ed/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Create.json
func ExampleManagedPrivateEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedPrivateEndpointsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleManagedVirtualNetworkName", "exampleManagedPrivateEndpointName", armdatafactory.ManagedPrivateEndpointResource{
		Properties: &armdatafactory.ManagedPrivateEndpoint{
			Fqdns:                 []*string{},
			GroupID:               to.Ptr("blob"),
			PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
		},
	}, &armdatafactory.ManagedPrivateEndpointsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedPrivateEndpointResource = armdatafactory.ManagedPrivateEndpointResource{
	// 	Name: to.Ptr("exampleManagedPrivateEndpointName"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/managedVirtualNetworks/managedPrivateEndpoints"),
	// 	Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/exampleManagedVirtualNetworkName/managedPrivateEndpoints/exampleManagedPrivateEndpointName"),
	// 	Properties: &armdatafactory.ManagedPrivateEndpoint{
	// 		ConnectionState: &armdatafactory.ConnectionStateProperties{
	// 			Description: to.Ptr(""),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr("Pending"),
	// 		},
	// 		Fqdns: []*string{
	// 		},
	// 		GroupID: to.Ptr("blob"),
	// 		PrivateLinkResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
