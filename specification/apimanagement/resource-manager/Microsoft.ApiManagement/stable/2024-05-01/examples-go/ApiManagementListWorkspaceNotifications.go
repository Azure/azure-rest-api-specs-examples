package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListWorkspaceNotifications.json
func ExampleWorkspaceNotificationClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceNotificationClient().NewListByServicePager("rg1", "apimService1", "wks1", &armapimanagement.WorkspaceNotificationClientListByServiceOptions{Top: nil,
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
		// page.NotificationCollection = armapimanagement.NotificationCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.NotificationContract{
		// 		{
		// 			Name: to.Ptr("RequestPublisherNotificationMessage"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/notifications"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/notifications/RequestPublisherNotificationMessage"),
		// 			Properties: &armapimanagement.NotificationContractProperties{
		// 				Description: to.Ptr("The following email recipients and users will receive email notifications about subscription requests for API products requiring approval."),
		// 				Recipients: &armapimanagement.RecipientsContractProperties{
		// 					Emails: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/notifications/RequestPublisherNotificationMessage/recipientEmails/contoso@live.com"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/notifications/RequestPublisherNotificationMessage/recipientEmails/foobar!live"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/notifications/RequestPublisherNotificationMessage/recipientEmails/foobar@live.com")},
		// 						Users: []*string{
		// 							to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/576823d0a40f7e74ec07d642")},
		// 						},
		// 						Title: to.Ptr("Subscription requests (requiring approval)"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("PurchasePublisherNotificationMessage"),
		// 					Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/notifications"),
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/notifications/PurchasePublisherNotificationMessage"),
		// 					Properties: &armapimanagement.NotificationContractProperties{
		// 						Description: to.Ptr("The following email recipients and users will receive email notifications about new API product subscriptions."),
		// 						Recipients: &armapimanagement.RecipientsContractProperties{
		// 							Emails: []*string{
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/notifications/RequestPublisherNotificationMessage/recipientEmails/contoso@live.com")},
		// 								Users: []*string{
		// 									to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1")},
		// 								},
		// 								Title: to.Ptr("New subscriptions"),
		// 							},
		// 					}},
		// 				}
	}
}
