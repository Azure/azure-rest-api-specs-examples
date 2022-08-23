package armconnectedvmware_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/CreateHybridIdentityMetadata.json
func ExampleHybridIdentityMetadataClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconnectedvmware.NewHybridIdentityMetadataClient("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "testrg", "ContosoVm", "default", armconnectedvmware.HybridIdentityMetadata{
		Properties: &armconnectedvmware.HybridIdentityMetadataProperties{
			PublicKey: to.Ptr("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2"),
			VMID:      to.Ptr("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
