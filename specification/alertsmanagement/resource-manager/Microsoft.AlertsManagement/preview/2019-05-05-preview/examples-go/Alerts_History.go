package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_History.json
func ExampleAlertsClient_GetHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().GetHistory(ctx, "66114d64-d9d9-478b-95c9-b789d6502100", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertModification = armalertsmanagement.AlertModification{
	// 	Name: to.Ptr("CPU Alert"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/alerts"),
	// 	ID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/providers/Microsoft.AlertsManagement/alerts/66114d64-d9d9-478b-95c9-b789d6502100/history/default"),
	// 	Properties: &armalertsmanagement.AlertModificationProperties{
	// 		AlertID: to.Ptr("66114d64-d9d9-478b-95c9-b789d6502100"),
	// 		Modifications: []*armalertsmanagement.AlertModificationItem{
	// 			{
	// 				Description: to.Ptr("State changed from 'New' to 'Acknowledged'"),
	// 				Comments: to.Ptr("Acknowledging alert"),
	// 				ModificationEvent: to.Ptr(armalertsmanagement.AlertModificationEventStateChange),
	// 				ModifiedAt: to.Ptr("2018-06-13T06:14:15.7378737Z"),
	// 				ModifiedBy: to.Ptr("vikramm@microsoft.com"),
	// 				NewValue: to.Ptr("Acknowledged"),
	// 				OldValue: to.Ptr("New"),
	// 			},
	// 			{
	// 				Description: to.Ptr("New Alert Object is created"),
	// 				Comments: to.Ptr(""),
	// 				ModificationEvent: to.Ptr(armalertsmanagement.AlertModificationEventAlertCreated),
	// 				ModifiedAt: to.Ptr("2018-06-13T06:09:01Z"),
	// 				ModifiedBy: to.Ptr("System"),
	// 				NewValue: to.Ptr(""),
	// 				OldValue: to.Ptr(""),
	// 		}},
	// 	},
	// }
}
