package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub/v2"
)

// Generated from example definition: 2025-05-01-preview/Clusters/ClusterQuotaConfigurationPatch.json
func ExampleConfigurationClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationClient().Patch(ctx, "ArunMonocle", "testCluster", armeventhub.ClusterQuotaConfigurationProperties{
		Settings: map[string]*string{
			"eventhub-per-namespace-quota": to.Ptr("20"),
			"namespaces-per-cluster-quota": to.Ptr("200"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armeventhub.ConfigurationClientPatchResponse{
	// 	ClusterQuotaConfigurationProperties: armeventhub.ClusterQuotaConfigurationProperties{
	// 		Settings: map[string]*string{
	// 			"eventhub-per-namespace-quota": to.Ptr("20"),
	// 			"namespaces-per-cluster-quota": to.Ptr("200"),
	// 		},
	// 	},
	// }
}
