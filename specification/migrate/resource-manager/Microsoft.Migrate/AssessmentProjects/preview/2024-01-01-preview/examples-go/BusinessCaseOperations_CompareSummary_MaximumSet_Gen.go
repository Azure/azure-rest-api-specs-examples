package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/BusinessCaseOperations_CompareSummary_MaximumSet_Gen.json
func ExampleBusinessCaseOperationsClient_BeginCompareSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBusinessCaseOperationsClient().BeginCompareSummary(ctx, "rgopenapi", "multipleto8617project", "sample-business-case", map[string]any{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CompareSummary = armmigrationassessment.CompareSummary{
	// 	AzureArcEnabledOnPremisesCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](21),
	// 		ComputeCost: to.Ptr[float32](14),
	// 		EsuSavings: to.Ptr[float32](22),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](12),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](24),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](28),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](8),
	// 		SecurityCost: to.Ptr[float32](19),
	// 		StorageCost: to.Ptr[float32](28),
	// 	},
	// 	AzureAvsCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	AzureIaasCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	AzurePaasCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	FutureAzureArcEnabledOnPremisesCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](21),
	// 		ComputeCost: to.Ptr[float32](14),
	// 		EsuSavings: to.Ptr[float32](22),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](12),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](24),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](28),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](8),
	// 		SecurityCost: to.Ptr[float32](19),
	// 		StorageCost: to.Ptr[float32](28),
	// 	},
	// 	FutureCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](21),
	// 		ComputeCost: to.Ptr[float32](14),
	// 		EsuSavings: to.Ptr[float32](22),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](12),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](24),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](28),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](8),
	// 		SecurityCost: to.Ptr[float32](19),
	// 		StorageCost: to.Ptr[float32](28),
	// 	},
	// 	OnPremisesAvsCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	OnPremisesAvsDecommissionedCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	OnPremisesIaasCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	OnPremisesIaasDecommissionedCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	OnPremisesPaasCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// 	OnPremisesPaasDecommissionedCostDetails: &armmigrationassessment.CostDetails{
	// 		AhubSavings: to.Ptr[float32](1),
	// 		ComputeCost: to.Ptr[float32](19),
	// 		EsuSavings: to.Ptr[float32](28),
	// 		FacilitiesCost: to.Ptr[float32](19),
	// 		ItLaborCost: to.Ptr[float32](29),
	// 		LinuxAhubSavings: to.Ptr[float32](10),
	// 		ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 			ManagementCost: to.Ptr[float32](22),
	// 			ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 					Value: to.Ptr[float32](16),
	// 			}},
	// 		},
	// 		NetworkCost: to.Ptr[float32](28),
	// 		SecurityCost: to.Ptr[float32](4),
	// 		StorageCost: to.Ptr[float32](27),
	// 	},
	// }
}
