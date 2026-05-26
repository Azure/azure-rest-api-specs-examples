package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListClientApplications.json
func ExampleClientApplicationClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClientApplicationClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page = armapimanagement.ClientApplicationClientListByServiceResponse{
		// 	ClientApplicationCollection: armapimanagement.ClientApplicationCollection{
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.ClientApplicationContract{
		// 			{
		// 				Name: to.Ptr("testAppId"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/clientApplications"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/clientApplications/testAppId"),
		// 				Properties: &armapimanagement.ClientApplicationContractProperties{
		// 					Description: to.Ptr("This is just an example application"),
		// 					DisplayName: to.Ptr("Test Application"),
		// 					EntraApplicationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					EntraTenantID: to.Ptr("00000000-0000-0000-0000-000000000010"),
		// 					OwnerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/userID1"),
		// 					State: to.Ptr(armapimanagement.ClientApplicationStateActive),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testAppId2"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/clientApplications"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/clientApplications/testAppId2"),
		// 				Properties: &armapimanagement.ClientApplicationContractProperties{
		// 					Description: to.Ptr("This is just an example application 2"),
		// 					DisplayName: to.Ptr("Test Application 2"),
		// 					EntraApplicationID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 					EntraTenantID: to.Ptr("00000000-0000-0000-0000-000000000011"),
		// 					OwnerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/userID2"),
		// 					State: to.Ptr(armapimanagement.ClientApplicationStateActive),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testAppId3"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/clientApplications"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/clientApplications/testAppId3"),
		// 				Properties: &armapimanagement.ClientApplicationContractProperties{
		// 					Description: to.Ptr("This is just an example application 3"),
		// 					DisplayName: to.Ptr("Test Application 3"),
		// 					EntraApplicationID: to.Ptr("00000000-0000-0000-0000-000000000002"),
		// 					EntraTenantID: to.Ptr("00000000-0000-0000-0000-000000000012"),
		// 					OwnerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/userID1"),
		// 					State: to.Ptr(armapimanagement.ClientApplicationStateActive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
