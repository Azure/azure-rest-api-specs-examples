package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/ApplicationGroup_ListByResourceGroup.json
func ExampleApplicationGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationGroupsClient().NewListByResourceGroupPager("resourceGroup1", &armdesktopvirtualization.ApplicationGroupsClientListByResourceGroupOptions{Filter: to.Ptr("applicationGroupType eq 'RailApplication'"),
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
		// page.ApplicationGroupList = armdesktopvirtualization.ApplicationGroupList{
		// 	Value: []*armdesktopvirtualization.ApplicationGroup{
		// 		{
		// 			Name: to.Ptr("applicationGroup1"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1"),
		// 			Location: to.Ptr("centralus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armdesktopvirtualization.ApplicationGroupProperties{
		// 				Description: to.Ptr("des1"),
		// 				ApplicationGroupType: to.Ptr(armdesktopvirtualization.ApplicationGroupTypeRemoteApp),
		// 				CloudPcResource: to.Ptr(false),
		// 				FriendlyName: to.Ptr("friendly"),
		// 				HostPoolArmPath: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				WorkspaceArmPath: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/workspaces/workspace1"),
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
		// 			Name: to.Ptr("applicationGroup2"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup2"),
		// 			Location: to.Ptr("centralus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armdesktopvirtualization.ApplicationGroupProperties{
		// 				Description: to.Ptr("des1"),
		// 				ApplicationGroupType: to.Ptr(armdesktopvirtualization.ApplicationGroupTypeRemoteApp),
		// 				CloudPcResource: to.Ptr(false),
		// 				FriendlyName: to.Ptr("friendly"),
		// 				HostPoolArmPath: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
		// 				ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
		// 				WorkspaceArmPath: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/workspaces/workspace1"),
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
