package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armworkloadssapvirtualinstance.OperationsClientListResponse{
		// 	OperationListResult: armworkloadssapvirtualinstance.OperationListResult{
		// 		Value: []*armworkloadssapvirtualinstance.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/RegisteredSubscriptions/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("RegisteredSubscriptions"),
		// 					Operation: to.Ptr("Gets/Lists registered subscriptions"),
		// 					Description: to.Ptr("Reads registered subscriptions"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 					Operation: to.Ptr("WordpressInstances_List"),
		// 					Description: to.Ptr("Lists WordpressInstances resources under a phpWorkload resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 					Operation: to.Ptr("WordpressInstances_Get"),
		// 					Description: to.Ptr("Gets a WordpressInstances resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 					Operation: to.Ptr("WordpressInstances_CreateOrUpdate"),
		// 					Description: to.Ptr("Create or updated WordpressInstances resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 					Operation: to.Ptr("WordpressInstances_Delete"),
		// 					Description: to.Ptr("Delete WordpressInstances resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/skus/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("Skus"),
		// 					Operation: to.Ptr("Gets the list of Microsoft.Workloads SKUs available for your Subscription"),
		// 					Description: to.Ptr("Gets the list of Microsoft.Workloads SKUs available for your Subscription"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/Operations/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("Operations"),
		// 					Operation: to.Ptr("read_Operations"),
		// 					Description: to.Ptr("read Operations"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/register/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("Microsoft.Workloads"),
		// 					Operation: to.Ptr("Register the Microsoft.Workloads"),
		// 					Description: to.Ptr("Register the subscription for Microsoft.Workloads"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/unregister/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("Microsoft.Workloads"),
		// 					Operation: to.Ptr("Unregister the Microsoft.Workloads"),
		// 					Description: to.Ptr("Unregister the subscription for Microsoft.Workloads"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/Locations/OperationStatuses/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("Locations/OperationStatuses"),
		// 					Operation: to.Ptr("read_OperationStatuses"),
		// 					Description: to.Ptr("read OperationStatuses"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/Locations/OperationStatuses/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("Locations/OperationStatuses"),
		// 					Operation: to.Ptr("write_OperationStatuses"),
		// 					Description: to.Ptr("write OperationStatuses"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads"),
		// 					Operation: to.Ptr("PhpWorkloads_ListBySubscription"),
		// 					Description: to.Ptr("Lists phpWorkload resources in a subscription"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads"),
		// 					Operation: to.Ptr("PhpWorkloads_ListByResourceGroup"),
		// 					Description: to.Ptr("Lists phpWorkload resources in a resource group"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads"),
		// 					Operation: to.Ptr("PhpWorkloads_Get"),
		// 					Description: to.Ptr("Gets a phpWorkload resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads"),
		// 					Operation: to.Ptr("PhpWorkloads_CreateOrUpdate"),
		// 					Description: to.Ptr("Create or updated phpWorkloads resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads"),
		// 					Operation: to.Ptr("PhpWorkloads_Delete"),
		// 					Description: to.Ptr("Delete phpWorkloads resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/phpWorkloads/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("phpWorkloads"),
		// 					Operation: to.Ptr("PhpWorkloads_Update"),
		// 					Description: to.Ptr("Update PHP workload resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 					Operation: to.Ptr("SAPApplicationServerInstances_Get"),
		// 					Description: to.Ptr("Gets the SAP Application Server Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 					Operation: to.Ptr("SAPApplicationServerInstances_Create"),
		// 					Description: to.Ptr("Puts the SAP Application Server Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 					Operation: to.Ptr("SAPApplicationServerInstances_Delete"),
		// 					Description: to.Ptr("Deletes the SAP Application Server Instance. <br><br>This operation will be used by service only. Delete by end user will return a Bad Request error."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 					Operation: to.Ptr("SAPApplicationServerInstances_Update"),
		// 					Description: to.Ptr("Puts the SAP Application Server Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 					Operation: to.Ptr("SAPApplicationServerInstances_List"),
		// 					Description: to.Ptr("Lists the SAP Application server Instances in an SVI."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/start/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 					Operation: to.Ptr("SAPApplicationServerInstances_StartInstance"),
		// 					Description: to.Ptr("Starts the SAP Application server Instance in an SVI."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/stop/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 					Operation: to.Ptr("SAPApplicationServerInstances_StopInstance"),
		// 					Description: to.Ptr("Stops the SAP Application server Instance in an SVI."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 					Operation: to.Ptr("SAPCentralInstances_Get"),
		// 					Description: to.Ptr("Gets the SAP Central Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 					Operation: to.Ptr("SAPCentralInstances_Create"),
		// 					Description: to.Ptr("Puts the SAP Central Instance. <br><br>This will be used by service only. PUT by end user will return a Bad Request error."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 					Operation: to.Ptr("SAPCentralInstances_Delete"),
		// 					Description: to.Ptr("Deletes the SAP Central Instance. <br><br>This will be used by service only. Delete by end user will return a Bad Request error."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 					Operation: to.Ptr("SAPCentralInstances_Update"),
		// 					Description: to.Ptr("Updates the SAP Central Instance. <br><br>This can be used to update tags."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 					Operation: to.Ptr("SAPCentralInstances_List"),
		// 					Description: to.Ptr("Lists the SAP Central Instances in an SVI."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/start/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 					Operation: to.Ptr("SAPCentralInstances_StartInstance"),
		// 					Description: to.Ptr("Starts the SAP Central server Instance in an SVI."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/stop/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 					Operation: to.Ptr("SAPCentralInstances_StopInstance"),
		// 					Description: to.Ptr("Stops the SAP Central server Instance in an SVI."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 					Operation: to.Ptr("SAPDatabaseInstances_Get"),
		// 					Description: to.Ptr("Gets the SAP Database Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 					Operation: to.Ptr("SAPDatabaseInstances_Create"),
		// 					Description: to.Ptr("Puts the SAP Database Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 					Operation: to.Ptr("SAPDatabaseInstances_Delete"),
		// 					Description: to.Ptr("Deletes the SAP Database Instance. <br><br>This will be used by service only. Delete by end user will return a Bad Request error."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 					Operation: to.Ptr("SAPDatabaseInstances_Update"),
		// 					Description: to.Ptr("Puts the SAP Database Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 					Operation: to.Ptr("SAPDatabaseInstances_List"),
		// 					Description: to.Ptr("Lists the SAP Database Instances in an SVI."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/start/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 					Operation: to.Ptr("SAPDatabaseInstances_StartInstance"),
		// 					Description: to.Ptr("Starts the database instance of the SAP system."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/stop/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 					Operation: to.Ptr("SAPDatabaseInstances_StopInstance"),
		// 					Description: to.Ptr("Stops the database instance of the SAP system."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors"),
		// 					Operation: to.Ptr("monitors_List"),
		// 					Description: to.Ptr("Gets a list of SAP monitors in the specified subscription. The operations returns various properties of each SAP monitor."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors"),
		// 					Operation: to.Ptr("monitors_ListByResourceGroup"),
		// 					Description: to.Ptr("Gets a list of SAP monitors in the specified resource group."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors"),
		// 					Operation: to.Ptr("monitors_Get"),
		// 					Description: to.Ptr("Gets properties of a SAP monitor for the specified subscription, resource group, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors"),
		// 					Operation: to.Ptr("monitors_Create"),
		// 					Description: to.Ptr("Creates a SAP monitor for the specified subscription, resource group, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors"),
		// 					Operation: to.Ptr("monitors_Delete"),
		// 					Description: to.Ptr("Deletes a SAP monitor with the specified subscription, resource group, and monitor name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors"),
		// 					Operation: to.Ptr("monitors_Update"),
		// 					Description: to.Ptr("Patches the Tags field of a SAP monitor for the specified subscription, resource group, and monitor name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/providerInstances"),
		// 					Operation: to.Ptr("ProviderInstances_List"),
		// 					Description: to.Ptr("Gets a list of provider instances in the specified SAP monitor. The operations returns various properties of each provider instances."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/providerInstances"),
		// 					Operation: to.Ptr("ProviderInstances_Get"),
		// 					Description: to.Ptr("Gets properties of a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/providerInstances"),
		// 					Operation: to.Ptr("ProviderInstances_Create"),
		// 					Description: to.Ptr("Creates a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/providerInstances"),
		// 					Operation: to.Ptr("ProviderInstances_Delete"),
		// 					Description: to.Ptr("Deletes a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/alerts/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/alerts"),
		// 					Operation: to.Ptr("Alerts_List"),
		// 					Description: to.Ptr("Gets a list of alert instances in the specified SAP monitor. The operations returns various properties of each provider instances."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/alerts/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/alerts"),
		// 					Operation: to.Ptr("Alerts_Get"),
		// 					Description: to.Ptr("Gets properties of a alert for the specified subscription, resource group, Monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/alerts/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/alerts"),
		// 					Operation: to.Ptr("Alerts_Create"),
		// 					Description: to.Ptr("Creates a alert for the specified subscription, resource group, Monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/alerts/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/alerts"),
		// 					Operation: to.Ptr("Alerts_Delete"),
		// 					Description: to.Ptr("Deletes a alert for the specified subscription, resource group, Monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/alertTemplates/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/alertTemplates"),
		// 					Operation: to.Ptr("AlertTemplates_List"),
		// 					Description: to.Ptr("Gets properties of an alert template for the specified subscription, resource group, SAP monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/alertTemplates/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/alertTemplates"),
		// 					Operation: to.Ptr("AlertTemplates_Get"),
		// 					Description: to.Ptr("Gets properties of a alert for the specified subscription, resource group, Monitor name, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 					Operation: to.Ptr("SapLandscapeMonitor_List"),
		// 					Description: to.Ptr("Gets a list of properties of a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 					Operation: to.Ptr("SapLandscapeMonitor_Get"),
		// 					Description: to.Ptr("Gets properties of a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 					Operation: to.Ptr("SapLandscapeMonitor_Create"),
		// 					Description: to.Ptr("Creates a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 					Operation: to.Ptr("SapLandscapeMonitor_Delete"),
		// 					Description: to.Ptr("Deletes a SAP Landscape monitor configuration with the specified subscription, resource group, and monitor name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_Get"),
		// 					Description: to.Ptr("Gets an SAP Virtual Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_Create"),
		// 					Description: to.Ptr("Creates an SAP Virtual Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_Delete"),
		// 					Description: to.Ptr("Deletes an SAP Virtual Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_Update"),
		// 					Description: to.Ptr("Updates an SAP Virtual Instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_ListByResourceGroup"),
		// 					Description: to.Ptr("Gets all SAP Virtual Instances in a resource group."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_ListBySubscription"),
		// 					Description: to.Ptr("Gets all SAP Virtual Instances in the subscription."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/start/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_Start"),
		// 					Description: to.Ptr("Starts the SAP System."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/stop/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapVirtualInstances"),
		// 					Operation: to.Ptr("SAPVirtualInstances_Stop"),
		// 					Description: to.Ptr("Stops the SAP System."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getSizingRecommendations/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 					Operation: to.Ptr("SAPSizingRecommendations"),
		// 					Description: to.Ptr("Get SAP sizing recommendations."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getSapSupportedSku/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 					Operation: to.Ptr("SAPSupportedSku"),
		// 					Description: to.Ptr("Get SAP supported SKUs."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getDiskConfigurations/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 					Operation: to.Ptr("SAPDiskConfigurations"),
		// 					Description: to.Ptr("Get SAP Disk Configurations."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getAvailabilityZoneDetails/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 					Operation: to.Ptr("SAPAvailabilityZoneDetails"),
		// 					Description: to.Ptr("Get SAP Availability Zone Details."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors"),
		// 					Operation: to.Ptr("Connectors_Get"),
		// 					Description: to.Ptr("Gets a connector resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors"),
		// 					Operation: to.Ptr("Connectors_Create"),
		// 					Description: to.Ptr("Creates a connector resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors"),
		// 					Operation: to.Ptr("Connectors_Delete"),
		// 					Description: to.Ptr("Deletes a connector resource and its child resources, which are the associated connection resources. All the child resources have to be deleted before deleting the connector resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors"),
		// 					Operation: to.Ptr("Connectors_Update"),
		// 					Description: to.Ptr("Updates a connector resource"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors"),
		// 					Operation: to.Ptr("Connectors_ListByResourceGroup"),
		// 					Description: to.Ptr("Gets all connector resources in a Resource Group."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors"),
		// 					Operation: to.Ptr("Connectors_ListBySubscription"),
		// 					Description: to.Ptr("Gets all connector resources in a Subscription."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/acssBackups"),
		// 					Operation: to.Ptr("ACSSBackupConnections_Get"),
		// 					Description: to.Ptr("Gets the backup connection resource of virtual instance for SAP."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/acssBackups"),
		// 					Operation: to.Ptr("ACSSBackupConnections_Create"),
		// 					Description: to.Ptr("Creates the backup connection resource of virtual instance for SAP."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/acssBackups"),
		// 					Operation: to.Ptr("ACSSBackupConnections_Delete"),
		// 					Description: to.Ptr("Deletes the backup connection resource of virtual instance for SAP."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/acssBackups"),
		// 					Operation: to.Ptr("ACSSBackupConnections_Update"),
		// 					Description: to.Ptr("Updates the backup connection resource of virtual instance for SAP. <br><br>This can be used to update tags on the resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/acssBackups"),
		// 					Operation: to.Ptr("ACSSBackupConnections_List"),
		// 					Description: to.Ptr("Lists the backup connection resources of virtual instance for SAP under the given connector resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/sapVirtualInstanceMonitors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/sapVirtualInstanceMonitors"),
		// 					Operation: to.Ptr("SapVirtualInstanceMonitorConnections_Get"),
		// 					Description: to.Ptr("Gets the monitor connection resource of virtual instance for SAP."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/sapVirtualInstanceMonitors/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/sapVirtualInstanceMonitors"),
		// 					Operation: to.Ptr("SapVirtualInstanceMonitorConnections_Create"),
		// 					Description: to.Ptr("Creates the monitor connection resource of virtual instance for SAP."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/sapVirtualInstanceMonitors/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/sapVirtualInstanceMonitors"),
		// 					Operation: to.Ptr("SapVirtualInstanceMonitorConnections_Delete"),
		// 					Description: to.Ptr("Deletes the monitor connection resource of virtual instance for SAP."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/sapVirtualInstanceMonitors/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/sapVirtualInstanceMonitors"),
		// 					Operation: to.Ptr("SapVirtualInstanceMonitorConnections_Update"),
		// 					Description: to.Ptr("Updates the monitor connection resource of virtual instance for SAP. <br><br>This can be used to update tags on the resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/connectors/sapVirtualInstanceMonitors/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("connectors/sapVirtualInstanceMonitors"),
		// 					Operation: to.Ptr("SapVirtualInstanceMonitorConnections_List"),
		// 					Description: to.Ptr("Lists the monitor connection resources of virtual instance for SAP under the given connector resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/insights/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("insights"),
		// 					Operation: to.Ptr("Insights_Get"),
		// 					Description: to.Ptr("Gets properties of Workloads Insights instance for the specified subscription, resource group and instance name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/insights/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("insights"),
		// 					Operation: to.Ptr("Insights_ListByResourceGroup"),
		// 					Description: to.Ptr("Gets a list of Workloads Insight instances in the specified subscription and resource group. The operations returns various properties of each instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/insights/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("insights"),
		// 					Operation: to.Ptr("Insights_List"),
		// 					Description: to.Ptr("Gets a list of Workloads Insight instances in the specified subscription. The operations returns various properties of each instance."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/insights/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("insights"),
		// 					Operation: to.Ptr("Insights_Create"),
		// 					Description: to.Ptr("Creates a Workloads Insights instance for the specified subscription, resource group, and instance name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/insights/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("insights"),
		// 					Operation: to.Ptr("Insights_Delete"),
		// 					Description: to.Ptr("Deletes a Workloads Insights instance for the specified subscription, resource group and instance name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/insights/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("insights"),
		// 					Operation: to.Ptr("Insights_Update"),
		// 					Description: to.Ptr("Patches the Workload Insights instance for the specified subscription, resource group, and instance name."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites"),
		// 					Operation: to.Ptr("SAPDiscoverySites_Get"),
		// 					Description: to.Ptr("Gets a SAP Migration discovery site resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites"),
		// 					Operation: to.Ptr("SAPDiscoverySites_Create"),
		// 					Description: to.Ptr("Creates a discovery site for SAP Migration."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites"),
		// 					Operation: to.Ptr("SAPDiscoverySites_Delete"),
		// 					Description: to.Ptr("Deletes a SAP Migration discovery site resource and its child resources, that is the associated SAP Instances and Server Instances."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites"),
		// 					Operation: to.Ptr("SAPDiscoverySites_Update"),
		// 					Description: to.Ptr("SAPDiscoverySites_Update."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites"),
		// 					Operation: to.Ptr("SAPDiscoverySites_ListByResourceGroup"),
		// 					Description: to.Ptr("Gets all SAP Migration discovery site resources in a Resource Group."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites"),
		// 					Operation: to.Ptr("SAPDiscoverySites_ListBySubscription"),
		// 					Description: to.Ptr("Gets all SAP Migration discovery site resources in a Subscription."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/importEntities/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites"),
		// 					Operation: to.Ptr("SAPDiscoverySites_ImportEntities"),
		// 					Description: to.Ptr("Import a SAP Migration discovery site resource and it's child resources, that is the SAP instances and Server instances."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances"),
		// 					Operation: to.Ptr("SAPInstances_Get"),
		// 					Description: to.Ptr("Gets the SAP Instance resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances"),
		// 					Operation: to.Ptr("SAPInstances_Create"),
		// 					Description: to.Ptr("Creates the SAP Instance resource. <br><br>This will be used by service only. PUT operation on this resource by end user will return a Bad Request error."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances"),
		// 					Operation: to.Ptr("SAPInstances_Delete"),
		// 					Description: to.Ptr("Deletes the SAP Instance resource. <br><br>This will be used by service only. Delete operation on this resource by end user will return a Bad Request error. You can delete the parent resource, which is the SAP Migration discovery site resource, using the delete operation on it."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances"),
		// 					Operation: to.Ptr("SAPInstances_Update"),
		// 					Description: to.Ptr("Updates the SAP Instance resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances"),
		// 					Operation: to.Ptr("SAPInstances_List"),
		// 					Description: to.Ptr("Lists the SAP Instance resources for the given SAP Migration discovery site resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/serverInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances/serverInstances"),
		// 					Operation: to.Ptr("ServerInstances_Get"),
		// 					Description: to.Ptr("Gets the Server Instance resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/serverInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances/serverInstances"),
		// 					Operation: to.Ptr("ServerInstances_Create"),
		// 					Description: to.Ptr("Creates the Server Instance resource. <br><br>This will be used by service only. PUT operation on this resource by end user will return a Bad Request error."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/serverInstances/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances/serverInstances"),
		// 					Operation: to.Ptr("ServerInstances_Delete"),
		// 					Description: to.Ptr("Deletes the Server Instance resource. <br><br>This will be used by service only. Delete operation on this resource by end user will return a Bad Request error. You can delete the parent resource, which is the SAP Migration discovery site resource, using the delete operation on it."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/serverInstances/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances/serverInstances"),
		// 					Operation: to.Ptr("ServerInstances_Update"),
		// 					Description: to.Ptr("Updates the Server Instance resource. This operation on a resource by end user will return a Bad Request error."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Workloads/sapDiscoverySites/sapInstances/serverInstances/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armworkloadssapvirtualinstance.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft.Workloads"),
		// 					Resource: to.Ptr("sapDiscoverySites/sapInstances/serverInstances"),
		// 					Operation: to.Ptr("ServerInstances_List"),
		// 					Description: to.Ptr("Lists the Server Instance resources for the given SAP Instance resource."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
