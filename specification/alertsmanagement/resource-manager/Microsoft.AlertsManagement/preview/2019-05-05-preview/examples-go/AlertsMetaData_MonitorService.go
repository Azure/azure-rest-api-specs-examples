package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6d2438481021a94793b07b226df06d5f3c61d51d/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/AlertsMetaData_MonitorService.json
func ExampleAlertsClient_MetaData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertsClient().MetaData(ctx, armalertsmanagement.IdentifierMonitorServiceList, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertsMetaData = armalertsmanagement.AlertsMetaData{
	// 	Properties: &armalertsmanagement.MonitorServiceList{
	// 		MetadataIdentifier: to.Ptr(armalertsmanagement.MetadataIdentifierMonitorServiceList),
	// 		Data: []*armalertsmanagement.MonitorServiceDetails{
	// 			{
	// 				Name: to.Ptr("ActivityLog Administrative"),
	// 				DisplayName: to.Ptr("Activity Log - Administrative"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ActivityLog Autoscale"),
	// 				DisplayName: to.Ptr("Activity Log - Autoscale"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ActivityLog Policy"),
	// 				DisplayName: to.Ptr("Activity Log - Policy"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ActivityLog Recommendation"),
	// 				DisplayName: to.Ptr("Activity Log - Recommendation"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ActivityLog Security"),
	// 				DisplayName: to.Ptr("Activity Log - Security"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Application Insights"),
	// 				DisplayName: to.Ptr("Application Insights"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Azure Backup"),
	// 				DisplayName: to.Ptr("Azure Backup"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Custom"),
	// 				DisplayName: to.Ptr("Custom"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Data Box Edge"),
	// 				DisplayName: to.Ptr("Data Box Edge"),
	// 			},
	// 			{
	// 				Name: to.Ptr("VM Insights"),
	// 				DisplayName: to.Ptr("VM Insights"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Log Analytics"),
	// 				DisplayName: to.Ptr("Log Analytics"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Nagios"),
	// 				DisplayName: to.Ptr("NAGIOS"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Platform"),
	// 				DisplayName: to.Ptr("Platform"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Resource Health"),
	// 				DisplayName: to.Ptr("Resource Health"),
	// 			},
	// 			{
	// 				Name: to.Ptr("SCOM"),
	// 				DisplayName: to.Ptr("SCOM"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ServiceHealth"),
	// 				DisplayName: to.Ptr("Service Health"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Smart Detector"),
	// 				DisplayName: to.Ptr("SmartDetector"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Zabbix"),
	// 				DisplayName: to.Ptr("ZABBIX"),
	// 		}},
	// 	},
	// }
}
