package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGroups.json
func ExampleGroupClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGroupClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.GroupClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
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
		// page.GroupCollection = armapimanagement.GroupCollection{
		// 	Count: to.Ptr[int64](4),
		// 	Value: []*armapimanagement.GroupContract{
		// 		{
		// 			Name: to.Ptr("5600b59375ff190048020001"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/groups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/5600b59375ff190048020001"),
		// 			Properties: &armapimanagement.GroupContractProperties{
		// 				Type: to.Ptr(armapimanagement.GroupTypeSystem),
		// 				Description: to.Ptr("Administrators is a built-in group. Its membership is managed by the system. Microsoft Azure subscription administrators fall into this group."),
		// 				BuiltIn: to.Ptr(true),
		// 				DisplayName: to.Ptr("Administrators"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("59306a29e4bbd510dc24e5f9"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/groups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/59306a29e4bbd510dc24e5f9"),
		// 			Properties: &armapimanagement.GroupContractProperties{
		// 				Type: to.Ptr(armapimanagement.GroupTypeExternal),
		// 				Description: to.Ptr("awesome group of people"),
		// 				BuiltIn: to.Ptr(false),
		// 				DisplayName: to.Ptr("AwesomeGroup (samiraad.onmicrosoft.com)"),
		// 				ExternalID: to.Ptr("aad://samiraad.onmicrosoft.com/groups/3773adf4-032e-4d25-9988-eaff9ca72eca"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5600b59375ff190048020002"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/groups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/5600b59375ff190048020002"),
		// 			Properties: &armapimanagement.GroupContractProperties{
		// 				Type: to.Ptr(armapimanagement.GroupTypeSystem),
		// 				Description: to.Ptr("Developers is a built-in group. Its membership is managed by the system. Signed-in users fall into this group."),
		// 				BuiltIn: to.Ptr(true),
		// 				DisplayName: to.Ptr("Developers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5600b59375ff190048020003"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/groups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/groups/5600b59375ff190048020003"),
		// 			Properties: &armapimanagement.GroupContractProperties{
		// 				Type: to.Ptr(armapimanagement.GroupTypeSystem),
		// 				Description: to.Ptr("Guests is a built-in group. Its membership is managed by the system. Unauthenticated users visiting the developer portal fall into this group."),
		// 				BuiltIn: to.Ptr(true),
		// 				DisplayName: to.Ptr("Guests"),
		// 			},
		// 	}},
		// }
	}
}
