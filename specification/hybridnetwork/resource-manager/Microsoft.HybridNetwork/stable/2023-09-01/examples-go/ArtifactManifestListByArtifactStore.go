package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ArtifactManifestListByArtifactStore.json
func ExampleArtifactManifestsClient_NewListByArtifactStorePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArtifactManifestsClient().NewListByArtifactStorePager("rg", "TestPublisher", "TestArtifactStore", nil)
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
		// page.ArtifactManifestListResult = armhybridnetwork.ArtifactManifestListResult{
		// 	Value: []*armhybridnetwork.ArtifactManifest{
		// 		{
		// 			Name: to.Ptr("TestManifest"),
		// 			Type: to.Ptr("microsoft.hybridnetwork/publishers/artifactStores/artifactManifests"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/microsoft.hybridnetwork/publishers/TestPublisher/artifactStores/TestArtifactStore/artifactManifests/TestManifest"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armhybridnetwork.ArtifactManifestPropertiesFormat{
		// 				ArtifactManifestState: to.Ptr(armhybridnetwork.ArtifactManifestStateUploaded),
		// 				Artifacts: []*armhybridnetwork.ManifestArtifactFormat{
		// 					{
		// 						ArtifactName: to.Ptr("fed-rbac"),
		// 						ArtifactType: to.Ptr(armhybridnetwork.ArtifactTypeOCIArtifact),
		// 						ArtifactVersion: to.Ptr("1.0.0"),
		// 					},
		// 					{
		// 						ArtifactName: to.Ptr("nginx"),
		// 						ArtifactType: to.Ptr(armhybridnetwork.ArtifactTypeOCIArtifact),
		// 						ArtifactVersion: to.Ptr("v1"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
