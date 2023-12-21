package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ListUpgradableVersionsPath_example.json
func ExampleClustersClient_ListUpgradableVersions_getUpgradePath() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClustersClient().ListUpgradableVersions(ctx, "resRg", "myCluster", &armservicefabric.ClustersClientListUpgradableVersionsOptions{VersionsDescription: &armservicefabric.UpgradableVersionsDescription{
		TargetVersion: to.Ptr("7.2.432.9590"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UpgradableVersionPathResult = armservicefabric.UpgradableVersionPathResult{
	// 	SupportedPath: []*string{
	// 		to.Ptr("6.4.664.9590"),
	// 		to.Ptr("7.0.466.9590"),
	// 		to.Ptr("7.0.470.9590"),
	// 		to.Ptr("7.2.432.9590")},
	// 	}
}
