package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_History.json
func ExampleSmartGroupsClient_GetHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSmartGroupsClient().GetHistory(ctx, "a808445e-bb38-4751-85c2-1b109ccc1059", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SmartGroupModification = armalertsmanagement.SmartGroupModification{
	// 	Name: to.Ptr("cpu alert"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/alerts"),
	// 	ID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/providers/Microsoft.AlertsManagement/smartGroups/a808445e-bb38-4751-85c2-1b109ccc1059/history/default"),
	// 	Properties: &armalertsmanagement.SmartGroupModificationProperties{
	// 		Modifications: []*armalertsmanagement.SmartGroupModificationItem{
	// 			{
	// 				Description: to.Ptr("New Smart Group is created"),
	// 				Comments: to.Ptr(""),
	// 				ModificationEvent: to.Ptr(armalertsmanagement.SmartGroupModificationEventSmartGroupCreated),
	// 				ModifiedAt: to.Ptr("2018-06-06T12:35:09Z"),
	// 				ModifiedBy: to.Ptr("System"),
	// 				NewValue: to.Ptr(""),
	// 				OldValue: to.Ptr(""),
	// 		}},
	// 		SmartGroupID: to.Ptr("a808445e-bb38-4751-85c2-1b109ccc1059"),
	// 	},
	// }
}
