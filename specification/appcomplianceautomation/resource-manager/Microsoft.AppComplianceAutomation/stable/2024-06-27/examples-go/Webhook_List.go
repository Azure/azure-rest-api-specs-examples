package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Webhook_List.json
func ExampleWebhookClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebhookClient().NewListPager("testReportName", &armappcomplianceautomation.WebhookClientListOptions{SkipToken: to.Ptr("1"),
		Top:                   to.Ptr[int32](100),
		Select:                nil,
		Filter:                nil,
		Orderby:               nil,
		OfferGUID:             nil,
		ReportCreatorTenantID: nil,
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
		// page.WebhookResourceListResult = armappcomplianceautomation.WebhookResourceListResult{
		// 	Value: []*armappcomplianceautomation.WebhookResource{
		// 		{
		// 			Name: to.Ptr("testWebhookName"),
		// 			Type: to.Ptr("Microsfot.AppComplianceAutomation/reports/webhooks"),
		// 			ID: to.Ptr("/provider/Microsfot.AppComplianceAutomation/reports/testReportName/snapshots/testWebhookName"),
		// 			SystemData: &armappcomplianceautomation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 				CreatedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				CreatedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-14T22:34:55.449Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				LastModifiedByType: to.Ptr(armappcomplianceautomation.CreatedByTypeUser),
		// 			},
		// 			Properties: &armappcomplianceautomation.WebhookProperties{
		// 				ContentType: to.Ptr(armappcomplianceautomation.ContentTypeApplicationJSON),
		// 				DeliveryStatus: to.Ptr(armappcomplianceautomation.DeliveryStatusSucceeded),
		// 				EnableSSLVerification: to.Ptr(armappcomplianceautomation.EnableSSLVerificationTrue),
		// 				Events: []*armappcomplianceautomation.NotificationEvent{
		// 					to.Ptr(armappcomplianceautomation.NotificationEventGenerateSnapshotFailed)},
		// 					PayloadURL: to.Ptr("https://example.com"),
		// 					ProvisioningState: to.Ptr(armappcomplianceautomation.ProvisioningStateSucceeded),
		// 					SendAllEvents: to.Ptr(armappcomplianceautomation.SendAllEventsFalse),
		// 					Status: to.Ptr(armappcomplianceautomation.WebhookStatusEnabled),
		// 					TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					WebhookID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					WebhookKeyEnabled: to.Ptr(armappcomplianceautomation.WebhookKeyEnabledFalse),
		// 				},
		// 		}},
		// 	}
	}
}
