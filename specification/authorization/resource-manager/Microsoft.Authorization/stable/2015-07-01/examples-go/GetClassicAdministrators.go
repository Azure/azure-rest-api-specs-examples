package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetClassicAdministrators.json
func ExampleClassicAdministratorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClassicAdministratorsClient().NewListPager(nil)
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
		// page.ClassicAdministratorListResult = armauthorization.ClassicAdministratorListResult{
		// 	Value: []*armauthorization.ClassicAdministrator{
		// 		{
		// 			Name: to.Ptr("classicadminid"),
		// 			Type: to.Ptr("Microsoft.Authorization/classicAdministrators"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Authorization/classicAdministrators/classicadminid"),
		// 			Properties: &armauthorization.ClassicAdministratorProperties{
		// 				EmailAddress: to.Ptr("test@test.com"),
		// 				Role: to.Ptr("ServiceAdministrator;AccountAdministrator"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("classicadminid2"),
		// 			Type: to.Ptr("Microsoft.Authorization/classicAdministrators"),
		// 			ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Authorization/classicAdministrators/classicadminid2"),
		// 			Properties: &armauthorization.ClassicAdministratorProperties{
		// 				EmailAddress: to.Ptr("coadmin@test.com"),
		// 				Role: to.Ptr("CoAdministrator"),
		// 			},
		// 	}},
		// }
	}
}
