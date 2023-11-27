package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PureProxyArtifact/ArtifactList.json
func ExampleProxyArtifactClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProxyArtifactClient().NewListPager("TestResourceGroup", "TestPublisher", "TestArtifactStoreName", nil)
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
		// page.ProxyArtifactOverviewListResult = armhybridnetwork.ProxyArtifactOverviewListResult{
		// 	Value: []*armhybridnetwork.ProxyArtifactListOverview{
		// 		{
		// 			Name: to.Ptr("fedrbac1"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/publishers/artifactStores/artifactVersions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/TestResourceGroup/providers/Microsoft.HybridNetwork/publishers/TestPublisher/artifactStores/TestArtifactStore/artifactVersions"),
		// 		},
		// 		{
		// 			Name: to.Ptr("fedrbac2"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/publishers/artifactStores/artifactVersions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/TestResourceGroup/providers/Microsoft.HybridNetwork/publishers/TestPublisher/artifactStores/TestArtifactStore/artifactVersions"),
		// 	}},
		// }
	}
}
