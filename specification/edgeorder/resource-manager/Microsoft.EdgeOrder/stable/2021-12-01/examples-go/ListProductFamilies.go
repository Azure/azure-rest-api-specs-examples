package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListProductFamilies.json
func ExampleManagementClient_NewListProductFamiliesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armedgeorder.NewManagementClient("YourSubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListProductFamiliesPager(armedgeorder.ProductFamiliesRequest{
		FilterableProperties: map[string][]*armedgeorder.FilterableProperty{
			"azurestackedge": {
				{
					Type: to.Ptr(armedgeorder.SupportedFilterTypesShipToCountries),
					SupportedValues: []*string{
						to.Ptr("US")},
				}},
		},
	},
		&armedgeorder.ManagementClientListProductFamiliesOptions{Expand: to.Ptr("configurations"),
			SkipToken: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
