package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Cluster_Delete_MaximumSet_Gen.json
func ExampleClusterClient_BeginDelete_clusterDeleteMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterClient().BeginDelete(ctx, "rgconfluent", "rwzpoelzgevhnkrvyqy", "gnijsroqxwwyyariafdnmkc", "zsvnfsirukovzkth", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
