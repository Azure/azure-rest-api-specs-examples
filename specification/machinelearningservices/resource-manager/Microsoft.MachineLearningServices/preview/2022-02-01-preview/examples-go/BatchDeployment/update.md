```go
package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/update.json
func ExampleBatchDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewBatchDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		"testDeploymentName",
		armmachinelearning.PartialBatchDeploymentPartialTrackedResource{
			Identity: &armmachinelearning.PartialManagedServiceIdentity{
				Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
				UserAssignedIdentities: map[string]interface{}{
					"string": map[string]interface{}{},
				},
			},
			Kind:     to.Ptr("string"),
			Location: to.Ptr("string"),
			Properties: &armmachinelearning.PartialBatchDeployment{
				Description: to.Ptr("string"),
				CodeConfiguration: &armmachinelearning.PartialCodeConfiguration{
					CodeID:        to.Ptr("string"),
					ScoringScript: to.Ptr("string"),
				},
				Compute:       to.Ptr("string"),
				EnvironmentID: to.Ptr("string"),
				EnvironmentVariables: map[string]*string{
					"string": to.Ptr("string"),
				},
				ErrorThreshold:            to.Ptr[int32](1),
				LoggingLevel:              to.Ptr(armmachinelearning.BatchLoggingLevelInfo),
				MaxConcurrencyPerInstance: to.Ptr[int32](1),
				MiniBatchSize:             to.Ptr[int64](1),
				Model: &armmachinelearning.PartialIDAssetReference{
					ReferenceType: to.Ptr(armmachinelearning.ReferenceTypeID),
					AssetID:       to.Ptr("string"),
				},
				OutputAction:   to.Ptr(armmachinelearning.BatchOutputActionSummaryOnly),
				OutputFileName: to.Ptr("string"),
				Properties: map[string]*string{
					"string": to.Ptr("string"),
				},
				RetrySettings: &armmachinelearning.PartialBatchRetrySettings{
					MaxRetries: to.Ptr[int32](1),
					Timeout:    to.Ptr("PT5M"),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearning%2Farmmachinelearning%2Fv2.0.0-beta.1/sdk/resourcemanager/machinelearning/armmachinelearning/README.md) on how to add the SDK to your project and authenticate.
