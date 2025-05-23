package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/AssociatedFiltersForDeployment_list.json
func ExampleListAssociatedTrafficFiltersClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelastic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewListAssociatedTrafficFiltersClient().List(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TrafficFilterResponse = armelastic.TrafficFilterResponse{
	// 	Rulesets: []*armelastic.TrafficFilter{
	// 		{
	// 			Name: to.Ptr("IPFromApi"),
	// 			Type: to.Ptr(armelastic.TypeIP),
	// 			Description: to.Ptr("created from azure"),
	// 			ID: to.Ptr("31d91b5afb6f4c2eaaf104c97b1991dd"),
	// 			IncludeByDefault: to.Ptr(false),
	// 			Region: to.Ptr("azure-eastus"),
	// 			Rules: []*armelastic.TrafficFilterRule{
	// 				{
	// 					Description: to.Ptr("Allow inbound traffic from IP address 192.168.131.0"),
	// 					ID: to.Ptr("f0297dad72af4a5e964cddf817f35e65"),
	// 					Source: to.Ptr("192.168.131.0"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Allow inbound traffic from IP address block 192.168.132.6/22"),
	// 					ID: to.Ptr("f9c00169f0e54931ae72aabde326b589"),
	// 					Source: to.Ptr("192.168.132.6/22"),
	// 			}},
	// 	}},
	// }
}
