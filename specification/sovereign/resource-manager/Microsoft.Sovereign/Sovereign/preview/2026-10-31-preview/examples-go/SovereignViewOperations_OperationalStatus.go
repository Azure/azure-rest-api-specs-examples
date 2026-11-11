package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_OperationalStatus.json
func ExampleViewOperationsClient_OperationalStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewViewOperationsClient().OperationalStatus(ctx, "ExampleResourceGroup", "ExampleSV", armsovereign.ViewOperationalStatusContent{
		Mode: to.Ptr(armsovereign.OperationalStatusRequestModeDetailed),
		EvaluationParameters: []*armsovereign.OperationalStatusEvaluationParameter{
			to.Ptr(armsovereign.OperationalStatusEvaluationParameterPermissionsEvaluation),
			to.Ptr(armsovereign.OperationalStatusEvaluationParameterPolicyAssignmentEvaluation),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientOperationalStatusResponse{
	// 	ViewOperationalStatusResult: armsovereign.ViewOperationalStatusResult{
	// 		Name: to.Ptr("ExampleSV"),
	// 		Status: to.Ptr(armsovereign.OperationalStatusStatesUnhealthy),
	// 		Details: []*armsovereign.ViewOperationalStatusDetails{
	// 			{
	// 				Reason: to.Ptr(armsovereign.OperationalStatusUnhealthyReasonMIPermissionMissing),
	// 				AffectedScopes: []*armsovereign.ViewAffectedScopes{
	// 					{
	// 						Name: to.Ptr("ExampleManagementGroup"),
	// 						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ExampleManagementGroup"),
	// 						Type: to.Ptr(armsovereign.ViewScopeTypeManagementGroup),
	// 						Status: to.Ptr(armsovereign.ScopeLevelOperationalStatusStatesUnhealthy),
	// 						ActionRequired: to.Ptr("ReaderRole"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Reason: to.Ptr(armsovereign.OperationalStatusUnhealthyReasonPolicyAssignmentMissing),
	// 				AffectedScopes: []*armsovereign.ViewAffectedScopes{
	// 					{
	// 						Name: to.Ptr("ExampleSubscription"),
	// 						ID: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f"),
	// 						Type: to.Ptr(armsovereign.ViewScopeTypeSubscription),
	// 						Status: to.Ptr(armsovereign.ScopeLevelOperationalStatusStatesUnhealthy),
	// 						ActionRequired: to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/providers/Microsoft.Authorization/policySetDefinitions/ExamplePolicySet"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
