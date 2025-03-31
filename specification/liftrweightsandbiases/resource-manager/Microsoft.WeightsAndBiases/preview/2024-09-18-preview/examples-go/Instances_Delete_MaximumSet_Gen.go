package armweightsandbiases_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/weightsandbiases/armweightsandbiases"
)

// Generated from example definition: 2024-09-18-preview/Instances_Delete_MaximumSet_Gen.json
func ExampleInstancesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armweightsandbiases.NewClientFactory("0BCB047F-CB71-4DFE-B0BD-88519F411C2F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInstancesClient().BeginDelete(ctx, "rgopenapi", "myinstance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
