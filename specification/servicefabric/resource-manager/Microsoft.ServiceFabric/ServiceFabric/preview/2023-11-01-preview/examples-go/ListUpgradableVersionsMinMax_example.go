package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ListUpgradableVersionsMinMax_example.json
func ExampleClustersClient_ListUpgradableVersions_getMinimumAndMaximumCodeVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListUpgradableVersions(ctx, "resRg", "myCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabric.ClustersClientListUpgradableVersionsResponse{
	// 	UpgradableVersionPathResult: armservicefabric.UpgradableVersionPathResult{
	// 		SupportedPath: []*string{
	// 			to.Ptr("7.0.0.0"),
	// 			to.Ptr("7.2.0.0"),
	// 		},
	// 	},
	// }
}
