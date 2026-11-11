package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_ComplianceDetails.json
func ExampleViewOperationsClient_ComplianceDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewViewOperationsClient().ComplianceDetails(ctx, "ExampleResourceGroup", "ExampleSV", armsovereign.ViewComplianceDetailsContent{
		Top:  to.Ptr[int32](50),
		Skip: to.Ptr[int32](25),
		ComplianceDimensions: []*armsovereign.ViewComplianceDetailsDimensionFilter{
			{
				Dimension: to.Ptr(armsovereign.ComplianceDimensionEncryptionAtRest),
				States: []*armsovereign.PolicyComplianceState{
					to.Ptr(armsovereign.PolicyComplianceStateCompliant),
				},
			},
			{
				Dimension: to.Ptr(armsovereign.ComplianceDimensionBusinessContinuity),
				States: []*armsovereign.PolicyComplianceState{
					to.Ptr(armsovereign.PolicyComplianceStateConstrained),
				},
			},
		},
		Filters: []*armsovereign.ViewComplianceDetailsFilters{
			{
				Field:    to.Ptr(armsovereign.ViewFilterFieldRegion),
				Operator: to.Ptr(armsovereign.ViewComplianceDetailsFilterOperatorsIn),
				Value: []*string{
					to.Ptr("northeurope"),
				},
			},
			{
				Field:    to.Ptr(armsovereign.ViewFilterFieldParentResourceName),
				Operator: to.Ptr(armsovereign.ViewComplianceDetailsFilterOperatorsIn),
				Value: []*string{
					to.Ptr("example-parent-vm"),
				},
			},
			{
				Field:    to.Ptr(armsovereign.ViewFilterFieldBusinessContinuity),
				Operator: to.Ptr(armsovereign.ViewComplianceDetailsFilterOperatorsIn),
				Value: []*string{
					to.Ptr("Constrained"),
				},
			},
		},
		Sort: []*armsovereign.ViewComplianceDetailsSort{
			{
				Field:     to.Ptr("ResourceId"),
				Direction: to.Ptr(armsovereign.ViewComplianceDetailsSortDirectionsAsc),
				SortingOrder: []*armsovereign.PolicyComplianceState{
					to.Ptr(armsovereign.PolicyComplianceStateCompliant),
					to.Ptr(armsovereign.PolicyComplianceStateNonCompliant),
					to.Ptr(armsovereign.PolicyComplianceStateConstrained),
				},
			},
		},
		Search: &armsovereign.ViewComplianceDetailsSearch{
			Field: to.Ptr("ResourceId"),
			Value: to.Ptr("example-vm"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientComplianceDetailsResponse{
	// 	ViewComplianceDetailsResult: armsovereign.ViewComplianceDetailsResult{
	// 		ReportStatus: &armsovereign.ReportStatus{
	// 			State: to.Ptr(armsovereign.ReportStateUpToDate),
	// 			NextAttempt: to.Ptr(time.Date(2026, time.August, 1, 6, 20, 21, 0, time.UTC)),
	// 		},
	// 		LastEvaluated: to.Ptr(time.Date(2026, time.March, 15, 12, 30, 0, 0, time.UTC)),
	// 		TotalResources: to.Ptr[int32](2),
	// 		Resources: []*armsovereign.ViewComplianceDetailResource{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Compute/virtualMachines/example-vm-01"),
	// 				ParentResourceID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup"),
	// 				Location: to.Ptr("northeurope"),
	// 				DimensionDetails: []*armsovereign.ViewDimensionDetail{
	// 					{
	// 						Dimension: to.Ptr(armsovereign.ComplianceDimensionEncryptionAtRest),
	// 						State: to.Ptr(armsovereign.PolicyComplianceStateCompliant),
	// 						EvaluatedAt: to.Ptr(time.Date(2026, time.March, 15, 10, 32, 29, 0, time.UTC)),
	// 						Sources: []*armsovereign.ViewResourceEvaluationSource{
	// 							to.Ptr(armsovereign.ViewResourceEvaluationSourcePolicy),
	// 						},
	// 						ComplianceDetails: []*armsovereign.ViewComplianceDetail{
	// 							{
	// 								PolicyInitiativeID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policySetDefinitions/ExamplePolicySet"),
	// 								PolicyAssignmentID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policyAssignments/68fcd240-8a80-4e3b-8590-52fefd22cdd8"),
	// 								EvaluatedAt: to.Ptr(time.Date(2026, time.March, 15, 10, 32, 29, 0, time.UTC)),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Storage/storageAccounts/examplest01"),
	// 				ParentResourceID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup"),
	// 				Location: to.Ptr("northeurope"),
	// 				DimensionDetails: []*armsovereign.ViewDimensionDetail{
	// 					{
	// 						Dimension: to.Ptr(armsovereign.ComplianceDimensionEncryptionAtRest),
	// 						State: to.Ptr(armsovereign.PolicyComplianceStateNonCompliant),
	// 						EvaluatedAt: to.Ptr(time.Date(2026, time.March, 15, 11, 15, 45, 0, time.UTC)),
	// 						Sources: []*armsovereign.ViewResourceEvaluationSource{
	// 							to.Ptr(armsovereign.ViewResourceEvaluationSourcePolicy),
	// 							to.Ptr(armsovereign.ViewResourceEvaluationSourceSovereignIntelligence),
	// 						},
	// 						ComplianceDetails: []*armsovereign.ViewComplianceDetail{
	// 							{
	// 								PolicyInitiativeID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policySetDefinitions/ExamplePolicySet"),
	// 								PolicyAssignmentID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policyAssignments/f868169e-9bc7-43d9-8ae0-362d24e926b7"),
	// 								EvaluatedAt: to.Ptr(time.Date(2026, time.March, 15, 11, 15, 45, 0, time.UTC)),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						Dimension: to.Ptr(armsovereign.ComplianceDimensionBusinessContinuity),
	// 						State: to.Ptr(armsovereign.PolicyComplianceStateConstrained),
	// 						EvaluatedAt: to.Ptr(time.Date(2026, time.March, 15, 11, 15, 45, 0, time.UTC)),
	// 						Sources: []*armsovereign.ViewResourceEvaluationSource{
	// 							to.Ptr(armsovereign.ViewResourceEvaluationSourceSovereignIntelligence),
	// 						},
	// 						ComplianceDetails: []*armsovereign.ViewComplianceDetail{
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
