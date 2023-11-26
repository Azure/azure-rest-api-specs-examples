package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/UserSession_ListByHostPool.json
func ExampleUserSessionsClient_NewListByHostPoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUserSessionsClient().NewListByHostPoolPager("resourceGroup1", "hostPool1", &armdesktopvirtualization.UserSessionsClientListByHostPoolOptions{Filter: to.Ptr("userPrincipalName eq 'user1@microsoft.com' and state eq 'active'"),
		PageSize:     to.Ptr[int32](10),
		IsDescending: to.Ptr(true),
		InitialSkip:  to.Ptr[int32](0),
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
		// page.UserSessionList = armdesktopvirtualization.UserSessionList{
		// 	Value: []*armdesktopvirtualization.UserSession{
		// 		{
		// 			Name: to.Ptr("1"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostPools/sessionHosts/userSessions"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1/sessionHosts/sessionHost1.microsoft.com/userSessions/1"),
		// 			Properties: &armdesktopvirtualization.UserSessionProperties{
		// 				ActiveDirectoryUserName: to.Ptr("WVDARM\\user1"),
		// 				ApplicationType: to.Ptr(armdesktopvirtualization.ApplicationTypeDesktop),
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				SessionState: to.Ptr(armdesktopvirtualization.SessionStateActive),
		// 				UserPrincipalName: to.Ptr("user1@microsoft.com"),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostPools/sessionHosts/userSessions"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1/sessionHosts/sessionHost1.microsoft.com/userSessions/2"),
		// 			Properties: &armdesktopvirtualization.UserSessionProperties{
		// 				ActiveDirectoryUserName: to.Ptr("WVDARM\\user1"),
		// 				ApplicationType: to.Ptr(armdesktopvirtualization.ApplicationTypeDesktop),
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				SessionState: to.Ptr(armdesktopvirtualization.SessionStateActive),
		// 				UserPrincipalName: to.Ptr("user1@microsoft.com"),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
