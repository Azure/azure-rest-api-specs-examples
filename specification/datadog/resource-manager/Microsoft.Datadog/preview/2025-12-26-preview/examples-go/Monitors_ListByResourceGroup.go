package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog/v2"
)

// Generated from example definition: 2025-12-26-preview/Monitors_ListByResourceGroup.json
func ExampleMonitorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armdatadog.MonitorsClientListByResourceGroupResponse{
		// 	MonitorResourceListResponse: armdatadog.MonitorResourceListResponse{
		// 		Value: []*armdatadog.MonitorResource{
		// 			{
		// 				Name: to.Ptr("myMonitor"),
		// 				Type: to.Ptr("Microsoft.Datadog/monitors"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armdatadog.MonitorProperties{
		// 					DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
		// 						Name: to.Ptr("myOrg"),
		// 						ID: to.Ptr("myOrg123"),
		// 					},
		// 					MarketplaceOfferDetails: &armdatadog.MarketplaceOfferDetails{
		// 						PublisherID: to.Ptr("datadog1591740804488"),
		// 						OfferID: to.Ptr("dd_liftr_v3_decoupled"),
		// 					},
		// 					LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
		// 					LiftrResourcePreference: to.Ptr[int32](1),
		// 					MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
		// 					ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"Environment": to.Ptr("Dev"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
