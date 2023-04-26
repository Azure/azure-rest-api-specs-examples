package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2023-01-01-preview/examples/SecurityOperators/GetSecurityOperatorByName_example.json
func ExampleOperatorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperatorsClient().Get(ctx, "CloudPosture", "DefenderCSPMSecurityOperator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperatorResource = armsecurity.OperatorResource{
	// 	Name: to.Ptr("DefenderCSPMSecurityOperator"),
	// 	Type: to.Ptr("Microsoft.Security/pricings/securityOperator"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/CloudPosture/securityOperators/DefenderCSPMSecurityOperator"),
	// 	Identity: &armsecurity.Identity{
	// 		Type: to.Ptr("SystemAssigned"),
	// 		PrincipalID: to.Ptr("44ee8e7e-7f52-4750-b937-27490fbf7663"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 	},
	// }
}
