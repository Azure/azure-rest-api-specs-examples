package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/OperationsGet.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationsList = armdataboxedge.OperationsList{
		// 	Value: []*armdataboxedge.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/users/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the share users"),
		// 				Operation: to.Ptr("List share users"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("share users"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/users/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the share users"),
		// 				Operation: to.Ptr("List share users"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("share users"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/users/write"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the share users"),
		// 				Operation: to.Ptr("Creates or updates share users"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("share users"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/users/delete"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Deletes the share users"),
		// 				Operation: to.Ptr("Delete share users"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("share users"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/bandwidthSchedules/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the bandwidth schedules"),
		// 				Operation: to.Ptr("List bandwidth schedules"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("bandwidth schedules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/bandwidthSchedules/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the bandwidth schedules"),
		// 				Operation: to.Ptr("List bandwidth schedules"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("bandwidth schedules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/bandwidthSchedules/write"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the bandwidth schedules"),
		// 				Operation: to.Ptr("Creates or updates bandwidth schedules"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("bandwidth schedules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/bandwidthSchedules/delete"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Deletes the bandwidth schedules"),
		// 				Operation: to.Ptr("Delete bandwidth schedules"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("bandwidth schedules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/roles/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the ArmApiRes_roles"),
		// 				Operation: to.Ptr("List ArmApiRes_roles"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("ArmApiRes_roles"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/roles/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the ArmApiRes_roles"),
		// 				Operation: to.Ptr("List ArmApiRes_roles"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("ArmApiRes_roles"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/roles/write"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the ArmApiRes_roles"),
		// 				Operation: to.Ptr("Creates or updates ArmApiRes_roles"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("ArmApiRes_roles"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/roles/delete"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Deletes the ArmApiRes_roles"),
		// 				Operation: to.Ptr("Delete ArmApiRes_roles"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("ArmApiRes_roles"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/shares/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the shares"),
		// 				Operation: to.Ptr("List shares"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("shares"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/shares/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the shares"),
		// 				Operation: to.Ptr("List shares"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("shares"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/shares/write"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the shares"),
		// 				Operation: to.Ptr("Creates or updates shares"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("shares"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/shares/refresh/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("ArmApiDesc_action_refresh_shares"),
		// 				Operation: to.Ptr("ArmApiOp_action_refresh_shares"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("shares"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/shares/delete"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Deletes the shares"),
		// 				Operation: to.Ptr("Delete shares"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("shares"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/uploadCertificate/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Upload certificate for device registration"),
		// 				Operation: to.Ptr("Upload certificates"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/write"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the Data Box Edge devices"),
		// 				Operation: to.Ptr("Creates or updates Data Box Edge devices"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Data Box Edge devices"),
		// 				Operation: to.Ptr("List Data Box Edge devices"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/delete"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Deletes the Data Box Edge devices"),
		// 				Operation: to.Ptr("Delete Data Box Edge devices"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Data Box Edge devices"),
		// 				Operation: to.Ptr("List Data Box Edge devices"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Data Box Edge devices"),
		// 				Operation: to.Ptr("List Data Box Edge devices"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/write"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the Data Box Edge devices"),
		// 				Operation: to.Ptr("Creates or updates Data Box Edge devices"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/getExtendedInformation/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("ArmApiDesc_action_getExtendedInformation_dataBoxEdgeDevices"),
		// 				Operation: to.Ptr("ArmApiOp_action_getExtendedInformation_dataBoxEdgeDevices"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/extendedInformation/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Retrieves resource extended information"),
		// 				Operation: to.Ptr("Gets resource extended information"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/networkSettings/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the Device network settings"),
		// 				Operation: to.Ptr("List Device network settings"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Device network settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/securitySettings/update/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Update security settings"),
		// 				Operation: to.Ptr("Update security settings"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Device security settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/updateSummary/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the update summary"),
		// 				Operation: to.Ptr("List update summary"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("update summary"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/scanForUpdates/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Scan for updates"),
		// 				Operation: to.Ptr("Scan for updates"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/downloadUpdates/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Download Updates in device"),
		// 				Operation: to.Ptr("Download Updates"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/installUpdates/action"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Install Updates on device"),
		// 				Operation: to.Ptr("Install Updates"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge devices"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/jobs/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the jobs"),
		// 				Operation: to.Ptr("List jobs"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccountCredentials/write"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the storage account credentials"),
		// 				Operation: to.Ptr("Creates or updates storage account credentials"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("storage account credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccountCredentials/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the storage account credentials"),
		// 				Operation: to.Ptr("List storage account credentials"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("storage account credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccountCredentials/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the storage account credentials"),
		// 				Operation: to.Ptr("List storage account credentials"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("storage account credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccountCredentials/delete"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Deletes the storage account credentials"),
		// 				Operation: to.Ptr("Delete storage account credentials"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("storage account credentials"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/alerts/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the alerts"),
		// 				Operation: to.Ptr("List alerts"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("alerts"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/alerts/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Lists or gets the alerts"),
		// 				Operation: to.Ptr("List alerts"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("alerts"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armdataboxedge.OperationDisplay{
		// 				Description: to.Ptr("Gets the available Data Box Edge device level metrics"),
		// 				Operation: to.Ptr("Read Data Box Edge device metric definition"),
		// 				Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 				Resource: to.Ptr("Data Box Edge device"),
		// 			},
		// 			Origin: to.Ptr("system"),
		// 			Properties: &armdataboxedge.OperationProperties{
		// 				ServiceSpecification: &armdataboxedge.ServiceSpecification{
		// 					MetricSpecifications: []*armdataboxedge.MetricSpecificationV1{
		// 						{
		// 							Name: to.Ptr("NICReadThroughput"),
		// 							AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 							Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 							Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 								{
		// 									Name: to.Ptr("InstanceName"),
		// 									DisplayName: to.Ptr("Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("The read throughput of the network interface on the device in the reporting period for all volumes in the gateway."),
		// 							DisplayName: to.Ptr("Read Throughput (Network)"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 								to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 								to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 								to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 								SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 									to.Ptr(armdataboxedge.TimeGrainPT1M),
		// 									to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 									to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 									Unit: to.Ptr(armdataboxedge.MetricUnitBytesPerSecond),
		// 								},
		// 								{
		// 									Name: to.Ptr("NICWriteThroughput"),
		// 									AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 									Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 									Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 										{
		// 											Name: to.Ptr("InstanceName"),
		// 											DisplayName: to.Ptr("Name"),
		// 											ToBeExportedForShoebox: to.Ptr(true),
		// 									}},
		// 									DisplayDescription: to.Ptr("The write throughput of the network interface on the device in the reporting period for all volumes in the gateway."),
		// 									DisplayName: to.Ptr("Write Throughput (Network)"),
		// 									FillGapWithZero: to.Ptr(false),
		// 									SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 										to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 										to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 										to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 										SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 											to.Ptr(armdataboxedge.TimeGrainPT1M),
		// 											to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 											to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 											Unit: to.Ptr(armdataboxedge.MetricUnitBytesPerSecond),
		// 										},
		// 										{
		// 											Name: to.Ptr("CloudReadThroughputPerShare"),
		// 											AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 											Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 											Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 												{
		// 													Name: to.Ptr("Share"),
		// 													DisplayName: to.Ptr("Share"),
		// 													ToBeExportedForShoebox: to.Ptr(true),
		// 											}},
		// 											DisplayDescription: to.Ptr("The download throughput to Azure from a share during the reporting period."),
		// 											DisplayName: to.Ptr("Cloud Download Throughput (Share)"),
		// 											FillGapWithZero: to.Ptr(false),
		// 											SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 												to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 												to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 												to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 												SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 													to.Ptr(armdataboxedge.TimeGrainPT1M),
		// 													to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 													to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 													Unit: to.Ptr(armdataboxedge.MetricUnitBytesPerSecond),
		// 												},
		// 												{
		// 													Name: to.Ptr("CloudUploadThroughputPerShare"),
		// 													AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 													Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 													Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 														{
		// 															Name: to.Ptr("Share"),
		// 															DisplayName: to.Ptr("Share"),
		// 															ToBeExportedForShoebox: to.Ptr(true),
		// 													}},
		// 													DisplayDescription: to.Ptr("The upload throughput to Azure from a share during the reporting period."),
		// 													DisplayName: to.Ptr("Cloud Upload Throughput (Share)"),
		// 													FillGapWithZero: to.Ptr(false),
		// 													SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 														to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 														to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 														to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 														SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 															to.Ptr(armdataboxedge.TimeGrainPT1M),
		// 															to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 															to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 															Unit: to.Ptr(armdataboxedge.MetricUnitBytesPerSecond),
		// 														},
		// 														{
		// 															Name: to.Ptr("BytesUploadedToCloudPerShare"),
		// 															AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 															Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 															Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 																{
		// 																	Name: to.Ptr("Share"),
		// 																	DisplayName: to.Ptr("Share"),
		// 																	ToBeExportedForShoebox: to.Ptr(true),
		// 															}},
		// 															DisplayDescription: to.Ptr("The total number of bytes that is uploaded to Azure from a share during the reporting period."),
		// 															DisplayName: to.Ptr("Cloud Bytes Uploaded (Share)"),
		// 															FillGapWithZero: to.Ptr(false),
		// 															SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 																to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 																to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 																SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 																	to.Ptr(armdataboxedge.TimeGrainPT1M),
		// 																	to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 																	to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 																	Unit: to.Ptr(armdataboxedge.MetricUnitBytes),
		// 																},
		// 																{
		// 																	Name: to.Ptr("TotalCapacity"),
		// 																	AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																	Category: to.Ptr(armdataboxedge.MetricCategoryCapacity),
		// 																	Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 																	},
		// 																	DisplayDescription: to.Ptr("Total Capacity"),
		// 																	DisplayName: to.Ptr("Total Capacity"),
		// 																	FillGapWithZero: to.Ptr(false),
		// 																	SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 																		to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																		to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 																		to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 																		SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 																			to.Ptr(armdataboxedge.TimeGrainPT5M),
		// 																			to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 																			to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 																			Unit: to.Ptr(armdataboxedge.MetricUnitBytes),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("AvailableCapacity"),
		// 																			AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																			Category: to.Ptr(armdataboxedge.MetricCategoryCapacity),
		// 																			Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 																			},
		// 																			DisplayDescription: to.Ptr("The available capacity in bytes during the reporting period."),
		// 																			DisplayName: to.Ptr("Available Capacity"),
		// 																			FillGapWithZero: to.Ptr(false),
		// 																			SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 																				to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																				to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 																				to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 																				SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 																					to.Ptr(armdataboxedge.TimeGrainPT5M),
		// 																					to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 																					to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 																					Unit: to.Ptr(armdataboxedge.MetricUnitBytes),
		// 																				},
		// 																				{
		// 																					Name: to.Ptr("CloudUploadThroughput"),
		// 																					AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																					Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 																					Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 																					},
		// 																					DisplayDescription: to.Ptr("The cloud upload throughput during the reporting period."),
		// 																					DisplayName: to.Ptr("Cloud Upload Throughput"),
		// 																					FillGapWithZero: to.Ptr(false),
		// 																					SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 																						to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																						to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 																						to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 																						SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 																							to.Ptr(armdataboxedge.TimeGrainPT5M),
		// 																							to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 																							to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 																							Unit: to.Ptr(armdataboxedge.MetricUnitBytesPerSecond),
		// 																						},
		// 																						{
		// 																							Name: to.Ptr("CloudReadThroughput"),
		// 																							AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																							Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 																							Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 																							},
		// 																							DisplayDescription: to.Ptr("The cloud download throughput during the reporting period."),
		// 																							DisplayName: to.Ptr("Cloud Read Throughput"),
		// 																							FillGapWithZero: to.Ptr(false),
		// 																							SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 																								to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																								to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 																								to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 																								SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 																									to.Ptr(armdataboxedge.TimeGrainPT5M),
		// 																									to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 																									to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 																									Unit: to.Ptr(armdataboxedge.MetricUnitBytesPerSecond),
		// 																								},
		// 																								{
		// 																									Name: to.Ptr("BytesUploadedToCloud"),
		// 																									AggregationType: to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																									Category: to.Ptr(armdataboxedge.MetricCategoryTransaction),
		// 																									Dimensions: []*armdataboxedge.MetricDimensionV1{
		// 																									},
		// 																									DisplayDescription: to.Ptr("The total number of bytes that is uploaded to Azure from a device during the reporting period."),
		// 																									DisplayName: to.Ptr("Cloud Bytes Uploaded (Device)"),
		// 																									FillGapWithZero: to.Ptr(false),
		// 																									SupportedAggregationTypes: []*armdataboxedge.MetricAggregationType{
		// 																										to.Ptr(armdataboxedge.MetricAggregationTypeAverage),
		// 																										to.Ptr(armdataboxedge.MetricAggregationTypeMinimum),
		// 																										to.Ptr(armdataboxedge.MetricAggregationTypeMaximum)},
		// 																										SupportedTimeGrainTypes: []*armdataboxedge.TimeGrain{
		// 																											to.Ptr(armdataboxedge.TimeGrainPT5M),
		// 																											to.Ptr(armdataboxedge.TimeGrainPT15M),
		// 																											to.Ptr(armdataboxedge.TimeGrainPT1H)},
		// 																											Unit: to.Ptr(armdataboxedge.MetricUnitBytes),
		// 																									}},
		// 																								},
		// 																							},
		// 																						},
		// 																						{
		// 																							Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 																							Display: &armdataboxedge.OperationDisplay{
		// 																								Description: to.Ptr("Creates or updates the diagnostics setting for the resource"),
		// 																								Operation: to.Ptr("Write diagnostics setting"),
		// 																								Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 																								Resource: to.Ptr("Data Box Edge device"),
		// 																							},
		// 																							Origin: to.Ptr("system"),
		// 																						},
		// 																						{
		// 																							Name: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 																							Display: &armdataboxedge.OperationDisplay{
		// 																								Description: to.Ptr("Gets the diagnostic setting for the resource"),
		// 																								Operation: to.Ptr("Read diagnostics setting"),
		// 																								Provider: to.Ptr("Microsoft.DataBoxEdge"),
		// 																								Resource: to.Ptr("Data Box Edge device"),
		// 																							},
		// 																							Origin: to.Ptr("system"),
		// 																					}},
		// 																				}
	}
}
