package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armworkloads.OperationListResult{
		// 	Value: []*armworkloads.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/RegisteredSubscriptions/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Reads registered subscriptions"),
		// 				Operation: to.Ptr("Gets/Lists registered subscriptions"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("RegisteredSubscriptions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Lists WordpressInstances resources under a phpWorkload resource"),
		// 				Operation: to.Ptr("WordpressInstances_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a WordpressInstances resource"),
		// 				Operation: to.Ptr("WordpressInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Create or updated WordpressInstances resource"),
		// 				Operation: to.Ptr("WordpressInstances_CreateOrUpdate"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/wordpressInstances/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Delete WordpressInstances resource"),
		// 				Operation: to.Ptr("WordpressInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads/wordpressInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/skus/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets the list of Microsoft.Workloads SKUs available for your Subscription"),
		// 				Operation: to.Ptr("Gets the list of Microsoft.Workloads SKUs available for your Subscription"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("Skus"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/Operations/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("read Operations"),
		// 				Operation: to.Ptr("read_Operations"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/register/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Register the subscription for Microsoft.Workloads"),
		// 				Operation: to.Ptr("Register the Microsoft.Workloads"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("Microsoft.Workloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/unregister/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Unregister the subscription for Microsoft.Workloads"),
		// 				Operation: to.Ptr("Unregister the Microsoft.Workloads"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("Microsoft.Workloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/Locations/OperationStatuses/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("read OperationStatuses"),
		// 				Operation: to.Ptr("read_OperationStatuses"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("Locations/OperationStatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/Locations/OperationStatuses/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("write OperationStatuses"),
		// 				Operation: to.Ptr("write_OperationStatuses"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("Locations/OperationStatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Lists phpWorkload resources in a subscription"),
		// 				Operation: to.Ptr("PhpWorkloads_ListBySubscription"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Lists phpWorkload resources in a resource group"),
		// 				Operation: to.Ptr("PhpWorkloads_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a phpWorkload resource"),
		// 				Operation: to.Ptr("PhpWorkloads_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Create or updated phpWorkloads resource"),
		// 				Operation: to.Ptr("PhpWorkloads_CreateOrUpdate"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Delete phpWorkloads resource"),
		// 				Operation: to.Ptr("PhpWorkloads_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/phpWorkloads/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Update PHP workload resource."),
		// 				Operation: to.Ptr("PhpWorkloads_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("phpWorkloads"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets the SAP Application Server Instance."),
		// 				Operation: to.Ptr("SAPApplicationServerInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Puts the SAP Application Server Instance."),
		// 				Operation: to.Ptr("SAPApplicationServerInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes the SAP Application Server Instance. <br><br>This operation will be used by service only. Delete by end user will return a Bad Request error."),
		// 				Operation: to.Ptr("SAPApplicationServerInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Puts the SAP Application Server Instance."),
		// 				Operation: to.Ptr("SAPApplicationServerInstances_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Lists the SAP Application server Instances in an SVI."),
		// 				Operation: to.Ptr("SAPApplicationServerInstances_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/start/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Starts the SAP Application server Instance in an SVI."),
		// 				Operation: to.Ptr("SAPApplicationServerInstances_StartInstance"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/applicationInstances/stop/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Stops the SAP Application server Instance in an SVI."),
		// 				Operation: to.Ptr("SAPApplicationServerInstances_StopInstance"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/applicationInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets the SAP Central Instance."),
		// 				Operation: to.Ptr("SAPCentralInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Puts the SAP Central Instance. <br><br>This will be used by service only. PUT by end user will return a Bad Request error."),
		// 				Operation: to.Ptr("SAPCentralInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes the SAP Central Instance. <br><br>This will be used by service only. Delete by end user will return a Bad Request error."),
		// 				Operation: to.Ptr("SAPCentralInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Updates the SAP Central Instance. <br><br>This can be used to update tags."),
		// 				Operation: to.Ptr("SAPCentralInstances_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Lists the SAP Central Instances in an SVI."),
		// 				Operation: to.Ptr("SAPCentralInstances_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/start/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Starts the SAP Central server Instance in an SVI."),
		// 				Operation: to.Ptr("SAPCentralInstances_StartInstance"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/centralInstances/stop/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Stops the SAP Central server Instance in an SVI."),
		// 				Operation: to.Ptr("SAPCentralInstances_StopInstance"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/centralInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets the SAP Database Instance."),
		// 				Operation: to.Ptr("SAPDatabaseInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Puts the SAP Database Instance."),
		// 				Operation: to.Ptr("SAPDatabaseInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes the SAP Database Instance. <br><br>This will be used by service only. Delete by end user will return a Bad Request error."),
		// 				Operation: to.Ptr("SAPDatabaseInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Puts the SAP Database Instance."),
		// 				Operation: to.Ptr("SAPDatabaseInstances_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Lists the SAP Database Instances in an SVI."),
		// 				Operation: to.Ptr("SAPDatabaseInstances_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/start/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Starts the database instance of the SAP system."),
		// 				Operation: to.Ptr("SAPDatabaseInstances_StartInstance"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/databaseInstances/stop/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Stops the database instance of the SAP system."),
		// 				Operation: to.Ptr("SAPDatabaseInstances_StopInstance"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances/databaseInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of SAP monitors in the specified subscription. The operations returns various properties of each SAP monitor."),
		// 				Operation: to.Ptr("monitors_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of SAP monitors in the specified resource group."),
		// 				Operation: to.Ptr("monitors_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets properties of a SAP monitor for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("monitors_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Creates a SAP monitor for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("monitors_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes a SAP monitor with the specified subscription, resource group, and monitor name."),
		// 				Operation: to.Ptr("monitors_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Patches the Tags field of a SAP monitor for the specified subscription, resource group, and monitor name."),
		// 				Operation: to.Ptr("monitors_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of provider instances in the specified SAP monitor. The operations returns various properties of each provider instances."),
		// 				Operation: to.Ptr("ProviderInstances_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets properties of a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				Operation: to.Ptr("ProviderInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Creates a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				Operation: to.Ptr("ProviderInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/providerInstances/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes a provider instance for the specified subscription, resource group, Monitor name, and resource name."),
		// 				Operation: to.Ptr("ProviderInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/providerInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of properties of a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets properties of a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Creates a SAP Landscape monitor configuration for the specified subscription, resource group, and resource name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes a SAP Landscape monitor configuration with the specified subscription, resource group, and monitor name."),
		// 				Operation: to.Ptr("SapLandscapeMonitor_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("monitors/sapLandscapeMonitor"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets an SAP Virtual Instance."),
		// 				Operation: to.Ptr("SAPVirtualInstances_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Creates an SAP Virtual Instance."),
		// 				Operation: to.Ptr("SAPVirtualInstances_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes an SAP Virtual Instance."),
		// 				Operation: to.Ptr("SAPVirtualInstances_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Updates an SAP Virtual Instance."),
		// 				Operation: to.Ptr("SAPVirtualInstances_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets all SAP Virtual Instances in a resource group."),
		// 				Operation: to.Ptr("SAPVirtualInstances_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets all SAP Virtual Instances in the subscription."),
		// 				Operation: to.Ptr("SAPVirtualInstances_ListBySubscription"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/start/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Starts the SAP System."),
		// 				Operation: to.Ptr("SAPVirtualInstances_Start"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/sapVirtualInstances/stop/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Stops the SAP System."),
		// 				Operation: to.Ptr("SAPVirtualInstances_Stop"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("sapVirtualInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getSizingRecommendations/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Get SAP sizing recommendations."),
		// 				Operation: to.Ptr("SAPSizingRecommendations"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getSapSupportedSku/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Get SAP supported SKUs."),
		// 				Operation: to.Ptr("SAPSupportedSku"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getDiskConfigurations/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Get SAP Disk Configurations."),
		// 				Operation: to.Ptr("SAPDiskConfigurations"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/locations/sapVirtualInstanceMetadata/getAvailabilityZoneDetails/action"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Get SAP Availability Zone Details."),
		// 				Operation: to.Ptr("SAPAvailabilityZoneDetails"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("locations/sapVirtualInstanceMetadata"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a connector resource"),
		// 				Operation: to.Ptr("Connectors_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Creates a connector resource"),
		// 				Operation: to.Ptr("Connectors_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes a connector resource and its child resources, which are the associated connection resources. All the child resources have to be deleted before deleting the connector resource."),
		// 				Operation: to.Ptr("Connectors_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Updates a connector resource"),
		// 				Operation: to.Ptr("Connectors_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets all connector resources in a Resource Group."),
		// 				Operation: to.Ptr("Connectors_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets all connector resources in a Subscription."),
		// 				Operation: to.Ptr("Connectors_ListBySubscription"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets the backup connection resource of virtual instance for SAP."),
		// 				Operation: to.Ptr("ACSSBackupConnections_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors/acssBackups"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Creates the backup connection resource of virtual instance for SAP."),
		// 				Operation: to.Ptr("ACSSBackupConnections_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors/acssBackups"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes the backup connection resource of virtual instance for SAP."),
		// 				Operation: to.Ptr("ACSSBackupConnections_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors/acssBackups"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Updates the backup connection resource of virtual instance for SAP. <br><br>This can be used to update tags on the resource."),
		// 				Operation: to.Ptr("ACSSBackupConnections_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors/acssBackups"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/connectors/acssBackups/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Lists the backup connection resources of virtual instance for SAP under the given connector resource."),
		// 				Operation: to.Ptr("ACSSBackupConnections_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("connectors/acssBackups"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/insights/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets properties of Workloads Insights instance for the specified subscription, resource group and instance name."),
		// 				Operation: to.Ptr("Insights_Get"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("insights"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/insights/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of Workloads Insight instances in the specified subscription and resource group. The operations returns various properties of each instance."),
		// 				Operation: to.Ptr("Insights_ListByResourceGroup"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("insights"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/insights/read"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of Workloads Insight instances in the specified subscription. The operations returns various properties of each instance."),
		// 				Operation: to.Ptr("Insights_List"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("insights"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/insights/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Creates a Workloads Insights instance for the specified subscription, resource group, and instance name."),
		// 				Operation: to.Ptr("Insights_Create"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("insights"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/insights/delete"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Deletes a Workloads Insights instance for the specified subscription, resource group and instance name."),
		// 				Operation: to.Ptr("Insights_Delete"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("insights"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Workloads/insights/write"),
		// 			Display: &armworkloads.OperationDisplay{
		// 				Description: to.Ptr("Patches the Workload Insights instance for the specified subscription, resource group, and instance name."),
		// 				Operation: to.Ptr("Insights_Update"),
		// 				Provider: to.Ptr("Microsoft.Workloads"),
		// 				Resource: to.Ptr("insights"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
