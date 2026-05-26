package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2023-01-01-preview/SecurityOperators/ListSecurityOperators_example.json
func ExampleOperatorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperatorsClient().NewListPager("CloudPosture", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armsecurity.OperatorsClientListResponse{
		// 	OperatorList: armsecurity.OperatorList{
		// 		Value: []*armsecurity.OperatorResource{
		// 			{
		// 				Name: to.Ptr("DefenderCSPMSecurityOperator"),
		// 				Type: to.Ptr("Microsoft.Security/pricings/securityOperator"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/pricings/CloudPosture/securityOperators/DefenderCSPMSecurityOperator"),
		// 				Identity: &armsecurity.Identity{
		// 					Type: to.Ptr("SystemAssigned"),
		// 					PrincipalID: to.Ptr("44ee8e7e-7f52-4750-b937-27490fbf7663"),
		// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
