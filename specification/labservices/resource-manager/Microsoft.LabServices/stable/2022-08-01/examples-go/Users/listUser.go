package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4c2cdccf6ca3281dd50ed8788ce1de2e0d480973/specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Users/listUser.json
func ExampleUsersClient_NewListByLabPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlabservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsersClient().NewListByLabPager("testrg123", "testlab", &armlabservices.UsersClientListByLabOptions{Filter: nil})
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
		// page.PagedUsers = armlabservices.PagedUsers{
		// 	Value: []*armlabservices.User{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.LabServices/User"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labs/testlab/users/testuser"),
		// 			Properties: &armlabservices.UserProperties{
		// 				AdditionalUsageQuota: to.Ptr("PT10H"),
		// 				DisplayName: to.Ptr("Test User"),
		// 				Email: to.Ptr("testuser@contoso.com"),
		// 				InvitationState: to.Ptr(armlabservices.InvitationStateSent),
		// 				ProvisioningState: to.Ptr(armlabservices.ProvisioningStateSucceeded),
		// 				RegistrationState: to.Ptr(armlabservices.RegistrationStateRegistered),
		// 				TotalUsage: to.Ptr("PT10H"),
		// 			},
		// 			SystemData: &armlabservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T10:00:00.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("identity123"),
		// 				CreatedByType: to.Ptr(armlabservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T09:12:28.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identity123"),
		// 				LastModifiedByType: to.Ptr(armlabservices.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
