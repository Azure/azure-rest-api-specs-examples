package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ExternalSecuritySolutions/GetExternalSecuritySolution_example.json
func ExampleExternalSecuritySolutionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExternalSecuritySolutionsClient().Get(ctx, "defaultresourcegroup-eus", "centralus", "aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.ExternalSecuritySolutionsClientGetResponse{
	// 	                            ExternalSecuritySolutionClassification: &armsecurity.AADExternalSecuritySolution{
	// 		Location: to.Ptr("eastus"),
	// 		Name: to.Ptr("aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus"),
	// 		Type: to.Ptr("Microsoft.Security/locations/externalSecuritySolutions"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/defaultresourcegroup-eus/providers/Microsoft.Security/locations/centralus/externalSecuritySolutions/aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus"),
	// 		Kind: to.Ptr(armsecurity.ExternalSecuritySolutionKindAAD),
	// 		Properties: &armsecurity.AADSolutionProperties{
	// 			DeviceType: to.Ptr("Azure Active Directory Identity Protection"),
	// 			DeviceVendor: to.Ptr("Microsoft"),
	// 			Workspace: &armsecurity.ConnectedWorkspace{
	// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/defaultresourcegroup-eus/providers/Microsoft.OperationalInsights/workspaces/defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus"),
	// 			},
	// 			AdditionalProperties: map[string]any{
	// 				"connectivityState": "Discovered",
	// 			},
	// 		},
	// 	},
	// 	                        }
}
