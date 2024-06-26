package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationsList = armalertsmanagement.OperationsList{
		// 	Value: []*armalertsmanagement.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/register/action"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Subscription Registration Action"),
		// 				Operation: to.Ptr("Subscription Registration Action"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("Subscription"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/register/action"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the Microsoft Alerts Management"),
		// 				Operation: to.Ptr("Register subscription"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("register"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/alerts/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get all the alerts for the input filters."),
		// 				Operation: to.Ptr("Read alerts"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("alerts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/alerts/changestate/action"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Change the state of the alert."),
		// 				Operation: to.Ptr("Resolve alerts"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("alerts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/alerts/history/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get history of the alert"),
		// 				Operation: to.Ptr("Read alert history"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("alerts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/smartDetectorAlertRules/write"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Create or update Smart Detector alert rule in a given subscription"),
		// 				Operation: to.Ptr("Create Smart Detector alert rule"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("smartDetectorAlertRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/smartDetectorAlertRules/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get all the Smart Detector alert rules for the input filters"),
		// 				Operation: to.Ptr("Read Smart Detector alert rules"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("smartDetectorAlertRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/smartDetectorAlertRules/delete"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Delete Smart Detector alert rule in a given subscription"),
		// 				Operation: to.Ptr("Delete Smart Detector alert rule"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("smartDetectorAlertRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/resourceHealthAlertRules/write"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Create or update Resource Health alert rule in a given subscription"),
		// 				Operation: to.Ptr("Create Resource Health alert rule"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("resourceHealthAlertRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/resourceHealthAlertRules/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get all the Resource Health alert rules for the input filters"),
		// 				Operation: to.Ptr("Read Resource Health alert rules"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("resourceHealthAlertRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/resourceHealthAlertRules/delete"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Delete Resource Health alert rule in a given subscription"),
		// 				Operation: to.Ptr("Delete Resource Health alert rule"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("resourceHealthAlertRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/migrateFromSmartDetection/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get the status of an asynchronous Smart Detection to smart alerts migration process"),
		// 				Operation: to.Ptr("Get Smart Detection Migration status"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("migrateFromSmartDetection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/migrateFromSmartDetection/action"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Starts an asynchronous migration process of Smart Detection to smart alerts in an Application Insights resource"),
		// 				Operation: to.Ptr("Migrate From Smart Detection"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("migrateFromSmartDetection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/alertsSummary/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get the summary of alerts"),
		// 				Operation: to.Ptr("Read alerts summary"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("alertsSummary"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/smartGroups/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get all the smart groups for the input filters"),
		// 				Operation: to.Ptr("Read smart groups"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("smartGroups"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/smartGroups/changestate/action"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Change the state of the smart group"),
		// 				Operation: to.Ptr("Read smart groups"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("smartGroups"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/smartGroups/history/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get history of the smart group"),
		// 				Operation: to.Ptr("Read smart group history"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("smartGroups"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/actionRules/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get all the alert processing rules for the input filters."),
		// 				Operation: to.Ptr("Read action rules"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("actionRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/actionRules/write"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Create or update alert processing rule in a given subscription"),
		// 				Operation: to.Ptr("Write action rule"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("actionRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/actionRules/delete"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Delete alert processing rule in a given subscription."),
		// 				Operation: to.Ptr("Delete action rule"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("actionRules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/alertsMetaData/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Get alerts meta data for the input parameter."),
		// 				Operation: to.Ptr("Read alerts meta data"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("Microsoft.AlertsManagement/alertsMetaData"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AlertsManagement/Operations/read"),
		// 			Display: &armalertsmanagement.OperationDisplay{
		// 				Description: to.Ptr("Reads the operations provided"),
		// 				Operation: to.Ptr("Read operations"),
		// 				Provider: to.Ptr("Microsoft.AlertsManagement"),
		// 				Resource: to.Ptr("operations"),
		// 			},
		// 	}},
		// }
	}
}
