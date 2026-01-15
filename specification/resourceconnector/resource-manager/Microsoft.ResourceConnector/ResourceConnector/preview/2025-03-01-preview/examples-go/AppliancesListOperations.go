package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: 2025-03-01-preview/AppliancesListOperations.json
func ExampleAppliancesClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAppliancesClient().NewListOperationsPager(nil)
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
		// page = armresourceconnector.AppliancesClientListOperationsResponse{
		// 	ApplianceOperationsList: armresourceconnector.ApplianceOperationsList{
		// 		Value: []*armresourceconnector.ApplianceOperation{
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/operations/read"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Gets list of Available Operations for Appliances"),
		// 					Operation: to.Ptr("List Available Operations for Appliances"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Operations"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/register/action"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Registers the subscription for Appliance resource provider and enables the creation of Appliance."),
		// 					Operation: to.Ptr("Registers the Appliance Resource Provider"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Appliances Resource Provider"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/appliances/read"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Gets an Appliance resource"),
		// 					Operation: to.Ptr("Get Appliance"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Appliances"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/appliances/write"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Creates or Updates Appliance resource"),
		// 					Operation: to.Ptr("Create or Update Appliance"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Appliances"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/appliances/delete"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Deletes Appliance resource"),
		// 					Operation: to.Ptr("Delete Appliance"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Appliances"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/appliances/listClusterUserCredential"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Get an appliance cluster user credential"),
		// 					Operation: to.Ptr("List User Cluster Credential"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Appliances"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/locations/operationsstatus/read"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Get result of Appliance operation"),
		// 					Operation: to.Ptr("Get status of Appliance operation"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Appliance Operation Status"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.ResourceConnector/locations/operationresults/read"),
		// 				Display: &armresourceconnector.ApplianceOperationValueDisplay{
		// 					Description: to.Ptr("Get result of Appliance operation"),
		// 					Operation: to.Ptr("Get the status of Appliance operation"),
		// 					Provider: to.Ptr("Microsoft ResourceConnector"),
		// 					Resource: to.Ptr("Appliance Operation Result"),
		// 				},
		// 				Origin: to.Ptr("user,system"),
		// 				IsDataAction: to.Ptr(false),
		// 			},
		// 		},
		// 	},
		// }
	}
}
