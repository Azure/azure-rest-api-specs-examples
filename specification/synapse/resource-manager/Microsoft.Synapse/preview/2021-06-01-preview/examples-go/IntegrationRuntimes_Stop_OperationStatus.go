package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Stop_OperationStatus.json
func ExampleGetClient_IntegrationRuntimeStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGetClient().IntegrationRuntimeStop(ctx, "drage-felles-prod-rg", "felles-prod-synapse-workspace", "SSIS-intergrationRuntime-Drage", "5752dcdf918e4aecb941245ddf6ebb83", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeStopOperationStatus = armsynapse.IntegrationRuntimeStopOperationStatus{
	// 	Properties: map[string]any{
	// 		"name": "5752dcdf918e4aecb941245ddf6ebb83",
	// 		"error": nil,
	// 		"properties": nil,
	// 		"status": "InProgress",
	// 	},
	// }
}
