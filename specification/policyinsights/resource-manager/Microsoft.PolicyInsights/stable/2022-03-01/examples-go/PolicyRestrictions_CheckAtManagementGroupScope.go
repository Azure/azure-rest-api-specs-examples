package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/05a9cdab363b8ec824094ee73950c04594325172/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtManagementGroupScope.json
func ExamplePolicyRestrictionsClient_CheckAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicyinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyRestrictionsClient().CheckAtManagementGroupScope(ctx, "financeMg", armpolicyinsights.CheckManagementGroupRestrictionsRequest{
		PendingFields: []*armpolicyinsights.PendingField{
			{
				Field: to.Ptr("type"),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckRestrictionsResult = armpolicyinsights.CheckRestrictionsResult{
	// 	ContentEvaluationResult: &armpolicyinsights.CheckRestrictionsResultContentEvaluationResult{
	// 		PolicyEvaluations: []*armpolicyinsights.PolicyEvaluationResult{
	// 		},
	// 	},
	// 	FieldRestrictions: []*armpolicyinsights.FieldRestrictions{
	// 		{
	// 			Field: to.Ptr("type"),
	// 			Restrictions: []*armpolicyinsights.FieldRestriction{
	// 				{
	// 					Policy: &armpolicyinsights.PolicyReference{
	// 						PolicyAssignmentID: to.Ptr("/providers/Microsoft.Management/managementGroups/financeMg/providers/microsoft.authorization/policyAssignments/7EB1508A"),
	// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/financeMg/providers/microsoft.authorization/policyDefinitions/allowedTypes"),
	// 						PolicyDefinitionReferenceID: to.Ptr("DefRef"),
	// 						PolicySetDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/financeMg/providers/microsoft.authorization/policySetDefinitions/735551F1"),
	// 					},
	// 					Result: to.Ptr(armpolicyinsights.FieldRestrictionResultRequired),
	// 					Values: []*string{
	// 						to.Ptr("Microsoft.Compute/virtualMachines")},
	// 				}},
	// 		}},
	// 	}
}
