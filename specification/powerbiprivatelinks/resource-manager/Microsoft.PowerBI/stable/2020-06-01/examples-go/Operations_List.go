package armpowerbiprivatelinks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiprivatelinks/armpowerbiprivatelinks/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/powerbiprivatelinks/resource-manager/Microsoft.PowerBI/stable/2020-06-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerbiprivatelinks.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armpowerbiprivatelinks.OperationListResult{
		// 	Value: []*armpowerbiprivatelinks.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/read"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Gets all resources for the tenant private link service"),
		// 				Operation: to.Ptr("Gets all resources"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/operationResults/operationId/read"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Gets operation result of Private Link Service Resources for Power BI."),
		// 				Operation: to.Ptr("Get Operation Result"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/read"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Gets resources for the tenant private link service by name"),
		// 				Operation: to.Ptr("Gets resource by name"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/write"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates resource for the tenant private link service"),
		// 				Operation: to.Ptr("Creates or updates private link service"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/delete"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Deletes resource for the tenant private link service"),
		// 				Operation: to.Ptr("Deletes private link service"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateLinkResources/read"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Gets private link resources in a Azure resource"),
		// 				Operation: to.Ptr("Gets private link resources by name"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnectionProxies/privateEndpointName.privateEndpointGuid/read"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Gets private endpoint connection proxy for the tenant"),
		// 				Operation: to.Ptr("Gets private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnectionProxies/privateEndpointName.privateEndpointGuid/write"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates private endpoint connection proxy for the tenant"),
		// 				Operation: to.Ptr("Creates or updates private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnectionProxies/privateEndpointName.privateEndpointGuid/delete"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Deletes private endpoint connection proxy for the tenant"),
		// 				Operation: to.Ptr("Deletes private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnectionProxies/privateEndpointName.privateEndpointGuid/validate/action"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Validates a private endpoint connection proxy before create or update"),
		// 				Operation: to.Ptr("Validates a private endpoint connection proxy"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnections/privateEndpointName.privateEndpointGuid/read"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Gets private endpoint connection for the tenant"),
		// 				Operation: to.Ptr("Gets private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnections/privateEndpointName.privateEndpointGuid/write"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates private endpoint connection for the tenant"),
		// 				Operation: to.Ptr("Creates or updates private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.PowerBI/privateLinkServicesForPowerBI/azureResourceName/privateEndpointConnections/privateEndpointName.privateEndpointGuid/delete"),
		// 			Display: &armpowerbiprivatelinks.OperationDisplay{
		// 				Description: to.Ptr("Deletes private endpoint connection for the tenant"),
		// 				Operation: to.Ptr("Deletes private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft Power BI"),
		// 				Resource: to.Ptr("Private link service resources"),
		// 			},
		// 	}},
		// }
	}
}
