package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-10-01-preview/enrichment/GetGeodataWithWorkspaceByIp.json
func ExampleClient_ListGeodataByIP() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("bd794837-4d29-4647-9105-6339bfdb4e6a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListGeodataByIP(ctx, "myRg", "myWorkspace", armsecurityinsights.EnrichmentTypeMain, armsecurityinsights.EnrichmentIPAddressBody{
		IPAddress: to.Ptr("1.2.3.4"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.ClientListGeodataByIPResponse{
	// 	EnrichmentIPGeodata: armsecurityinsights.EnrichmentIPGeodata{
	// 		Asn: to.Ptr("12345"),
	// 		Carrier: to.Ptr("Microsoft"),
	// 		City: to.Ptr("Redmond"),
	// 		Continent: to.Ptr("north america"),
	// 		Country: to.Ptr("united states"),
	// 		IPAddr: to.Ptr("1.2.3.4"),
	// 		IPRoutingType: to.Ptr("fixed"),
	// 		Latitude: to.Ptr("40.2436"),
	// 		Longitude: to.Ptr("-100.8891"),
	// 		Organization: to.Ptr("Microsoft"),
	// 		OrganizationType: to.Ptr("tech"),
	// 		Region: to.Ptr("western usa"),
	// 		State: to.Ptr("washington"),
	// 		StateCode: to.Ptr("wa"),
	// 	},
	// }
}
