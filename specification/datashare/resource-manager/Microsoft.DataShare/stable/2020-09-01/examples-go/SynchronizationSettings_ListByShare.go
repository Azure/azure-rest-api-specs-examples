package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/SynchronizationSettings_ListByShare.json
func ExampleSynchronizationSettingsClient_NewListBySharePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSynchronizationSettingsClient().NewListBySharePager("SampleResourceGroup", "Account1", "Share1", &armdatashare.SynchronizationSettingsClientListByShareOptions{SkipToken: nil})
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
		// page.SynchronizationSettingList = armdatashare.SynchronizationSettingList{
		// 	Value: []armdatashare.SynchronizationSettingClassification{
		// 		&armdatashare.ScheduledSynchronizationSetting{
		// 			Name: to.Ptr("SynchronizationSetting1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/shares/synchronizationSettings"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shares/Share1/synchronizationSettings/SynchronizationSetting1"),
		// 			Kind: to.Ptr(armdatashare.SynchronizationSettingKindScheduleBased),
		// 			Properties: &armdatashare.ScheduledSynchronizationSettingProperties{
		// 				RecurrenceInterval: to.Ptr(armdatashare.RecurrenceIntervalDay),
		// 				SynchronizationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-14T04:47:52.961Z"); return t}()),
		// 				UserName: to.Ptr("John Smith"),
		// 			},
		// 	}},
		// }
	}
}
