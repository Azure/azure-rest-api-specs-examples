package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Restart.json
func ExampleIntegrationServiceEnvironmentsClient_Restart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationServiceEnvironmentsClient("f34b22a3-2202-4fb1-b040-1332bd928c84", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Restart(ctx,
		"testResourceGroup",
		"testIntegrationServiceEnvironment",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
