package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountBatchConfigurations_List.json
func ExampleIntegrationAccountBatchConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationAccountBatchConfigurationsClient().NewListPager("testResourceGroup", "testIntegrationAccount", nil)
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
		// page.BatchConfigurationCollection = armlogic.BatchConfigurationCollection{
		// 	Value: []*armlogic.BatchConfiguration{
		// 		{
		// 			Name: to.Ptr("testBatchConfiguration"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationAccounts/batchConfigurations"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationAccounts/testIntegrationAccount/batchConfigurations/testBatchConfiguration"),
		// 			Properties: &armlogic.BatchConfigurationProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T06:14:16.704Z"); return t}()),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T06:14:16.651Z"); return t}()),
		// 				BatchGroupName: to.Ptr("DEFAULT"),
		// 				ReleaseCriteria: &armlogic.BatchReleaseCriteria{
		// 					BatchSize: to.Ptr[int32](234567),
		// 					MessageCount: to.Ptr[int32](10),
		// 					Recurrence: &armlogic.WorkflowTriggerRecurrence{
		// 						Frequency: to.Ptr(armlogic.RecurrenceFrequencyMinute),
		// 						Interval: to.Ptr[int32](1),
		// 						StartTime: to.Ptr("2017-03-24T11:43:00"),
		// 						TimeZone: to.Ptr("India Standard Time"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
