package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationEntityListResult = armstoragesync.OperationEntityListResult{
		// 	Value: []*armstoragesync.OperationEntity{
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/register/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Registers the subscription for the Storage Sync Provider"),
		// 				Operation: to.Ptr("Registers the Storage Sync Provider"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Storage Sync Provider"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/unregister/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Unregisters the subscription for the Storage Sync Provider"),
		// 				Operation: to.Ptr("Unregisters the Storage Sync Provider"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Storage Sync Provider"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/operations/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets a list of the Supported Operations"),
		// 				Operation: to.Ptr("Gets the Supported Operations"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Supported Operations"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Storage Sync Services"),
		// 				Operation: to.Ptr("Read Storage Sync Services"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Storage Sync Services"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/write"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Create or Update any Storage Sync Services"),
		// 				Operation: to.Ptr("Create or Update Storage Sync Services"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Storage Sync Services"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/delete"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Storage Sync Services"),
		// 				Operation: to.Ptr("Delete Storage Sync Services"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Storage Sync Services"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the available metrics for Storage Sync Services"),
		// 				Operation: to.Ptr("Read Storage Sync Services metric definitions"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Storage Sync Metrics"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 			Properties: &armstoragesync.OperationProperties{
		// 				ServiceSpecification: &armstoragesync.OperationResourceServiceSpecification{
		// 					MetricSpecifications: []*armstoragesync.OperationResourceMetricSpecification{
		// 						{
		// 							Name: to.Ptr("ServerSyncSessionResult"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Metric that logs a value of 1 each time the Server Endpoint successfully completes a Sync Session with the Cloud Endpoint"),
		// 							DisplayName: to.Ptr("Sync Session Result"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncSyncSessionAppliedFilesCount"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of Files synced"),
		// 							DisplayName: to.Ptr("Files Synced"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncSyncSessionPerItemErrorsCount"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of files failed to sync"),
		// 							DisplayName: to.Ptr("Files not syncing"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncBatchTransferredFileBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total file size transferred for Sync Sessions"),
		// 							DisplayName: to.Ptr("Bytes synced"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncServerHeartbeat"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("ServerName"),
		// 									DisplayName: to.Ptr("Server Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Metric that logs a value of 1 each time the resigtered server successfully records a heartbeat with the Cloud Endpoint"),
		// 							DisplayName: to.Ptr("Server Online Status"),
		// 							FillGapWithZero: to.Ptr(true),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncRecallIOTotalSizeBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("ServerName"),
		// 									DisplayName: to.Ptr("Server Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total size of data recalled by the server"),
		// 							DisplayName: to.Ptr("Cloud tiering recall"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncRecalledTotalNetworkBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerName"),
		// 									DisplayName: to.Ptr("Server Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Size of data recalled"),
		// 							DisplayName: to.Ptr("Cloud tiering recall size"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncRecallThroughputBytesPerSecond"),
		// 							AggregationType: to.Ptr("Average"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerName"),
		// 									DisplayName: to.Ptr("Server Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Size of data recall throughput"),
		// 							DisplayName: to.Ptr("Cloud tiering recall throughput"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("BytesPerSecond"),
		// 						},
		// 						{
		// 							Name: to.Ptr("StorageSyncRecalledNetworkBytesByApplication"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerName"),
		// 									DisplayName: to.Ptr("Server Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ApplicationName"),
		// 									DisplayName: to.Ptr("Application Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Size of data recalled by application"),
		// 							DisplayName: to.Ptr("Cloud tiering recall size by application"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateLinkResources/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Private Link Resources"),
		// 				Operation: to.Ptr("Read Private Link Resources"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Link Resources"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateEndpointConnections/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Private Endpoint Connections"),
		// 				Operation: to.Ptr("Read Private Endpoint Connections"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Endpoint Connections"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateEndpointConnections/write"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Create or Update any Private Endpoint Connections"),
		// 				Operation: to.Ptr("Create or Update Private Endpoint Connections"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Endpoint Connections"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateEndpointConnections/delete"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Private Endpoint Connections"),
		// 				Operation: to.Ptr("Delete Private Endpoint Connections"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Endpoint Connections"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Validate any Private Endpoint ConnectionProxies"),
		// 				Operation: to.Ptr("Validate Private Endpoint ConnectionProxies"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Endpoint ConnectionProxies"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateEndpointConnectionProxies/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Private Endpoint ConnectionProxies"),
		// 				Operation: to.Ptr("Read Private Endpoint ConnectionProxies"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Endpoint ConnectionProxies"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateEndpointConnectionProxies/write"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Create or Update any Private Endpoint ConnectionProxies"),
		// 				Operation: to.Ptr("Create or Update Private Endpoint ConnectionProxies"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Endpoint ConnectionProxies"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/privateEndpointConnectionProxies/delete"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Private Endpoint ConnectionProxies"),
		// 				Operation: to.Ptr("Delete Private Endpoint ConnectionProxies"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Private Endpoint ConnectionProxies"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Sync Groups"),
		// 				Operation: to.Ptr("Read Sync Groups"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Sync Groups"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/write"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Create or Update any Sync Groups"),
		// 				Operation: to.Ptr("Create or Update Sync Groups"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Sync Groups"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/delete"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Sync Groups"),
		// 				Operation: to.Ptr("Delete Sync Groups"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Sync Groups"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the available metrics for Sync Groups"),
		// 				Operation: to.Ptr("Read Sync Groups metric definitions"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Sync Group Metrics"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 			Properties: &armstoragesync.OperationProperties{
		// 				ServiceSpecification: &armstoragesync.OperationResourceServiceSpecification{
		// 					MetricSpecifications: []*armstoragesync.OperationResourceMetricSpecification{
		// 						{
		// 							Name: to.Ptr("SyncGroupSyncSessionAppliedFilesCount"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of Files synced"),
		// 							DisplayName: to.Ptr("Files Synced"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("SyncGroupSyncSessionPerItemErrorsCount"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of files failed to sync"),
		// 							DisplayName: to.Ptr("Files not syncing"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("SyncGroupBatchTransferredFileBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("SyncGroupName"),
		// 									DisplayName: to.Ptr("Sync Group Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total file size transferred for Sync Sessions"),
		// 							DisplayName: to.Ptr("Bytes synced"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Cloud Endpoints"),
		// 				Operation: to.Ptr("Read Cloud Endpoints"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/write"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Create or Update any Cloud Endpoints"),
		// 				Operation: to.Ptr("Create or Update Cloud Endpoints"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/delete"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Cloud Endpoints"),
		// 				Operation: to.Ptr("Delete Cloud Endpoints"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/prebackup/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Call this action before backup"),
		// 				Operation: to.Ptr("prebackup"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/postbackup/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Call this action after backup"),
		// 				Operation: to.Ptr("postbackup"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/prerestore/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Call this action before restore"),
		// 				Operation: to.Ptr("prerestore"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/postrestore/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Call this action after restore"),
		// 				Operation: to.Ptr("postrestore"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/restoreheartbeat/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Restore heartbeat"),
		// 				Operation: to.Ptr("restoreheartbeat"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/operationresults/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the status of an asynchronous backup/restore operation"),
		// 				Operation: to.Ptr("Read cloudEndpoints/operationresults"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/cloudEndpoints/triggerChangeDetection/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Call this action to trigger detection of changes on a cloud endpoint's file share"),
		// 				Operation: to.Ptr("triggerChangeDetection"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Cloud Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/serverEndpoints/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Server Endpoints"),
		// 				Operation: to.Ptr("Read Server Endpoints"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Server Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/serverEndpoints/write"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Create or Update any Server Endpoints"),
		// 				Operation: to.Ptr("Create or Update Server Endpoints"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Server Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/serverEndpoints/delete"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Server Endpoints"),
		// 				Operation: to.Ptr("Delete Server Endpoints"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Server Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/serverEndpoints/recallAction/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Call this action to recall files to a server"),
		// 				Operation: to.Ptr("recallAction"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Server Endpoints"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/syncGroups/serverEndpoints/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the available metrics for Server Endpoints"),
		// 				Operation: to.Ptr("Read Server Endpoints metric definitions"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Server Endpoint Metrics"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 			Properties: &armstoragesync.OperationProperties{
		// 				ServiceSpecification: &armstoragesync.OperationResourceServiceSpecification{
		// 					MetricSpecifications: []*armstoragesync.OperationResourceMetricSpecification{
		// 						{
		// 							Name: to.Ptr("ServerEndpointSyncSessionAppliedFilesCount"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of Files synced"),
		// 							DisplayName: to.Ptr("Files Synced"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("ServerEndpointSyncSessionPerItemErrorsCount"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Count of files failed to sync"),
		// 							DisplayName: to.Ptr("Files not syncing"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("ServerEndpointBatchTransferredFileBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("ServerEndpointName"),
		// 									DisplayName: to.Ptr("Server Endpoint Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("SyncDirection"),
		// 									DisplayName: to.Ptr("Sync Direction"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total file size transferred for Sync Sessions"),
		// 							DisplayName: to.Ptr("Bytes synced"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/registeredServers/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Registered Server"),
		// 				Operation: to.Ptr("Read Registered Server"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Registered Server"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/registeredServers/write"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Create or Update any Registered Server"),
		// 				Operation: to.Ptr("Create or Update Registered Server"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Registered Server"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/registeredServers/delete"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Registered Server"),
		// 				Operation: to.Ptr("Delete Registered Server"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Registered Server"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/registeredServers/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the available metrics for Registered Server"),
		// 				Operation: to.Ptr("Read Registered Server metric definitions"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Registered Server Metrics"),
		// 			},
		// 			Origin: to.Ptr("System"),
		// 			Properties: &armstoragesync.OperationProperties{
		// 				ServiceSpecification: &armstoragesync.OperationResourceServiceSpecification{
		// 					MetricSpecifications: []*armstoragesync.OperationResourceMetricSpecification{
		// 						{
		// 							Name: to.Ptr("ServerHeartbeat"),
		// 							AggregationType: to.Ptr("Maximum"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("ServerResourceId"),
		// 									DisplayName: to.Ptr("Registered Server Id"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerName"),
		// 									DisplayName: to.Ptr("Server Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Metric that logs a value of 1 each time the resigtered server successfully records a heartbeat with the Cloud Endpoint"),
		// 							DisplayName: to.Ptr("Server Online Status"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Count"),
		// 						},
		// 						{
		// 							Name: to.Ptr("ServerRecallIOTotalSizeBytes"),
		// 							AggregationType: to.Ptr("Total"),
		// 							Dimensions: []*armstoragesync.OperationResourceMetricSpecificationDimension{
		// 								{
		// 									Name: to.Ptr("ServerResourceId"),
		// 									DisplayName: to.Ptr("Registered Server Id"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 								},
		// 								{
		// 									Name: to.Ptr("ServerName"),
		// 									DisplayName: to.Ptr("Server Name"),
		// 									ToBeExportedForShoebox: to.Ptr(true),
		// 							}},
		// 							DisplayDescription: to.Ptr("Total size of data recalled by the server"),
		// 							DisplayName: to.Ptr("Cloud tiering recall"),
		// 							FillGapWithZero: to.Ptr(false),
		// 							Unit: to.Ptr("Bytes"),
		// 					}},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/workflows/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Read Workflows"),
		// 				Operation: to.Ptr("Read Workflows"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/workflows/operationresults/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the status of an asynchronous operation"),
		// 				Operation: to.Ptr("Read workflows/operationresults"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/storageSyncServices/workflows/operations/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the status of an asynchronous operation"),
		// 				Operation: to.Ptr("Read workflows/operations"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/locations/checkNameAvailability/action"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Checks that storage sync service name is valid and is not in use."),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Name Availability"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/locations/workflows/operations/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the status of an asynchronous operation"),
		// 				Operation: to.Ptr("Read locations/workflows/operations"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("Workflow"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/locations/operationresults/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the result for an asynchronous operation"),
		// 				Operation: to.Ptr("Read locations/operationresults"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("operationresults"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 		},
		// 		{
		// 			Name: to.Ptr("microsoft.storagesync/locations/operations/read"),
		// 			Display: &armstoragesync.OperationDisplayInfo{
		// 				Description: to.Ptr("Gets the status for an azure asynchronous operation"),
		// 				Operation: to.Ptr("Read locations/operations"),
		// 				Provider: to.Ptr("microsoft.storagesync"),
		// 				Resource: to.Ptr("operations"),
		// 			},
		// 			Origin: to.Ptr("User"),
		// 	}},
		// }
	}
}
