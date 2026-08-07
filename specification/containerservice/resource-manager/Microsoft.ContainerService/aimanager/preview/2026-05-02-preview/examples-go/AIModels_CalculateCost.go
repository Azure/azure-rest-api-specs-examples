package armcontainerserviceaimanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerserviceaimanager/armcontainerserviceaimanager"
)

// Generated from example definition: 2026-05-02-preview/AIModels_CalculateCost.json
func ExampleAIModelsClient_CalculateCost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerserviceaimanager.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAIModelsClient().CalculateCost(ctx, "eastus", "9806f0c862fdd920", armcontainerserviceaimanager.CalculateCostRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerserviceaimanager.AIModelsClientCalculateCostResponse{
	// 	CalculateCostResponse: armcontainerserviceaimanager.CalculateCostResponse{
	// 		Currency: to.Ptr("USD"),
	// 		Plans: []*armcontainerserviceaimanager.CalculateCostPlan{
	// 			{
	// 				VMSize: to.Ptr("Standard_NC24ads_A100_v4"),
	// 				Quantization: to.Ptr("fp16"),
	// 				VMsPerReplica: to.Ptr[int32](1),
	// 				MaxAvailableReplicas: to.Ptr[int32](4),
	// 				ServingPerformanceEstimation: &armcontainerserviceaimanager.ServingPerformanceEstimation{
	// 					RelativeLatencyScore: to.Ptr[float32](0.72),
	// 					RelativeThroughputScore: to.Ptr[float32](0.65),
	// 				},
	// 				VMHourlyPrice: to.Ptr[float64](3.67),
	// 				TotalHourlyPrice: to.Ptr[float64](3.67),
	// 				PriceAsOf: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-05T06:05:51.775Z"); return t}()),
	// 				Feasible: to.Ptr(true),
	// 			},
	// 			{
	// 				VMSize: to.Ptr("Standard_ND96isr_H100_v5"),
	// 				Quantization: to.Ptr("fp16"),
	// 				VMsPerReplica: to.Ptr[int32](1),
	// 				MaxAvailableReplicas: to.Ptr[int32](0),
	// 				ServingPerformanceEstimation: &armcontainerserviceaimanager.ServingPerformanceEstimation{
	// 					RelativeLatencyScore: to.Ptr[float32](1),
	// 					RelativeThroughputScore: to.Ptr[float32](1),
	// 				},
	// 				VMHourlyPrice: to.Ptr[float64](98.32),
	// 				PriceAsOf: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-05T06:05:51.775Z"); return t}()),
	// 				Feasible: to.Ptr(false),
	// 				InfeasibilityReason: &armcontainerserviceaimanager.InfeasibilityReason{
	// 					Code: to.Ptr(armcontainerserviceaimanager.InfeasibleCodeInsufficientQuota),
	// 					Message: to.Ptr("Subscription does not have sufficient H100 quota in the target region."),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
