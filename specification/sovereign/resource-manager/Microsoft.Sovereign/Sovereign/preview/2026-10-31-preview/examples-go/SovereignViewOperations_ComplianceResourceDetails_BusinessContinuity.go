package armsovereign_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_ComplianceResourceDetails_BusinessContinuity.json
func ExampleViewOperationsClient_ComplianceResourceDetails_sovereignViewOperationsComplianceResourceDetailsBusinessContinuity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewViewOperationsClient().ComplianceResourceDetails(ctx, "ExampleResourceGroup", "ExampleSV", armsovereign.ViewComplianceResourceDetailsContent{
		ResourceID:  to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Sql/servers/exampleserver/databases/exampledatabase"),
		Dimension:   to.Ptr(armsovereign.ComplianceDimensionBusinessContinuity),
		EvaluatedAt: to.Ptr(time.Date(2026, time.October, 15, 10, 0, 0, 0, time.UTC)),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientComplianceResourceDetailsResponse{
	// 	ViewComplianceResourceDetailsResult: armsovereign.ViewComplianceResourceDetailsResult{
	// 		ResourceID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Sql/servers/exampleserver/databases/exampledatabase"),
	// 		Dimension: to.Ptr(armsovereign.ComplianceDimensionBusinessContinuity),
	// 		EvaluatedAt: to.Ptr(time.Date(2026, time.October, 15, 10, 0, 0, 0, time.UTC)),
	// 		State: to.Ptr(armsovereign.PolicyComplianceStateConstrained),
	// 		Evaluations: []*armsovereign.ViewDataEvaluation{
	// 			{
	// 				EvaluationSourceID: to.Ptr("Microsoft.Sql"),
	// 				Intent: to.Ptr(armsovereign.ViewEvaluationIntentBusinessContinuity),
	// 				Source: to.Ptr(armsovereign.ViewLinkedResourceEvaluationSourceBusinessContinuity),
	// 				EvaluationOutcomes: []*armsovereign.ViewDataEvaluationOutcome{
	// 					{
	// 						ResourceID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Sql/servers/exampleserver/databases/exampledatabase"),
	// 						State: to.Ptr(armsovereign.PolicyComplianceStateConstrained),
	// 						ActualValue: to.Ptr(""),
	// 						ExpectedValue: to.Ptr(""),
	// 						Reason: to.Ptr(""),
	// 						ReasonCodeID: to.Ptr("reasonCode056"),
	// 						AdditionalInfo: to.Ptr(""),
	// 						AdditionalInfoCodeID: to.Ptr("009"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
