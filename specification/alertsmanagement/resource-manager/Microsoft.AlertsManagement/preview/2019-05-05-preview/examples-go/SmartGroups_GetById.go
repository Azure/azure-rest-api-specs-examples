package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_GetById.json
func ExampleSmartGroupsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSmartGroupsClient().GetByID(ctx, "603675da-9851-4b26-854a-49fc53d32715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SmartGroup = armalertsmanagement.SmartGroup{
	// 	Name: to.Ptr("cpu alert"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/smartGroups"),
	// 	ID: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/providers/Microsoft.AlertsManagement/smartGroups/a808445e-bb38-4751-85c2-1b109ccc1059"),
	// 	Properties: &armalertsmanagement.SmartGroupProperties{
	// 		AlertSeverities: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Sev3"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		AlertStates: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("New"),
	// 				Count: to.Ptr[int64](1941),
	// 			},
	// 			{
	// 				Name: to.Ptr("Acknowledged"),
	// 				Count: to.Ptr[int64](1),
	// 		}},
	// 		AlertsCount: to.Ptr[int64](1942),
	// 		LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T06:30:09.000Z"); return t}()),
	// 		LastModifiedUserName: to.Ptr("System"),
	// 		MonitorConditions: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Fired"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		MonitorServices: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Application Insights"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		ResourceGroups: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("alertscorrelationrg"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		ResourceTypes: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("components"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		Resources: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/resourcegroups/alertscorrelationrg/providers/microsoft.insights/components/alertscorrelationworkerrole_int"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		Severity: to.Ptr(armalertsmanagement.SeveritySev3),
	// 		SmartGroupState: to.Ptr(armalertsmanagement.StateNew),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-06T12:35:09.000Z"); return t}()),
	// 	},
	// }
}
