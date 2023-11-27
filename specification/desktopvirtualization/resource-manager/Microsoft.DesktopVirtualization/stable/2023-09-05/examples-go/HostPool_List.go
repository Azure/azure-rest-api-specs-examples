package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/HostPool_List.json
func ExampleHostPoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHostPoolsClient().NewListPager(&armdesktopvirtualization.HostPoolsClientListOptions{PageSize: to.Ptr[int32](10),
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
		// page.HostPoolList = armdesktopvirtualization.HostPoolList{
		// 	Value: []*armdesktopvirtualization.HostPool{
		// 		{
		// 			Name: to.Ptr("hostPool1"),
		// 			Type: to.Ptr("/Microsoft.DesktopVirtualization/hostPools"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
		// 			Location: to.Ptr("centralus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armdesktopvirtualization.HostPoolProperties{
		// 				Description: to.Ptr("des1"),
		// 				AgentUpdate: &armdesktopvirtualization.AgentUpdateProperties{
		// 					Type: to.Ptr(armdesktopvirtualization.SessionHostComponentUpdateTypeScheduled),
		// 					MaintenanceWindowTimeZone: to.Ptr("Alaskan Standard Time"),
		// 					MaintenanceWindows: []*armdesktopvirtualization.MaintenanceWindowProperties{
		// 						{
		// 							DayOfWeek: to.Ptr(armdesktopvirtualization.DayOfWeekFriday),
		// 							Hour: to.Ptr[int32](7),
		// 						},
		// 						{
		// 							DayOfWeek: to.Ptr(armdesktopvirtualization.DayOfWeekSaturday),
		// 							Hour: to.Ptr[int32](8),
		// 					}},
		// 					UseSessionHostLocalTime: to.Ptr(false),
		// 				},
		// 				CloudPcResource: to.Ptr(false),
		// 				FriendlyName: to.Ptr("friendly"),
		// 				HostPoolType: to.Ptr(armdesktopvirtualization.HostPoolTypePooled),
		// 				LoadBalancerType: to.Ptr(armdesktopvirtualization.LoadBalancerTypeBreadthFirst),
		// 				MaxSessionLimit: to.Ptr[int32](999999),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				PersonalDesktopAssignmentType: to.Ptr(armdesktopvirtualization.PersonalDesktopAssignmentTypeAutomatic),
		// 				PreferredAppGroupType: to.Ptr(armdesktopvirtualization.PreferredAppGroupTypeDesktop),
		// 				RegistrationInfo: &armdesktopvirtualization.RegistrationInfo{
		// 					ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
		// 					RegistrationTokenOperation: to.Ptr(armdesktopvirtualization.RegistrationTokenOperationUpdate),
		// 					Token: to.Ptr("token"),
		// 				},
		// 				SsoClientID: to.Ptr("client"),
		// 				SsoClientSecretKeyVaultPath: to.Ptr("https://keyvault/secret"),
		// 				SsoSecretType: to.Ptr(armdesktopvirtualization.SSOSecretTypeSharedKey),
		// 				SsoadfsAuthority: to.Ptr("https://adfs"),
		// 				StartVMOnConnect: to.Ptr(false),
		// 				VMTemplate: to.Ptr("{json:json}"),
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
		// 			Name: to.Ptr("hostPool2"),
		// 			Type: to.Ptr("/Microsoft.DesktopVirtualization/hostPools"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool2"),
		// 			Location: to.Ptr("centralus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armdesktopvirtualization.HostPoolProperties{
		// 				Description: to.Ptr("des1"),
		// 				AgentUpdate: &armdesktopvirtualization.AgentUpdateProperties{
		// 					Type: to.Ptr(armdesktopvirtualization.SessionHostComponentUpdateTypeScheduled),
		// 					MaintenanceWindowTimeZone: to.Ptr("Alaskan Standard Time"),
		// 					MaintenanceWindows: []*armdesktopvirtualization.MaintenanceWindowProperties{
		// 						{
		// 							DayOfWeek: to.Ptr(armdesktopvirtualization.DayOfWeekFriday),
		// 							Hour: to.Ptr[int32](7),
		// 						},
		// 						{
		// 							DayOfWeek: to.Ptr(armdesktopvirtualization.DayOfWeekSaturday),
		// 							Hour: to.Ptr[int32](8),
		// 					}},
		// 					UseSessionHostLocalTime: to.Ptr(false),
		// 				},
		// 				CloudPcResource: to.Ptr(false),
		// 				FriendlyName: to.Ptr("friendly"),
		// 				HostPoolType: to.Ptr(armdesktopvirtualization.HostPoolTypePooled),
		// 				LoadBalancerType: to.Ptr(armdesktopvirtualization.LoadBalancerTypeBreadthFirst),
		// 				MaxSessionLimit: to.Ptr[int32](999999),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				PersonalDesktopAssignmentType: to.Ptr(armdesktopvirtualization.PersonalDesktopAssignmentTypeAutomatic),
		// 				PreferredAppGroupType: to.Ptr(armdesktopvirtualization.PreferredAppGroupTypeDesktop),
		// 				RegistrationInfo: &armdesktopvirtualization.RegistrationInfo{
		// 					ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
		// 					RegistrationTokenOperation: to.Ptr(armdesktopvirtualization.RegistrationTokenOperationUpdate),
		// 					Token: to.Ptr("token"),
		// 				},
		// 				SsoClientID: to.Ptr("client"),
		// 				SsoClientSecretKeyVaultPath: to.Ptr("https://keyvault/secret"),
		// 				SsoSecretType: to.Ptr(armdesktopvirtualization.SSOSecretTypeSharedKey),
		// 				SsoadfsAuthority: to.Ptr("https://adfs"),
		// 				StartVMOnConnect: to.Ptr(false),
		// 				VMTemplate: to.Ptr("{json:json}"),
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
