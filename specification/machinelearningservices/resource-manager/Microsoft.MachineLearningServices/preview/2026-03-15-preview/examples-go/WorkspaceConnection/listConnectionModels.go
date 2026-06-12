package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/WorkspaceConnection/listConnectionModels.json
func ExampleConnectionClient_GetAllModels() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectionClient().GetAllModels(ctx, "test-rg", "aml-workspace-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.ConnectionClientGetAllModelsResponse{
	// 	EndpointModels: armmachinelearning.EndpointModels{
	// 		Value: []*armmachinelearning.EndpointModelProperties{
	// 			{
	// 				Name: to.Ptr("ada"),
	// 				Format: to.Ptr("OpenAI"),
	// 				Capabilities: map[string]*string{
	// 					"FineTuneTokensMaxValue": to.Ptr("2000000000"),
	// 					"completion": to.Ptr("true"),
	// 					"fineTune": to.Ptr("true"),
	// 					"inference": to.Ptr("false"),
	// 					"scaleType": to.Ptr("Manual"),
	// 					"search": to.Ptr("true"),
	// 				},
	// 				Deprecation: &armmachinelearning.EndpointModelDeprecationProperties{
	// 					FineTune: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-05T00:00:00+00:00"); return t}()),
	// 					Inference: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-05T00:00:00+00:00"); return t}()),
	// 				},
	// 				FinetuneCapabilities: map[string]*string{
	// 					"FineTuneTokensMaxValue": to.Ptr("2000000000"),
	// 					"completion": to.Ptr("true"),
	// 					"scaleType": to.Ptr("Manual,Standard"),
	// 					"search": to.Ptr("true"),
	// 				},
	// 				IsDefaultVersion: to.Ptr(false),
	// 				MaxCapacity: to.Ptr[int32](3),
	// 				SKUs: []*armmachinelearning.EndpointModelSKUProperties{
	// 					{
	// 						Name: to.Ptr("Standard"),
	// 						Capacity: &armmachinelearning.EndpointModelSKUCapacityProperties{
	// 							Maximum: to.Ptr[int32](10000),
	// 						},
	// 						ConnectionIDs: []*string{
	// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-1"),
	// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-2"),
	// 						},
	// 						DeprecationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-01T00:00:00+00:00"); return t}()),
	// 						RateLimits: []*armmachinelearning.EndpointModelSKURateLimitProperties{
	// 							{
	// 								Count: to.Ptr[float32](1),
	// 								RenewalPeriod: to.Ptr[float32](10),
	// 								Rules: []*armmachinelearning.EndpointModelSKURateLimitRuleProperties{
	// 								},
	// 							},
	// 							{
	// 								Count: to.Ptr[float32](1000),
	// 								RenewalPeriod: to.Ptr[float32](60),
	// 								Rules: []*armmachinelearning.EndpointModelSKURateLimitRuleProperties{
	// 								},
	// 							},
	// 						},
	// 						UsageName: to.Ptr("OpenAI.Standard.ada"),
	// 					},
	// 				},
	// 				SystemData: &armmachinelearning.SystemData{
	// 					CreatedBy: to.Ptr("Microsoft"),
	// 					LastModifiedBy: to.Ptr("Microsoft"),
	// 				},
	// 				Version: to.Ptr("1"),
	// 			},
	// 			{
	// 				Name: to.Ptr("babbage"),
	// 				Format: to.Ptr("OpenAI"),
	// 				Capabilities: map[string]*string{
	// 					"FineTuneTokensMaxValue": to.Ptr("2000000000"),
	// 					"completion": to.Ptr("true"),
	// 					"fineTune": to.Ptr("true"),
	// 					"inference": to.Ptr("false"),
	// 					"scaleType": to.Ptr("Manual"),
	// 					"search": to.Ptr("true"),
	// 				},
	// 				Deprecation: &armmachinelearning.EndpointModelDeprecationProperties{
	// 					FineTune: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-05T00:00:00+00:00"); return t}()),
	// 					Inference: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-05T00:00:00+00:00"); return t}()),
	// 				},
	// 				FinetuneCapabilities: map[string]*string{
	// 					"FineTuneTokensMaxValue": to.Ptr("2000000000"),
	// 					"completion": to.Ptr("true"),
	// 					"scaleType": to.Ptr("Manual,Standard"),
	// 					"search": to.Ptr("true"),
	// 				},
	// 				IsDefaultVersion: to.Ptr(false),
	// 				MaxCapacity: to.Ptr[int32](3),
	// 				SKUs: []*armmachinelearning.EndpointModelSKUProperties{
	// 					{
	// 						Name: to.Ptr("Standard"),
	// 						Capacity: &armmachinelearning.EndpointModelSKUCapacityProperties{
	// 							Maximum: to.Ptr[int32](10000),
	// 						},
	// 						ConnectionIDs: []*string{
	// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-1"),
	// 						},
	// 						DeprecationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-01T00:00:00+00:00"); return t}()),
	// 						RateLimits: []*armmachinelearning.EndpointModelSKURateLimitProperties{
	// 							{
	// 								Count: to.Ptr[float32](1),
	// 								RenewalPeriod: to.Ptr[float32](10),
	// 								Rules: []*armmachinelearning.EndpointModelSKURateLimitRuleProperties{
	// 								},
	// 							},
	// 							{
	// 								Count: to.Ptr[float32](1000),
	// 								RenewalPeriod: to.Ptr[float32](60),
	// 								Rules: []*armmachinelearning.EndpointModelSKURateLimitRuleProperties{
	// 								},
	// 							},
	// 						},
	// 						UsageName: to.Ptr("OpenAI.Standard.babbage"),
	// 					},
	// 				},
	// 				SystemData: &armmachinelearning.SystemData{
	// 					CreatedBy: to.Ptr("Microsoft"),
	// 					LastModifiedBy: to.Ptr("Microsoft"),
	// 				},
	// 				Version: to.Ptr("1"),
	// 			},
	// 		},
	// 	},
	// }
}
