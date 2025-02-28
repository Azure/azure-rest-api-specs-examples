package armdatabasewatcher_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
)

// Generated from example definition: 2025-01-02/Targets_Delete_MaximumSet_Gen.json
func ExampleTargetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabasewatcher.NewClientFactory("49e0fbd3-75e8-44e7-96fd-5b64d9ad818d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTargetsClient().Delete(ctx, "apiTest-ddat4p", "databasemo3ej9ih", "monitoringh22eed", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatabasewatcher.TargetsClientDeleteResponse{
	// }
}
