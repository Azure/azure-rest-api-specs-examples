package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/SessionHost_List.json
func ExampleSessionHostsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSessionHostsClient().NewListPager("resourceGroup1", "hostPool1", &armdesktopvirtualization.SessionHostsClientListOptions{PageSize: to.Ptr[int32](10),
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
		// page.SessionHostList = armdesktopvirtualization.SessionHostList{
		// 	Value: []*armdesktopvirtualization.SessionHost{
		// 		{
		// 			Name: to.Ptr("sessionHost1.microsoft.com"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostPools/sessionhosts"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1/sessionHosts/sessionHost1.microsoft.com"),
		// 			Properties: &armdesktopvirtualization.SessionHostProperties{
		// 				AgentVersion: to.Ptr("1.0.0.1391"),
		// 				AllowNewSession: to.Ptr(true),
		// 				AssignedUser: to.Ptr("user1@microsoft.com"),
		// 				FriendlyName: to.Ptr("friendly"),
		// 				LastHeartBeat: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				OSVersion: to.Ptr("10.0.17763"),
		// 				ResourceID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.Compute/virtualMachines/sessionHost1"),
		// 				SessionHostHealthCheckResults: []*armdesktopvirtualization.SessionHostHealthCheckReport{
		// 					{
		// 						AdditionalFailureDetails: &armdesktopvirtualization.SessionHostHealthCheckFailureDetails{
		// 							ErrorCode: to.Ptr[int32](0),
		// 							LastHealthCheckDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-14T02:09:37.6236843Z"); return t}()),
		// 							Message: to.Ptr("SessionHost healthy: is joined to domain ≤wvdarmtest1.net≥"),
		// 						},
		// 						HealthCheckName: to.Ptr(armdesktopvirtualization.HealthCheckNameDomainJoinedCheck),
		// 						HealthCheckResult: to.Ptr(armdesktopvirtualization.HealthCheckResultHealthCheckSucceeded),
		// 				}},
		// 				Sessions: to.Ptr[int32](1),
		// 				Status: to.Ptr(armdesktopvirtualization.StatusAvailable),
		// 				StatusTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				SxSStackVersion: to.Ptr("rdp-sxs190816002"),
		// 				UpdateErrorMessage: to.Ptr(""),
		// 				UpdateState: to.Ptr(armdesktopvirtualization.UpdateStateSucceeded),
		// 				VirtualMachineID: to.Ptr("29491b54-c033-4dec-b09a-18bf0ebafaef"),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sessionHost2.microsoft.com"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/hostPools/sessionhosts"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1/sessionHosts/sessionHost2microsoft.com"),
		// 			Properties: &armdesktopvirtualization.SessionHostProperties{
		// 				AgentVersion: to.Ptr("1.0.0.1391"),
		// 				AllowNewSession: to.Ptr(true),
		// 				AssignedUser: to.Ptr("user2@microsoft.com"),
		// 				FriendlyName: to.Ptr("friendly"),
		// 				LastHeartBeat: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				OSVersion: to.Ptr("10.0.17763"),
		// 				ResourceID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.Compute/virtualMachines/sessionHost2"),
		// 				SessionHostHealthCheckResults: []*armdesktopvirtualization.SessionHostHealthCheckReport{
		// 					{
		// 						AdditionalFailureDetails: &armdesktopvirtualization.SessionHostHealthCheckFailureDetails{
		// 							ErrorCode: to.Ptr[int32](0),
		// 							LastHealthCheckDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-14T02:09:37.6236843Z"); return t}()),
		// 							Message: to.Ptr("SessionHost healthy: is joined to domain ≤wvdarmtest1.net≥"),
		// 						},
		// 						HealthCheckName: to.Ptr(armdesktopvirtualization.HealthCheckNameDomainJoinedCheck),
		// 						HealthCheckResult: to.Ptr(armdesktopvirtualization.HealthCheckResultHealthCheckSucceeded),
		// 				}},
		// 				Sessions: to.Ptr[int32](1),
		// 				Status: to.Ptr(armdesktopvirtualization.StatusAvailable),
		// 				StatusTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
		// 				SxSStackVersion: to.Ptr("rdp-sxs190816002"),
		// 				UpdateErrorMessage: to.Ptr(""),
		// 				UpdateState: to.Ptr(armdesktopvirtualization.UpdateStateSucceeded),
		// 				VirtualMachineID: to.Ptr("39491b54-c033-4dec-b09a-18bf0ebafaef"),
		// 			},
		// 			SystemData: &armdesktopvirtualization.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
