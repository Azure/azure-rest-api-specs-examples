package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/ListPeeringOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armpeering.OperationListResult{
		// 	Value: []*armpeering.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/register/action"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Registers the subscription for the Peering Resource Provider and enables the creation of peerings"),
		// 				Operation: to.Ptr("Registers the Peering Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("Peering Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerAsns/write"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Write any peerAsns"),
		// 				Operation: to.Ptr("Write PeerAsn"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peerAsns"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerAsns/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any peerAsns"),
		// 				Operation: to.Ptr("Read PeerAsn"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peerAsns"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerAsns/delete"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any peerAsns"),
		// 				Operation: to.Ptr("Delete PeerAsn"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peerAsns"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringLocations/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any peeringLocations"),
		// 				Operation: to.Ptr("Read PeeringLocation"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peeringLocations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/legacyPeerings/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any legacyPeerings"),
		// 				Operation: to.Ptr("Read LegacyPeering"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("legacyPeerings"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/write"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Write any peerings"),
		// 				Operation: to.Ptr("Write Peering"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peerings"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any peerings"),
		// 				Operation: to.Ptr("Read Peering"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peerings"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/delete"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any peerings"),
		// 				Operation: to.Ptr("Delete Peering"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peerings"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServiceLocations/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any peeringServiceLocations"),
		// 				Operation: to.Ptr("Read PeeringServiceLocations"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peeringServiceLocations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServiceProviders/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any peeringServiceProviders"),
		// 				Operation: to.Ptr("Read PeeringServiceProviders"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peeringServiceProviders"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any peeringServices"),
		// 				Operation: to.Ptr("Read PeeringServices"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peeringServices"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/write"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Write any peeringServices"),
		// 				Operation: to.Ptr("Write PeeringService"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peeringServices"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/delete"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any peeringServices"),
		// 				Operation: to.Ptr("Delete PeeringServices"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("peeringServices"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/prefixes/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any prefixes"),
		// 				Operation: to.Ptr("Read PeeringServicePrefixes"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("prefixes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/prefixes/write"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Write any prefixes"),
		// 				Operation: to.Ptr("Write PeeringServicePrefixes"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("prefixes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/prefixes/delete"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any prefixes"),
		// 				Operation: to.Ptr("Delete PeeringServicePrefixes"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("prefixes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/registeredPrefixes/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any registeredPrefixes"),
		// 				Operation: to.Ptr("Read RegisteredPrefixes"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("registeredPrefixes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/registeredPrefixes/write"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Write any registeredPrefixes"),
		// 				Operation: to.Ptr("Write RegisteredPrefixes"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("registeredPrefixes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/registeredPrefixes/delete"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any registeredPrefixes"),
		// 				Operation: to.Ptr("Delete RegisteredPrefixes"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("registeredPrefixes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/registeredAsns/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any registeredAsns"),
		// 				Operation: to.Ptr("Read RegisteredAsns"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("registeredAsns"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/registeredAsns/write"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Write any registeredAsns"),
		// 				Operation: to.Ptr("Write RegisteredAsns"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("registeredAsns"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/registeredAsns/delete"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any registeredAsns"),
		// 				Operation: to.Ptr("Delete RegisteredAsns"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("registeredAsns"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peerings/receivedRoutes/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any receivedRoutes"),
		// 				Operation: to.Ptr("Read PeeringReceivedRoutes"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("receivedRoutes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/connectionMonitorTests/read"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Read any Peering Service Connection Monitor Tests"),
		// 				Operation: to.Ptr("Read Peering Service Connection Monitor Tests"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("connectionMonitorTests"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/connectionMonitorTests/write"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Write any connectionMonitorTests"),
		// 				Operation: to.Ptr("Write Peering Service Connection Monitor Tests"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("connectionMonitorTests"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Peering/peeringServices/connectionMonitorTests/delete"),
		// 			Display: &armpeering.OperationDisplayInfo{
		// 				Description: to.Ptr("Delete any Peering Service Connection Monitor Tests"),
		// 				Operation: to.Ptr("Delete Peering Service Connection Monitor Tests"),
		// 				Provider: to.Ptr("Microsoft.Peering"),
		// 				Resource: to.Ptr("connectionMonitorTests"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
