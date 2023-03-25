package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/GetHybridIdentityMetadata.json
func ExampleHybridIdentityMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHybridIdentityMetadataClient().Get(ctx, "testrg", "ContosoTargetCluster", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HybridIdentityMetadata = armhybridcontainerservice.HybridIdentityMetadata{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.HybridContainerService/provisionedClusters/hybridIdentityMetadata"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridContainerService/provisionedClusters/ContosoTargetCluster/hybridIdentityMetadata/default"),
	// 	Properties: &armhybridcontainerservice.HybridIdentityMetadataProperties{
	// 		Identity: &armhybridcontainerservice.ProvisionedClusterIdentity{
	// 			Type: to.Ptr(armhybridcontainerservice.ResourceIdentityTypeSystemAssigned),
	// 			PrincipalID: to.Ptr("7b5129bc-8642-4a6a-95f8-63400ca6ec4d"),
	// 			TenantID: to.Ptr("ec46ca82-5d4a-4e3e-b4b7-e27f9318645d"),
	// 		},
	// 		PublicKey: to.Ptr("8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2"),
	// 		ResourceUID: to.Ptr("f8b82dff-38ef-4220-99ef-d3a3f86ddc6c"),
	// 	},
	// }
}
