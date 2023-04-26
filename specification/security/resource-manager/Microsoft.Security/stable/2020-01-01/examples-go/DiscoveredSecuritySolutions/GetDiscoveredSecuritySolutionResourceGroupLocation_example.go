package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/DiscoveredSecuritySolutions/GetDiscoveredSecuritySolutionResourceGroupLocation_example.json
func ExampleDiscoveredSecuritySolutionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiscoveredSecuritySolutionsClient().Get(ctx, "myRg2", "centralus", "paloalto7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiscoveredSecuritySolution = armsecurity.DiscoveredSecuritySolution{
	// 	Location: to.Ptr("eastus2"),
	// 	Name: to.Ptr("paloalto7"),
	// 	Type: to.Ptr("Microsoft.Security/locations/discoveredSecuritySolutions"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg2/providers/Microsoft.Security/locations/centralus/discoveredSecuritySolutions/paloalto7"),
	// 	Properties: &armsecurity.DiscoveredSecuritySolutionProperties{
	// 		Offer: to.Ptr("vmseries1"),
	// 		Publisher: to.Ptr("paloaltonetworks"),
	// 		SecurityFamily: to.Ptr(armsecurity.SecurityFamilyNgfw),
	// 		SKU: to.Ptr("byol"),
	// 	},
	// }
}
