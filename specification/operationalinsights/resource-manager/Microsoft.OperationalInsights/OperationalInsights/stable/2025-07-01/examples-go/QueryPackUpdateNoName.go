package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v3"
)

// Generated from example definition: 2025-07-01/QueryPackUpdateNoName.json
func ExampleQueryPacksClient_CreateOrUpdateWithoutName_queryPackUpdateNoName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("86dc51d3-92ed-4d7e-947a-775ea79b4919", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewQueryPacksClient().CreateOrUpdateWithoutName(ctx, "my-resource-group", armoperationalinsights.LogAnalyticsQueryPack{
		Location:   to.Ptr("South Central US"),
		Properties: &armoperationalinsights.LogAnalyticsQueryPackProperties{},
		Tags: map[string]*string{
			"Tag1": to.Ptr("Value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
