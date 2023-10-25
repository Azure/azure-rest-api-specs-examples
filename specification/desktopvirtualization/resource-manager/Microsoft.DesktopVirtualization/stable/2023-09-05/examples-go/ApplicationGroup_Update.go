package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ApplicationGroup_Update.json
func ExampleApplicationGroupsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationGroupsClient().Update(ctx, "resourceGroup1", "applicationGroup1", &armdesktopvirtualization.ApplicationGroupsClientUpdateOptions{ApplicationGroup: &armdesktopvirtualization.ApplicationGroupPatch{
		Properties: &armdesktopvirtualization.ApplicationGroupPatchProperties{
			Description:  to.Ptr("des1"),
			FriendlyName: to.Ptr("friendly"),
			ShowInFeed:   to.Ptr(true),
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGroup = armdesktopvirtualization.ApplicationGroup{
	// 	Name: to.Ptr("applicationGroup1"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1"),
	// 	Location: to.Ptr("centralus"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armdesktopvirtualization.ApplicationGroupProperties{
	// 		Description: to.Ptr("des1"),
	// 		ApplicationGroupType: to.Ptr(armdesktopvirtualization.ApplicationGroupTypeRemoteApp),
	// 		CloudPcResource: to.Ptr(false),
	// 		FriendlyName: to.Ptr("friendly"),
	// 		HostPoolArmPath: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1"),
	// 		ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
	// 		ShowInFeed: to.Ptr(true),
	// 		WorkspaceArmPath: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/workspaces/workspace1"),
	// 	},
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// }
}
