package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/operations-list-all.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().List(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationCollection = armmediaservices.OperationCollection{
	// 	Value: []*armmediaservices.Operation{
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/register/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Registers the subscription for the Media Services resource provider and enables the creation of Media Services accounts"),
	// 				Operation: to.Ptr("Registers the Media Services Resource Provider"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Microsoft Media Services"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/unregister/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Unregisters the subscription for the Media Services resource provider"),
	// 				Operation: to.Ptr("Unregisters the Media Services Resource Provider"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Microsoft Media Services"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/checknameavailability/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Checks if a Media Services account name is available"),
	// 				Operation: to.Ptr("Check Name Availability"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Microsoft Media Services"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/operations/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Get Available Operations"),
	// 				Operation: to.Ptr("Get Available Operations"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Available Operations"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Media Services Account"),
	// 				Operation: to.Ptr("Read Media Services Account"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Media Services Account"),
	// 				Operation: to.Ptr("Create or Update Media Services Account"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Media Services Account"),
	// 				Operation: to.Ptr("Delete Media Services Account"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/regenerateKey/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Regenerate a Media Services ACS key"),
	// 				Operation: to.Ptr("Regenerate Key"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/listKeys/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("List the ACS keys for the Media Services account"),
	// 				Operation: to.Ptr("List Keys"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/syncStorageKeys/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Synchronize the Storage Keys for an attached Azure Storage account"),
	// 				Operation: to.Ptr("Synchronize Storage Keys"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/listEdgePolicies/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("List policies for an edge device."),
	// 				Operation: to.Ptr("List policies for an edge device."),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/PrivateEndpointConnectionsApproval/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Approve Private Endpoint Connections"),
	// 				Operation: to.Ptr("Approve Private Endpoint Connections"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Media Services Account"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/eventGridFilters/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Event Grid Filter"),
	// 				Operation: to.Ptr("Read Event Grid Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Event Grid Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/eventGridFilters/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Event Grid Filter"),
	// 				Operation: to.Ptr("Create or Update Event Grid Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Event Grid Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/eventGridFilters/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Event Grid Filter"),
	// 				Operation: to.Ptr("Delete Event Grid Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Event Grid Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/accountfilters/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Account Filter"),
	// 				Operation: to.Ptr("Read Account Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Account Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/accountfilters/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Account Filter"),
	// 				Operation: to.Ptr("Create or Update Account Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Account Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/accountfilters/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Account Filter"),
	// 				Operation: to.Ptr("Delete Account Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Account Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Asset"),
	// 				Operation: to.Ptr("Read Asset"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Asset"),
	// 				Operation: to.Ptr("Create or Update Asset"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Asset"),
	// 				Operation: to.Ptr("Delete Asset"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/listContainerSas/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("List Asset Container SAS URLs"),
	// 				Operation: to.Ptr("List Asset Container SAS URLs"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/getEncryptionKey/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Get Asset Encryption Key"),
	// 				Operation: to.Ptr("Get Asset Encryption Key"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/listStreamingLocators/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("List Streaming Locators for Asset"),
	// 				Operation: to.Ptr("List Streaming Locators for Asset"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/assetfilters/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Asset Filter"),
	// 				Operation: to.Ptr("Read Asset Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/assetfilters/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Asset Filter"),
	// 				Operation: to.Ptr("Create or Update Asset Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/assets/assetfilters/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Asset Filter"),
	// 				Operation: to.Ptr("Delete Asset Filter"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Asset Filter"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Streaming Policy"),
	// 				Operation: to.Ptr("Read Streaming Policy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Policy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Streaming Policy"),
	// 				Operation: to.Ptr("Create or Update Streaming Policy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Policy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingPolicies/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Streaming Policy"),
	// 				Operation: to.Ptr("Delete Streaming Policy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Policy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingLocators/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Streaming Locator"),
	// 				Operation: to.Ptr("Read Streaming Locator"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Locator"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingLocators/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Streaming Locator"),
	// 				Operation: to.Ptr("Create or Update Streaming Locator"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Locator"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingLocators/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Streaming Locator"),
	// 				Operation: to.Ptr("Delete Streaming Locator"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Locator"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingLocators/listContentKeys/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("List Content Keys"),
	// 				Operation: to.Ptr("List Content Keys"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Locator"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingLocators/listPaths/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("List Paths"),
	// 				Operation: to.Ptr("List Paths"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Locator"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Content Key Policy"),
	// 				Operation: to.Ptr("Read Content Key Policy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Content Key Policy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Content Key Policy"),
	// 				Operation: to.Ptr("Create or Update Content Key Policy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Content Key Policy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Content Key Policy"),
	// 				Operation: to.Ptr("Delete Content Key Policy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Content Key Policy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/contentKeyPolicies/getPolicyPropertiesWithSecrets/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Get Policy Properties With Secrets"),
	// 				Operation: to.Ptr("Get Policy Properties With Secrets"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Content Key Policy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/transforms/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Transform"),
	// 				Operation: to.Ptr("Read Transform"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Transform"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/transforms/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Transform"),
	// 				Operation: to.Ptr("Create or Update Transform"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Transform"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/transforms/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Transform"),
	// 				Operation: to.Ptr("Delete Transform"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Transform"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/transforms/jobs/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Job"),
	// 				Operation: to.Ptr("Read Job"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Job"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/transforms/jobs/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Job"),
	// 				Operation: to.Ptr("Create or Update Job"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Job"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/transforms/jobs/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Job"),
	// 				Operation: to.Ptr("Delete Job"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Job"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/transforms/jobs/cancelJob/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Cancel Job"),
	// 				Operation: to.Ptr("Cancel Job"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Job"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateLinkResources/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Private Link Resource"),
	// 				Operation: to.Ptr("Read Private Link Resource"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateLinkResource"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnectionProxies/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Private Endpoint Connection Proxy"),
	// 				Operation: to.Ptr("Read Private Endpoint Connection Proxy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnectionProxies/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create Private Endpoint Connection Proxy"),
	// 				Operation: to.Ptr("Create Private Endpoint Connection Proxy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnectionProxies/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete Private Endpoint Connection Proxy"),
	// 				Operation: to.Ptr("Delete Private Endpoint Connection Proxy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnectionProxies/validate/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Validate Private Endpoint Connection Proxy"),
	// 				Operation: to.Ptr("Validate Private Endpoint Connection Proxy"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnections/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Private Endpoint Connection"),
	// 				Operation: to.Ptr("Read Private Endpoint Connection"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateEndpointConnection"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnections/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create Private Endpoint Connection"),
	// 				Operation: to.Ptr("Create Private Endpoint Connection"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateEndpointConnection"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnections/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete Private Endpoint Connection"),
	// 				Operation: to.Ptr("Delete Private Endpoint Connection"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("PrivateEndpointConnection"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnectionOperations/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Private Endpoint Connection Operation"),
	// 				Operation: to.Ptr("Read Private Endpoint Connection Operation"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Private Endpoint Connection Operation"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Read any Streaming Endpoint"),
	// 				Operation: to.Ptr("Read Streaming Endpoint"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoint"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Create or Update any Streaming Endpoint"),
	// 				Operation: to.Ptr("Create or Update Streaming Endpoint"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoint"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/delete"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Delete any Streaming Endpoint"),
	// 				Operation: to.Ptr("Delete Streaming Endpoint"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoint"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/start/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Start any Streaming Endpoint Operation"),
	// 				Operation: to.Ptr("Start Streaming Endpoint Operation"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoint"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/stop/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Stop any Streaming Endpoint Operation"),
	// 				Operation: to.Ptr("Stop Streaming Endpoint Operation"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoint"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/scale/action"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Scale any Streaming Endpoint Operation"),
	// 				Operation: to.Ptr("Scale Streaming Endpoint Operation"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoint"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/providers/Microsoft.Insights/diagnosticSettings/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Gets the diagnostic setting for the resource."),
	// 				Operation: to.Ptr("Read diagnostic setting"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoints"),
	// 			},
	// 			Origin: to.Ptr("system"),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/providers/Microsoft.Insights/diagnosticSettings/write"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
	// 				Operation: to.Ptr("Write diagnostic setting"),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoints"),
	// 			},
	// 			Origin: to.Ptr("system"),
	// 		},
	// 		{
	// 			Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpoints/providers/Microsoft.Insights/metricDefinitions/read"),
	// 			Display: &armmediaservices.OperationDisplay{
	// 				Description: to.Ptr("Get list of Media Services Streaming Endpoint Metrics definitions."),
	// 				Operation: to.Ptr("Get list of Media Services Streaming Endpoint Metrics definitions."),
	// 				Provider: to.Ptr("Microsoft Media Services"),
	// 				Resource: to.Ptr("Streaming Endpoints"),
	// 			},
	// 			Origin: to.Ptr("system"),
	// 			Properties: &armmediaservices.Properties{
	// 				ServiceSpecification: &armmediaservices.ServiceSpecification{
	// 					MetricSpecifications: []*armmediaservices.MetricSpecification{
	// 						{
	// 							Name: to.Ptr("Egress"),
	// 							AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeTotal),
	// 							Dimensions: []*armmediaservices.MetricDimension{
	// 								{
	// 									Name: to.Ptr("OutputFormat"),
	// 									DisplayName: to.Ptr("Output Format"),
	// 									ToBeExportedForShoebox: to.Ptr(true),
	// 							}},
	// 							DisplayDescription: to.Ptr("The amount of Egress data, in bytes."),
	// 							DisplayName: to.Ptr("Egress"),
	// 							SupportedAggregationTypes: []*string{
	// 								to.Ptr("Total")},
	// 								Unit: to.Ptr(armmediaservices.MetricUnitBytes),
	// 							},
	// 							{
	// 								Name: to.Ptr("SuccessE2ELatency"),
	// 								AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 								Dimensions: []*armmediaservices.MetricDimension{
	// 									{
	// 										Name: to.Ptr("OutputFormat"),
	// 										DisplayName: to.Ptr("Output Format"),
	// 										ToBeExportedForShoebox: to.Ptr(true),
	// 								}},
	// 								DisplayDescription: to.Ptr("The average latency for successful requests in milliseconds."),
	// 								DisplayName: to.Ptr("Success end to end Latency"),
	// 								SupportedAggregationTypes: []*string{
	// 									to.Ptr("Average"),
	// 									to.Ptr("Minimum"),
	// 									to.Ptr("Maximum")},
	// 									Unit: to.Ptr(armmediaservices.MetricUnitMilliseconds),
	// 								},
	// 								{
	// 									Name: to.Ptr("Requests"),
	// 									AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeTotal),
	// 									Dimensions: []*armmediaservices.MetricDimension{
	// 										{
	// 											Name: to.Ptr("OutputFormat"),
	// 											DisplayName: to.Ptr("Output Format"),
	// 											ToBeExportedForShoebox: to.Ptr(true),
	// 										},
	// 										{
	// 											Name: to.Ptr("HttpStatusCode"),
	// 											DisplayName: to.Ptr("HTTP Status Code"),
	// 											ToBeExportedForShoebox: to.Ptr(true),
	// 										},
	// 										{
	// 											Name: to.Ptr("ErrorCode"),
	// 											DisplayName: to.Ptr("Error Code"),
	// 											ToBeExportedForShoebox: to.Ptr(false),
	// 									}},
	// 									DisplayDescription: to.Ptr("Requests to a Streaming Endpoint."),
	// 									DisplayName: to.Ptr("Requests"),
	// 									SupportedAggregationTypes: []*string{
	// 										to.Ptr("Total")},
	// 										Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 									},
	// 									{
	// 										Name: to.Ptr("EgressBandwidth"),
	// 										AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 										Dimensions: []*armmediaservices.MetricDimension{
	// 										},
	// 										DisplayDescription: to.Ptr("Egress bandwidth in bits per second."),
	// 										DisplayName: to.Ptr("Egress bandwidth"),
	// 										LockAggregationType: to.Ptr(armmediaservices.MetricAggregationTypeTotal),
	// 										SupportedAggregationTypes: []*string{
	// 											to.Ptr("Average"),
	// 											to.Ptr("Minimum"),
	// 											to.Ptr("Maximum")},
	// 											Unit: to.Ptr(armmediaservices.MetricUnit("BitsPerSecond")),
	// 										},
	// 										{
	// 											Name: to.Ptr("CPU"),
	// 											AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 											Dimensions: []*armmediaservices.MetricDimension{
	// 											},
	// 											DisplayDescription: to.Ptr("CPU usage for premium streaming endpoints. This data is not available for standard streaming endpoints."),
	// 											DisplayName: to.Ptr("CPU usage"),
	// 											SupportedAggregationTypes: []*string{
	// 												to.Ptr("Average"),
	// 												to.Ptr("Minimum"),
	// 												to.Ptr("Maximum")},
	// 												Unit: to.Ptr(armmediaservices.MetricUnit("Percent")),
	// 										}},
	// 									},
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/read"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Read any Live Event"),
	// 									Operation: to.Ptr("Read Live Event"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/write"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Create or Update any Live Event"),
	// 									Operation: to.Ptr("Create or Update Live Event"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/delete"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Delete any Live Event"),
	// 									Operation: to.Ptr("Delete Live Event"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/start/action"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Start any Live Event Operation"),
	// 									Operation: to.Ptr("Start Live Event Operation"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/stop/action"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Stop any Live Event Operation"),
	// 									Operation: to.Ptr("Stop Live Event Operation"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/reset/action"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Reset any Live Event Operation"),
	// 									Operation: to.Ptr("Reset Live Event Operation"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/providers/Microsoft.Insights/diagnosticSettings/read"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Gets the diagnostic setting for the resource."),
	// 									Operation: to.Ptr("Read diagnostic setting"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 								Origin: to.Ptr("system"),
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/providers/Microsoft.Insights/diagnosticSettings/write"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
	// 									Operation: to.Ptr("Write diagnostic setting"),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 								Origin: to.Ptr("system"),
	// 							},
	// 							{
	// 								Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/providers/Microsoft.Insights/metricDefinitions/read"),
	// 								Display: &armmediaservices.OperationDisplay{
	// 									Description: to.Ptr("Get a list of Media Services Live Event Metrics definitions."),
	// 									Operation: to.Ptr("Get a list of Media Services Live Event Metrics definitions."),
	// 									Provider: to.Ptr("Microsoft Media Services"),
	// 									Resource: to.Ptr("Live Event"),
	// 								},
	// 								Origin: to.Ptr("system"),
	// 								Properties: &armmediaservices.Properties{
	// 									ServiceSpecification: &armmediaservices.ServiceSpecification{
	// 										MetricSpecifications: []*armmediaservices.MetricSpecification{
	// 											{
	// 												Name: to.Ptr("IngestBitrate"),
	// 												AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 												Dimensions: []*armmediaservices.MetricDimension{
	// 													{
	// 														Name: to.Ptr("TrackName"),
	// 														DisplayName: to.Ptr("Track name"),
	// 														ToBeExportedForShoebox: to.Ptr(true),
	// 												}},
	// 												DisplayDescription: to.Ptr("The incoming bitrate ingested for a live event, in bits per second."),
	// 												DisplayName: to.Ptr("Live Event ingest bitrate"),
	// 												EnableRegionalMdmAccount: to.Ptr(true),
	// 												SourceMdmNamespace: to.Ptr("MicrosoftMediaLiveEvent"),
	// 												SupportedAggregationTypes: []*string{
	// 													to.Ptr("Average"),
	// 													to.Ptr("Minimum"),
	// 													to.Ptr("Maximum")},
	// 													Unit: to.Ptr(armmediaservices.MetricUnit("BitsPerSecond")),
	// 												},
	// 												{
	// 													Name: to.Ptr("IngestLastTimestamp"),
	// 													AggregationType: to.Ptr(armmediaservices.MetricAggregationType("Maximum")),
	// 													Dimensions: []*armmediaservices.MetricDimension{
	// 														{
	// 															Name: to.Ptr("TrackName"),
	// 															DisplayName: to.Ptr("Track name"),
	// 															ToBeExportedForShoebox: to.Ptr(true),
	// 													}},
	// 													DisplayDescription: to.Ptr("Last timestamp ingested for a live event."),
	// 													DisplayName: to.Ptr("Live Event ingest last timestamp"),
	// 													EnableRegionalMdmAccount: to.Ptr(true),
	// 													SourceMdmNamespace: to.Ptr("MicrosoftMediaLiveEvent"),
	// 													SupportedAggregationTypes: []*string{
	// 														to.Ptr("Maximum")},
	// 														Unit: to.Ptr(armmediaservices.MetricUnitMilliseconds),
	// 													},
	// 													{
	// 														Name: to.Ptr("IngestDriftValue"),
	// 														AggregationType: to.Ptr(armmediaservices.MetricAggregationType("Maximum")),
	// 														Dimensions: []*armmediaservices.MetricDimension{
	// 															{
	// 																Name: to.Ptr("TrackName"),
	// 																DisplayName: to.Ptr("Track name"),
	// 																ToBeExportedForShoebox: to.Ptr(true),
	// 														}},
	// 														DisplayDescription: to.Ptr("Drift between the timestamp of the ingested content and the system clock, measured in seconds per minute. A non zero value indicates that the ingested content is arriving slower than system clock time."),
	// 														DisplayName: to.Ptr("Live Event ingest drift value"),
	// 														EnableRegionalMdmAccount: to.Ptr(true),
	// 														SourceMdmNamespace: to.Ptr("MicrosoftMediaLiveEvent"),
	// 														SupportedAggregationTypes: []*string{
	// 															to.Ptr("Maximum")},
	// 															Unit: to.Ptr(armmediaservices.MetricUnit("Seconds")),
	// 														},
	// 														{
	// 															Name: to.Ptr("LiveOutputLastTimestamp"),
	// 															AggregationType: to.Ptr(armmediaservices.MetricAggregationType("Maximum")),
	// 															Dimensions: []*armmediaservices.MetricDimension{
	// 																{
	// 																	Name: to.Ptr("TrackName"),
	// 																	DisplayName: to.Ptr("Track name"),
	// 																	ToBeExportedForShoebox: to.Ptr(true),
	// 															}},
	// 															DisplayDescription: to.Ptr("Timestamp of the last fragment uploaded to storage for a live event output."),
	// 															DisplayName: to.Ptr("Last output timestamp"),
	// 															EnableRegionalMdmAccount: to.Ptr(true),
	// 															SourceMdmNamespace: to.Ptr("MicrosoftMediaLiveEvent"),
	// 															SupportedAggregationTypes: []*string{
	// 																to.Ptr("Maximum")},
	// 																Unit: to.Ptr(armmediaservices.MetricUnitMilliseconds),
	// 														}},
	// 													},
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/liveOutputs/read"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Read any Live Output"),
	// 													Operation: to.Ptr("Read Live Output"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Live Output"),
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/liveOutputs/write"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Create or Update any Live Output"),
	// 													Operation: to.Ptr("Create or Update Live Output"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Live Output"),
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/liveEvents/liveOutputs/delete"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Delete any Live Output"),
	// 													Operation: to.Ptr("Delete Live Output"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Live Output"),
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/streamingEndpointOperations/read"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Read any Streaming Endpoint Operation"),
	// 													Operation: to.Ptr("Read Streaming Endpoint Operation"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Streaming Endpoint Operation"),
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/liveEventOperations/read"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Read any Live Event Operation"),
	// 													Operation: to.Ptr("Read Live Event Operation"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Live Event Operation"),
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/liveOutputOperations/read"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Read any Live Output Operation"),
	// 													Operation: to.Ptr("Read Live Output Operation"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Live Output Operation"),
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/providers/Microsoft.Insights/diagnosticSettings/read"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Gets the diagnostic setting for the resource."),
	// 													Operation: to.Ptr("Read diagnostic setting"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Media Services Account"),
	// 												},
	// 												Origin: to.Ptr("system"),
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/providers/Microsoft.Insights/diagnosticSettings/write"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Creates or updates the diagnostic setting for the resource."),
	// 													Operation: to.Ptr("Write diagnostic setting"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Media Services Account"),
	// 												},
	// 												Origin: to.Ptr("system"),
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/providers/Microsoft.Insights/logDefinitions/read"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Gets the available logs for a Media Services Account"),
	// 													Operation: to.Ptr("Read mediaservices log definitions"),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("The log definition of mediaservices"),
	// 												},
	// 												Origin: to.Ptr("system"),
	// 												Properties: &armmediaservices.Properties{
	// 													ServiceSpecification: &armmediaservices.ServiceSpecification{
	// 														LogSpecifications: []*armmediaservices.LogSpecification{
	// 															{
	// 																Name: to.Ptr("KeyDeliveryRequests"),
	// 																BlobDuration: to.Ptr("PT1H"),
	// 																DisplayName: to.Ptr("Key Delivery Requests"),
	// 														}},
	// 													},
	// 												},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Media/mediaservices/providers/Microsoft.Insights/metricDefinitions/read"),
	// 												Display: &armmediaservices.OperationDisplay{
	// 													Description: to.Ptr("Get list of Media Services Metric definitions."),
	// 													Operation: to.Ptr("Get list of Media Services Metric definitions."),
	// 													Provider: to.Ptr("Microsoft Media Services"),
	// 													Resource: to.Ptr("Media Service"),
	// 												},
	// 												Origin: to.Ptr("system"),
	// 												Properties: &armmediaservices.Properties{
	// 													ServiceSpecification: &armmediaservices.ServiceSpecification{
	// 														MetricSpecifications: []*armmediaservices.MetricSpecification{
	// 															{
	// 																Name: to.Ptr("AssetQuota"),
	// 																AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																DisplayDescription: to.Ptr("How many assets are allowed for current media service account"),
	// 																DisplayName: to.Ptr("Asset quota"),
	// 																EnableRegionalMdmAccount: to.Ptr(true),
	// 																SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																SupportedAggregationTypes: []*string{
	// 																	to.Ptr("Average")},
	// 																	SupportedTimeGrainTypes: []*string{
	// 																		to.Ptr("PT1H"),
	// 																		to.Ptr("PT6H"),
	// 																		to.Ptr("PT12H"),
	// 																		to.Ptr("P1D")},
	// 																		Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																	},
	// 																	{
	// 																		Name: to.Ptr("AssetCount"),
	// 																		AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																		DisplayDescription: to.Ptr("How many assets are already created in current media service account"),
	// 																		DisplayName: to.Ptr("Asset count"),
	// 																		EnableRegionalMdmAccount: to.Ptr(true),
	// 																		SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																		SupportedAggregationTypes: []*string{
	// 																			to.Ptr("Average")},
	// 																			SupportedTimeGrainTypes: []*string{
	// 																				to.Ptr("PT1H"),
	// 																				to.Ptr("PT6H"),
	// 																				to.Ptr("PT12H"),
	// 																				to.Ptr("P1D")},
	// 																				Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																			},
	// 																			{
	// 																				Name: to.Ptr("AssetQuotaUsedPercentage"),
	// 																				AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																				DisplayDescription: to.Ptr("Asset used percentage in current media service account"),
	// 																				DisplayName: to.Ptr("Asset quota used percentage"),
	// 																				EnableRegionalMdmAccount: to.Ptr(true),
	// 																				SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																				SupportedAggregationTypes: []*string{
	// 																					to.Ptr("Average")},
	// 																					SupportedTimeGrainTypes: []*string{
	// 																						to.Ptr("PT1H"),
	// 																						to.Ptr("PT6H"),
	// 																						to.Ptr("PT12H"),
	// 																						to.Ptr("P1D")},
	// 																						Unit: to.Ptr(armmediaservices.MetricUnit("Percent")),
	// 																					},
	// 																					{
	// 																						Name: to.Ptr("ContentKeyPolicyQuota"),
	// 																						AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																						DisplayDescription: to.Ptr("How many content key polices are allowed for current media service account"),
	// 																						DisplayName: to.Ptr("Content Key Policy quota"),
	// 																						EnableRegionalMdmAccount: to.Ptr(true),
	// 																						SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																						SupportedAggregationTypes: []*string{
	// 																							to.Ptr("Average")},
	// 																							SupportedTimeGrainTypes: []*string{
	// 																								to.Ptr("PT1H"),
	// 																								to.Ptr("PT6H"),
	// 																								to.Ptr("PT12H"),
	// 																								to.Ptr("P1D")},
	// 																								Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																							},
	// 																							{
	// 																								Name: to.Ptr("ContentKeyPolicyCount"),
	// 																								AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																								DisplayDescription: to.Ptr("How many content key policies are already created in current media service account"),
	// 																								DisplayName: to.Ptr("Content Key Policy count"),
	// 																								EnableRegionalMdmAccount: to.Ptr(true),
	// 																								SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																								SupportedAggregationTypes: []*string{
	// 																									to.Ptr("Average")},
	// 																									SupportedTimeGrainTypes: []*string{
	// 																										to.Ptr("PT1H"),
	// 																										to.Ptr("PT6H"),
	// 																										to.Ptr("PT12H"),
	// 																										to.Ptr("P1D")},
	// 																										Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																									},
	// 																									{
	// 																										Name: to.Ptr("ContentKeyPolicyQuotaUsedPercentage"),
	// 																										AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																										DisplayDescription: to.Ptr("Content Key Policy used percentage in current media service account"),
	// 																										DisplayName: to.Ptr("Content Key Policy quota used percentage"),
	// 																										EnableRegionalMdmAccount: to.Ptr(true),
	// 																										SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																										SupportedAggregationTypes: []*string{
	// 																											to.Ptr("Average")},
	// 																											SupportedTimeGrainTypes: []*string{
	// 																												to.Ptr("PT1H"),
	// 																												to.Ptr("PT6H"),
	// 																												to.Ptr("PT12H"),
	// 																												to.Ptr("P1D")},
	// 																												Unit: to.Ptr(armmediaservices.MetricUnit("Percent")),
	// 																											},
	// 																											{
	// 																												Name: to.Ptr("StreamingPolicyQuota"),
	// 																												AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																												DisplayDescription: to.Ptr("How many streaming policies are allowed for current media service account"),
	// 																												DisplayName: to.Ptr("Streaming Policy quota"),
	// 																												EnableRegionalMdmAccount: to.Ptr(true),
	// 																												SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																												SupportedAggregationTypes: []*string{
	// 																													to.Ptr("Average")},
	// 																													SupportedTimeGrainTypes: []*string{
	// 																														to.Ptr("PT1H"),
	// 																														to.Ptr("PT6H"),
	// 																														to.Ptr("PT12H"),
	// 																														to.Ptr("P1D")},
	// 																														Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																													},
	// 																													{
	// 																														Name: to.Ptr("StreamingPolicyCount"),
	// 																														AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																														DisplayDescription: to.Ptr("How many streaming policies are already created in current media service account"),
	// 																														DisplayName: to.Ptr("Streaming Policy count"),
	// 																														EnableRegionalMdmAccount: to.Ptr(true),
	// 																														SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																														SupportedAggregationTypes: []*string{
	// 																															to.Ptr("Average")},
	// 																															SupportedTimeGrainTypes: []*string{
	// 																																to.Ptr("PT1H"),
	// 																																to.Ptr("PT6H"),
	// 																																to.Ptr("PT12H"),
	// 																																to.Ptr("P1D")},
	// 																																Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																															},
	// 																															{
	// 																																Name: to.Ptr("StreamingPolicyQuotaUsedPercentage"),
	// 																																AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																																DisplayDescription: to.Ptr("Streaming Policy used percentage in current media service account"),
	// 																																DisplayName: to.Ptr("Streaming Policy quota used percentage"),
	// 																																EnableRegionalMdmAccount: to.Ptr(true),
	// 																																SourceMdmNamespace: to.Ptr("MediaServiceQuotaAndUsage"),
	// 																																SupportedAggregationTypes: []*string{
	// 																																	to.Ptr("Average")},
	// 																																	SupportedTimeGrainTypes: []*string{
	// 																																		to.Ptr("PT1H"),
	// 																																		to.Ptr("PT6H"),
	// 																																		to.Ptr("PT12H"),
	// 																																		to.Ptr("P1D")},
	// 																																		Unit: to.Ptr(armmediaservices.MetricUnit("Percent")),
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("ChannelsAndLiveEventsCount"),
	// 																																		AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																																		DisplayDescription: to.Ptr("The total number of live events in the current media services account"),
	// 																																		DisplayName: to.Ptr("Live event count"),
	// 																																		EnableRegionalMdmAccount: to.Ptr(true),
	// 																																		SourceMdmNamespace: to.Ptr("ClusterResource_ChannelsAndLiveEvents"),
	// 																																		SupportedAggregationTypes: []*string{
	// 																																			to.Ptr("Average")},
	// 																																			Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																																		},
	// 																																		{
	// 																																			Name: to.Ptr("RunningChannelsAndLiveEventsCount"),
	// 																																			AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																																			DisplayDescription: to.Ptr("The total number of running live events in the current media services account"),
	// 																																			DisplayName: to.Ptr("Running live event count"),
	// 																																			EnableRegionalMdmAccount: to.Ptr(true),
	// 																																			SourceMdmNamespace: to.Ptr("ClusterResource_ChannelsAndLiveEvents"),
	// 																																			SupportedAggregationTypes: []*string{
	// 																																				to.Ptr("Average")},
	// 																																				Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																																			},
	// 																																			{
	// 																																				Name: to.Ptr("MaxChannelsAndLiveEventsCount"),
	// 																																				AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																																				DisplayDescription: to.Ptr("The maximum number of live events allowed in the current media services account"),
	// 																																				DisplayName: to.Ptr("Max live event quota"),
	// 																																				EnableRegionalMdmAccount: to.Ptr(true),
	// 																																				SourceMdmNamespace: to.Ptr("ClusterResource_ChannelsAndLiveEvents"),
	// 																																				SupportedAggregationTypes: []*string{
	// 																																					to.Ptr("Average")},
	// 																																					Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																																				},
	// 																																				{
	// 																																					Name: to.Ptr("MaxRunningChannelsAndLiveEventsCount"),
	// 																																					AggregationType: to.Ptr(armmediaservices.MetricAggregationTypeAverage),
	// 																																					DisplayDescription: to.Ptr("The maximum number of running live events allowed in the current media services account"),
	// 																																					DisplayName: to.Ptr("Max running live event quota"),
	// 																																					EnableRegionalMdmAccount: to.Ptr(true),
	// 																																					SourceMdmNamespace: to.Ptr("ClusterResource_ChannelsAndLiveEvents"),
	// 																																					SupportedAggregationTypes: []*string{
	// 																																						to.Ptr("Average")},
	// 																																						Unit: to.Ptr(armmediaservices.MetricUnitCount),
	// 																																				}},
	// 																																			},
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read a Video Analyzer Account"),
	// 																																			Operation: to.Ptr("Read a Video Analyzer Account"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Analyzer Account"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create or Update a Video Analyzer Account"),
	// 																																			Operation: to.Ptr("Create or Update a Video Analyzer Account"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Analyzer Account"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete a Video Analyzer Account"),
	// 																																			Operation: to.Ptr("Delete a Video Analyzer Account"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Analyzer Account"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/PrivateEndpointConnectionsApproval/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Approve Private Endpoint Connections"),
	// 																																			Operation: to.Ptr("Approve Private Endpoint Connections"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Analyzer Account"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/videos/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Video"),
	// 																																			Operation: to.Ptr("Read Video"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/videos/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create or Update any Video"),
	// 																																			Operation: to.Ptr("Create or Update Video"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/videos/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete any Video"),
	// 																																			Operation: to.Ptr("Delete Video"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/videos/listStreamingToken/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Generates a streaming token which can be used for video playback"),
	// 																																			Operation: to.Ptr("Generates a streaming token which can be used for video playback"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/videos/listContentToken/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Generates a content token which can be used for video playback"),
	// 																																			Operation: to.Ptr("Generates a content token which can be used for video playback"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Video Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/accessPolicies/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Access Policy"),
	// 																																			Operation: to.Ptr("Read Access Policy"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Access Policy Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/accessPolicies/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create or Update any Access Policy"),
	// 																																			Operation: to.Ptr("Create or Update Access Policy"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Access Policy Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/accessPolicies/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete any Access Policy"),
	// 																																			Operation: to.Ptr("Delete Access Policy"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Access Policy Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/edgeModules/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Edge Module"),
	// 																																			Operation: to.Ptr("Read Edge Module"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Edge Module Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/edgeModules/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create or Update any Edge Module"),
	// 																																			Operation: to.Ptr("Create or Update Edge Module"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Edge Module Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/edgeModules/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete any Edge Module"),
	// 																																			Operation: to.Ptr("Delete Edge Module"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Edge Module Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/edgeModules/listProvisioningToken/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Creates a new provisioning token. A provisioning token allows for a single instance of Azure Video analyzer IoT edge module to be initialized and authorized to the cloud account. The provisioning token itself is short lived and it is only used for the initial handshake between IoT edge module and the cloud. After the initial handshake, the IoT edge module will agree on a set of authentication keys which will be auto-rotated as long as the module is able to periodically connect to the cloud. A new provisioning token can be generated for the same IoT edge module in case the module state lost or reset"),
	// 																																			Operation: to.Ptr("Creates a new provisioning token"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Edge Module Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineTopologies/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Pipeline Topology"),
	// 																																			Operation: to.Ptr("Read Pipeline Topology"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Topology Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineTopologies/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create or Update any Pipeline Topology"),
	// 																																			Operation: to.Ptr("Create or Update Pipeline Topology"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Topology Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineTopologies/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete any Pipeline Topology"),
	// 																																			Operation: to.Ptr("Delete Pipeline Topology"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Topology Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/livePipelines/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Live Pipeline"),
	// 																																			Operation: to.Ptr("Read Live Pipeline"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Live Pipeline Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/livePipelines/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create or Update any Live Pipeline"),
	// 																																			Operation: to.Ptr("Create or Update Live Pipeline"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Live Pipeline Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/livePipelines/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete any Live Pipeline"),
	// 																																			Operation: to.Ptr("Delete Live Pipeline"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Live Pipeline Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/livePipelines/activate/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Activate any Live Pipeline"),
	// 																																			Operation: to.Ptr("Activate Live Pipeline"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Live Pipeline Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/livePipelines/deactivate/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Deactivate any Live Pipeline"),
	// 																																			Operation: to.Ptr("Deactivate Live Pipeline"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Live Pipeline Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/livePipelines/operationsStatus/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Live Pipeline operation status"),
	// 																																			Operation: to.Ptr("Read Live Pipeline operation status"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Live Pipeline operation status Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineJobs/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Pipeline Job"),
	// 																																			Operation: to.Ptr("Read Pipeline Job"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Job Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineJobs/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create or Update any Pipeline Job"),
	// 																																			Operation: to.Ptr("Create or Update Pipeline Job"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Job Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineJobs/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete any Pipeline Job"),
	// 																																			Operation: to.Ptr("Delete Pipeline Job"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Job Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineJobs/cancel/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Cancel any Pipeline Job"),
	// 																																			Operation: to.Ptr("Cancel Pipeline Job"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Job Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/pipelineJobs/operationsStatus/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Pipeline Job operation status"),
	// 																																			Operation: to.Ptr("Read Pipeline Job operation status"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Pipeline Job operation status Resource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateLinkResources/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Private Link Resource"),
	// 																																			Operation: to.Ptr("Read Private Link Resource"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateLinkResource"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnectionProxies/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Private Endpoint Connection Proxy"),
	// 																																			Operation: to.Ptr("Read Private Endpoint Connection Proxy"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnectionProxies/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create Private Endpoint Connection Proxy"),
	// 																																			Operation: to.Ptr("Create Private Endpoint Connection Proxy"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnectionProxies/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete Private Endpoint Connection Proxy"),
	// 																																			Operation: to.Ptr("Delete Private Endpoint Connection Proxy"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnectionProxies/validate/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Validate Private Endpoint Connection Proxy"),
	// 																																			Operation: to.Ptr("Validate Private Endpoint Connection Proxy"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateEndpointConnectionProxy"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnections/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Private Endpoint Connection"),
	// 																																			Operation: to.Ptr("Read Private Endpoint Connection"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateEndpointConnection"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnections/write"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Create Private Endpoint Connection"),
	// 																																			Operation: to.Ptr("Create Private Endpoint Connection"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateEndpointConnection"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnections/delete"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Delete Private Endpoint Connection"),
	// 																																			Operation: to.Ptr("Delete Private Endpoint Connection"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("PrivateEndpointConnection"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/videoAnalyzers/privateEndpointConnectionOperations/read"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Read any Private Endpoint Connection Operation"),
	// 																																			Operation: to.Ptr("Read Private Endpoint Connection Operation"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Private Endpoint Connection Operation"),
	// 																																		},
	// 																																	},
	// 																																	{
	// 																																		Name: to.Ptr("Microsoft.Media/locations/checkNameAvailability/action"),
	// 																																		Display: &armmediaservices.OperationDisplay{
	// 																																			Description: to.Ptr("Checks if a Media Services account name is available"),
	// 																																			Operation: to.Ptr("Check Name Availability"),
	// 																																			Provider: to.Ptr("Microsoft Media Services"),
	// 																																			Resource: to.Ptr("Microsoft Media Services"),
	// 																																		},
	// 																																}},
	// 																															}
}
