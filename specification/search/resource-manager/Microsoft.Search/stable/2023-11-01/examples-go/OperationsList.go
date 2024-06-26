package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7e29dd59eef13ef347d09e41a63f2585be77b3ca/specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armsearch.OperationListResult{
		// 	Value: []*armsearch.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/operations/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Lists all of the available operations of the Microsoft.Search provider."),
		// 				Operation: to.Ptr("List all available operations"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/register/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the search resource provider and enables the creation of search services."),
		// 				Operation: to.Ptr("Register the Search Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/write"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the search service."),
		// 				Operation: to.Ptr("Set Search Service"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Reads the search service."),
		// 				Operation: to.Ptr("Get Search Service"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/delete"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Deletes the search service."),
		// 				Operation: to.Ptr("Delete Search Service"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/start/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Starts the search service."),
		// 				Operation: to.Ptr("Start Search Service"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/stop/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Stops the search service."),
		// 				Operation: to.Ptr("Stop Search Service"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/listAdminKeys/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Reads the admin keys."),
		// 				Operation: to.Ptr("Get Admin Key"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/regenerateAdminKey/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Regenerates the admin key."),
		// 				Operation: to.Ptr("Regenerate Admin Key"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/listQueryKeys/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of query API keys for the given Azure Search service."),
		// 				Operation: to.Ptr("Get Query Keys"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("API Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/createQueryKey/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Creates the query key."),
		// 				Operation: to.Ptr("Create Query Key"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Search Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/deleteQueryKey/delete"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Deletes the query key."),
		// 				Operation: to.Ptr("Delete Query Key"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("API Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/checkNameAvailability/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Checks availability of the service name."),
		// 				Operation: to.Ptr("Check Service Name Availability"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Service Name Availability"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/diagnosticSettings/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Gets the diganostic setting read for the resource"),
		// 				Operation: to.Ptr("Get Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Diagnostic Settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/diagnosticSettings/write"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diganostic setting for the resource"),
		// 				Operation: to.Ptr("Set Diagnostic Setting"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Diagnostic Settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/metricDefinitions/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for the search service"),
		// 				Operation: to.Ptr("Read search service metric definitions"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("The metric definitions for the search service"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/logDefinitions/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for the search service"),
		// 				Operation: to.Ptr("Read search service log definitions"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("The log definition for the search service"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Validates a private endpoint connection create call from NRP side"),
		// 				Operation: to.Ptr("Validate Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnectionProxies/write"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Creates a private endpoint connection proxy with the specified parameters or updates the properties or tags for the specified private endpoint connection proxy"),
		// 				Operation: to.Ptr("Create Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnectionProxies/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connection proxies or gets the properties for the specified private endpoint connection proxy"),
		// 				Operation: to.Ptr("Get Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnectionProxies/delete"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connection proxy"),
		// 				Operation: to.Ptr("Delete Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnections/write"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Creates a private endpoint connections with the specified parameters or updates the properties or tags for the specified private endpoint connections"),
		// 				Operation: to.Ptr("Create Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnections/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connections or gets the properties for the specified private endpoint connections"),
		// 				Operation: to.Ptr("Get Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/privateEndpointConnections/delete"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connections"),
		// 				Operation: to.Ptr("Delete Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/sharedPrivateLinkResources/write"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Creates a new shared private link resource with the specified parameters or updates the properties for the specified shared private link resource"),
		// 				Operation: to.Ptr("Create Shared Private Link Resource"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Shared Private Link Resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/sharedPrivateLinkResources/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of shared private link resources or gets the properties for the specified shared private link resource"),
		// 				Operation: to.Ptr("Get Shared Private Link Resource"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Shared Private Link Resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/sharedPrivateLinkResources/delete"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing shared private link resource"),
		// 				Operation: to.Ptr("Delete Shared Private Link Resource"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Shared Private Link Resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Search/searchServices/sharedPrivateLinkResources/operationStatuses/read"),
		// 			Display: &armsearch.OperationDisplay{
		// 				Description: to.Ptr("Get the details of a long running shared private link resource operation"),
		// 				Operation: to.Ptr("Get Operation Status"),
		// 				Provider: to.Ptr("Microsoft Search"),
		// 				Resource: to.Ptr("Shared Private Link Resource"),
		// 			},
		// 	}},
		// }
	}
}
