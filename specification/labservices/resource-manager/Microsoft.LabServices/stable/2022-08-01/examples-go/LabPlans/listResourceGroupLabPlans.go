package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4c2cdccf6ca3281dd50ed8788ce1de2e0d480973/specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/listResourceGroupLabPlans.json
func ExampleLabPlansClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlabservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLabPlansClient().NewListByResourceGroupPager("testrg123", nil)
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
		// page.PagedLabPlans = armlabservices.PagedLabPlans{
		// 	Value: []*armlabservices.LabPlan{
		// 		{
		// 			Name: to.Ptr("testlabplan"),
		// 			Type: to.Ptr("Microsoft.LabServices/LabPlan"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labPlans/testlabplan"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armlabservices.LabPlanProperties{
		// 				DefaultAutoShutdownProfile: &armlabservices.AutoShutdownProfile{
		// 					DisconnectDelay: to.Ptr("PT5M"),
		// 					IdleDelay: to.Ptr("PT5M"),
		// 					NoConnectDelay: to.Ptr("PT5M"),
		// 					ShutdownOnDisconnect: to.Ptr(armlabservices.EnableStateEnabled),
		// 					ShutdownOnIdle: to.Ptr(armlabservices.ShutdownOnIdleModeUserAbsence),
		// 					ShutdownWhenNotConnected: to.Ptr(armlabservices.EnableStateEnabled),
		// 				},
		// 				DefaultConnectionProfile: &armlabservices.ConnectionProfile{
		// 					ClientRdpAccess: to.Ptr(armlabservices.ConnectionTypePublic),
		// 					ClientSSHAccess: to.Ptr(armlabservices.ConnectionTypePublic),
		// 					WebRdpAccess: to.Ptr(armlabservices.ConnectionTypeNone),
		// 					WebSSHAccess: to.Ptr(armlabservices.ConnectionTypeNone),
		// 				},
		// 				DefaultNetworkProfile: &armlabservices.LabPlanNetworkProfile{
		// 					SubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
		// 				},
		// 				SharedGalleryID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Compute/galleries/testsig"),
		// 				SupportInfo: &armlabservices.SupportInfo{
		// 					Email: to.Ptr("help@contoso.com"),
		// 					Instructions: to.Ptr("Contact support for help."),
		// 					Phone: to.Ptr("+1-202-555-0123"),
		// 					URL: to.Ptr("help.contoso.com"),
		// 				},
		// 				ProvisioningState: to.Ptr(armlabservices.ProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armlabservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T10:00:00Z"); return t}()),
		// 				CreatedBy: to.Ptr("identity123"),
		// 				CreatedByType: to.Ptr(armlabservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T09:12:28Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identity123"),
		// 				LastModifiedByType: to.Ptr(armlabservices.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
