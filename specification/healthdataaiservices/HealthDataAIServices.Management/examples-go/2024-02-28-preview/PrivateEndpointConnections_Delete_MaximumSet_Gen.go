package armhealthdataaiservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/PrivateEndpointConnections_Delete_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginDelete(ctx, "rgopenapi", "deidTest", "kgwgrrpabvrsrrvpcgcnfmyfgyrl", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
