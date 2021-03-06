package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/createOrUpdate.json
func ExampleBatchDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewBatchDeploymentsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		"testDeploymentName",
		armmachinelearning.BatchDeploymentData{
			Location: to.Ptr("string"),
			Tags:     map[string]*string{},
			Identity: &armmachinelearning.ManagedServiceIdentity{
				Type: to.Ptr(armmachinelearning.ManagedServiceIdentityTypeSystemAssigned),
				UserAssignedIdentities: map[string]*armmachinelearning.UserAssignedIdentity{
					"string": {},
				},
			},
			Kind: to.Ptr("string"),
			Properties: &armmachinelearning.BatchDeploymentDetails{
				Description: to.Ptr("string"),
				CodeConfiguration: &armmachinelearning.CodeConfiguration{
					CodeID:        to.Ptr("string"),
					ScoringScript: to.Ptr("string"),
				},
				EnvironmentID: to.Ptr("string"),
				EnvironmentVariables: map[string]*string{
					"string": to.Ptr("string"),
				},
				Properties: map[string]*string{
					"string": to.Ptr("string"),
				},
				Compute:                   to.Ptr("string"),
				ErrorThreshold:            to.Ptr[int32](1),
				LoggingLevel:              to.Ptr(armmachinelearning.BatchLoggingLevelInfo),
				MaxConcurrencyPerInstance: to.Ptr[int32](1),
				MiniBatchSize:             to.Ptr[int64](1),
				Model: &armmachinelearning.IDAssetReference{
					ReferenceType: to.Ptr(armmachinelearning.ReferenceTypeID),
					AssetID:       to.Ptr("string"),
				},
				OutputAction:   to.Ptr(armmachinelearning.BatchOutputActionSummaryOnly),
				OutputFileName: to.Ptr("string"),
				Resources: &armmachinelearning.ResourceConfiguration{
					InstanceCount: to.Ptr[int32](1),
					InstanceType:  to.Ptr("string"),
					Properties: map[string]interface{}{
						"string": map[string]interface{}{
							"cd3c37dc-2876-4ca4-8a54-21bd7619724a": nil,
						},
					},
				},
				RetrySettings: &armmachinelearning.BatchRetrySettings{
					MaxRetries: to.Ptr[int32](1),
					Timeout:    to.Ptr("PT5M"),
				},
			},
			SKU: &armmachinelearning.SKU{
				Name:     to.Ptr("string"),
				Capacity: to.Ptr[int32](1),
				Family:   to.Ptr("string"),
				Size:     to.Ptr("string"),
				Tier:     to.Ptr(armmachinelearning.SKUTierFree),
			},
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
