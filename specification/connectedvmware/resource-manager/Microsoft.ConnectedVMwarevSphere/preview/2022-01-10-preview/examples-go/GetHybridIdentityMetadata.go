package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7473936304533e6716fc4563401bf265dda4cb64/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetHybridIdentityMetadata.json
func ExampleHybridIdentityMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconnectedvmware.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHybridIdentityMetadataClient().Get(ctx, "testrg", "ContosoVm", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HybridIdentityMetadata = armconnectedvmware.HybridIdentityMetadata{
	// 	Name: to.Ptr("testItem"),
	// 	Type: to.Ptr("Microsoft.ConnectedVMwarevSphere/VitualMachines/hybridIdentityMetadatas"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ConnectedVMwarevSphere/VitualMachines/ContosoVm/hybridIdentityMetadatas/default"),
	// 	Properties: &armconnectedvmware.HybridIdentityMetadataProperties{
	// 		Identity: &armconnectedvmware.Identity{
	// 			Type: to.Ptr(armconnectedvmware.IdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("7b5129bc-8642-4a6a-95f8-63400ca6ec4d"),
	// 			TenantID: to.Ptr("ec46ca82-5d4a-4e3e-b4b7-e27f9318645d"),
	// 		},
	// 		PublicKey: to.Ptr("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2"),
	// 		VMID: to.Ptr("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c"),
	// 	},
	// }
}
