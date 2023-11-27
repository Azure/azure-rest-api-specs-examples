package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ShareSubscriptions_ListSourceShareSynchronizationSettings.json
func ExampleShareSubscriptionsClient_NewListSourceShareSynchronizationSettingsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewShareSubscriptionsClient().NewListSourceShareSynchronizationSettingsPager("SampleResourceGroup", "Account1", "ShareSub1", &armdatashare.ShareSubscriptionsClientListSourceShareSynchronizationSettingsOptions{SkipToken: nil})
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
		// page.SourceShareSynchronizationSettingList = armdatashare.SourceShareSynchronizationSettingList{
		// 	Value: []armdatashare.SourceShareSynchronizationSettingClassification{
		// 		&armdatashare.ScheduledSourceSynchronizationSetting{
		// 			Kind: to.Ptr(armdatashare.SourceShareSynchronizationSettingKindScheduleBased),
		// 			Properties: &armdatashare.ScheduledSourceShareSynchronizationSettingProperties{
		// 				RecurrenceInterval: to.Ptr(armdatashare.RecurrenceIntervalHour),
		// 				SynchronizationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-15T19:45:58.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
