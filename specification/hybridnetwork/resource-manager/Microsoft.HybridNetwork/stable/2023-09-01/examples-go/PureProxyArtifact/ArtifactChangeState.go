package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PureProxyArtifact/ArtifactChangeState.json
func ExampleProxyArtifactClient_BeginUpdateState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProxyArtifactClient().BeginUpdateState(ctx, "TestResourceGroup", "TestPublisher", "TestArtifactStoreName", "fedrbac", "1.0.0", armhybridnetwork.ArtifactChangeState{
		Properties: &armhybridnetwork.ArtifactChangeStateProperties{
			ArtifactState: to.Ptr(armhybridnetwork.ArtifactStateDeprecated),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProxyArtifactVersionsListOverview = armhybridnetwork.ProxyArtifactVersionsListOverview{
	// 	Name: to.Ptr("fedrbac"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/publishers/artifactStores/artifactVersions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/TestResourceGroup/providers/Microsoft.HybridNetwork/publishers/TestPublisher/artifactStores/TestArtifactStore/artifactVersions/1.0.0"),
	// 	Properties: &armhybridnetwork.ProxyArtifactOverviewPropertiesValue{
	// 		ArtifactState: to.Ptr(armhybridnetwork.ArtifactStateDeprecated),
	// 		ArtifactType: to.Ptr(armhybridnetwork.ArtifactTypeOCIArtifact),
	// 		ArtifactVersion: to.Ptr("1.0.0"),
	// 	},
	// }
}
