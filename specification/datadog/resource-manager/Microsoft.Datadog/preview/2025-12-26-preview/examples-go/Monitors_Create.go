package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog/v2"
)

// Generated from example definition: 2025-12-26-preview/Monitors_Create.json
func ExampleMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatadog.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginCreate(ctx, "myResourceGroup", "myMonitor", armdatadog.MonitorResource{
		Location: to.Ptr("West US"),
		Properties: &armdatadog.MonitorProperties{
			DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
				Name:               to.Ptr("myOrg"),
				Cspm:               to.Ptr(false),
				EnterpriseAppID:    to.Ptr("00000000-0000-0000-0000-000000000000"),
				ID:                 to.Ptr("myOrg123"),
				LinkingAuthCode:    to.Ptr("someAuthCode"),
				LinkingClientID:    to.Ptr("00000000-0000-0000-0000-000000000000"),
				ResourceCollection: to.Ptr(false),
			},
			MarketplaceOfferDetails: &armdatadog.MarketplaceOfferDetails{
				PublisherID: to.Ptr("datadog1591740804488"),
				OfferID:     to.Ptr("dd_liftr_v3_decoupled"),
			},
			MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
			UserInfo: &armdatadog.UserInfo{
				Name:         to.Ptr("Alice"),
				EmailAddress: to.Ptr("alice@microsoft.com"),
				PhoneNumber:  to.Ptr("123-456-7890"),
			},
		},
		SKU: &armdatadog.ResourceSKU{
			Name: to.Ptr("free_Monthly"),
		},
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatadog.MonitorsClientCreateResponse{
	// 	MonitorResource: &armdatadog.MonitorResource{
	// 		Name: to.Ptr("myMonitor"),
	// 		Type: to.Ptr("Microsoft.Datadog/monitors"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armdatadog.MonitorProperties{
	// 			DatadogOrganizationProperties: &armdatadog.OrganizationProperties{
	// 				Name: to.Ptr("myOrg"),
	// 				ID: to.Ptr("myOrg123"),
	// 			},
	// 			MarketplaceOfferDetails: &armdatadog.MarketplaceOfferDetails{
	// 				PublisherID: to.Ptr("datadog1591740804488"),
	// 				OfferID: to.Ptr("dd_liftr_v3_decoupled"),
	// 			},
	// 			LiftrResourceCategory: to.Ptr(armdatadog.LiftrResourceCategoriesMonitorLogs),
	// 			LiftrResourcePreference: to.Ptr[int32](1),
	// 			MonitoringStatus: to.Ptr(armdatadog.MonitoringStatusEnabled),
	// 			ProvisioningState: to.Ptr(armdatadog.ProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armdatadog.ResourceSKU{
	// 			Name: to.Ptr("free_Monthly"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"Environment": to.Ptr("Dev"),
	// 		},
	// 	},
	// }
}
