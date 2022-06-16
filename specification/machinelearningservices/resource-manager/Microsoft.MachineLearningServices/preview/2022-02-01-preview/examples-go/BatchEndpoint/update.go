package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchEndpoint/update.json
func ExampleBatchEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewBatchEndpointsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		armmachinelearning.PartialBatchEndpointPartialTrackedResource{
			Identity: &armmachinelearning.PartialManagedServiceIdentity{
				Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
				UserAssignedIdentities: map[string]interface{}{
					"string": map[string]interface{}{},
				},
			},
			Kind:     to.Ptr("string"),
			Location: to.Ptr("string"),
			Properties: &armmachinelearning.PartialBatchEndpoint{
				Defaults: &armmachinelearning.BatchEndpointDefaults{
					DeploymentName: to.Ptr("string"),
				},
			},
			SKU: &armmachinelearning.PartialSKU{
				Name:     to.Ptr("string"),
				Capacity: to.Ptr[int32](1),
				Family:   to.Ptr("string"),
				Size:     to.Ptr("string"),
				Tier:     to.Ptr(armmachinelearning.SKUTierFree),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
