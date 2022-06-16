package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsCreate.json
func ExampleReferenceDataSetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtimeseriesinsights.NewReferenceDataSetsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"env1",
		"rds1",
		armtimeseriesinsights.ReferenceDataSetCreateOrUpdateParameters{
			Location: to.Ptr("West US"),
			Properties: &armtimeseriesinsights.ReferenceDataSetCreationProperties{
				KeyProperties: []*armtimeseriesinsights.ReferenceDataSetKeyProperty{
					{
						Name: to.Ptr("DeviceId1"),
						Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeString),
					},
					{
						Name: to.Ptr("DeviceFloor"),
						Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeDouble),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
