package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/CreateHybridIdentityMetadata.json
func ExampleHybridIdentityMetadataClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHybridIdentityMetadataClient().Put(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster", armhybridcontainerservice.HybridIdentityMetadata{
		Properties: &armhybridcontainerservice.HybridIdentityMetadataProperties{
			PublicKey:   to.Ptr("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2"),
			ResourceUID: to.Ptr("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HybridIdentityMetadata = armhybridcontainerservice.HybridIdentityMetadata{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances/hybridIdentityMetadata"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/hybridIdentityMetadata/default"),
	// 	Properties: &armhybridcontainerservice.HybridIdentityMetadataProperties{
	// 		PublicKey: to.Ptr("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2"),
	// 		ResourceUID: to.Ptr("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c"),
	// 	},
	// }
}
