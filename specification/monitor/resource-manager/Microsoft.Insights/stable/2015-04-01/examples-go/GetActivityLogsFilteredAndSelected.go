package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetActivityLogsFilteredAndSelected.json
func ExampleActivityLogsClient_NewListPager_getActivityLogsWithFilterAndSelect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewActivityLogsClient().NewListPager("eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and resourceGroupName eq 'MSSupportGroup'", &armmonitor.ActivityLogsClientListOptions{Select: to.Ptr("eventName,id,resourceGroupName,resourceProviderName,operationName,status,eventTimestamp,correlationId,submissionTimestamp,level")})
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
		// page.EventDataCollection = armmonitor.EventDataCollection{
		// 	Value: []*armmonitor.EventData{
		// 		{
		// 			CorrelationID: to.Ptr("1e121103-0ba6-4300-ac9d-952bb5d0c80f"),
		// 			EventName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("End request"),
		// 				Value: to.Ptr("EndRequest"),
		// 			},
		// 			EventTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:26.979Z"); return t}()),
		// 			ID: to.Ptr("/subscriptions/089bd33f-d4ec-47fe-8ba5-0753aa5c5b33/resourceGroups/MSSupportGroup/providers/microsoft.support/supporttickets/115012112305841/events/44ade6b4-3813-45e6-ae27-7420a95fa2f8/ticks/635574752669792776"),
		// 			Level: to.Ptr(armmonitor.EventLevelInformational),
		// 			OperationName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support/supporttickets/write"),
		// 				Value: to.Ptr("microsoft.support/supporttickets/write"),
		// 			},
		// 			ResourceGroupName: to.Ptr("MSSupportGroup"),
		// 			ResourceProviderName: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("microsoft.support"),
		// 				Value: to.Ptr("microsoft.support"),
		// 			},
		// 			Status: &armmonitor.LocalizableString{
		// 				LocalizedValue: to.Ptr("Succeeded"),
		// 				Value: to.Ptr("Succeeded"),
		// 			},
		// 			SubmissionTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-01-21T22:14:39.993Z"); return t}()),
		// 	}},
		// }
	}
}
