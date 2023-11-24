package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/ReferenceDataSetsCreate.json
func ExampleReferenceDataSetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReferenceDataSetsClient().CreateOrUpdate(ctx, "rg1", "env1", "rds1", armtimeseriesinsights.ReferenceDataSetCreateOrUpdateParameters{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReferenceDataSetResource = armtimeseriesinsights.ReferenceDataSetResource{
	// 	Name: to.Ptr("rds1"),
	// 	Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments/ReferenceDataSets"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1/referenceDataSets/rds1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armtimeseriesinsights.ReferenceDataSetResourceProperties{
	// 		KeyProperties: []*armtimeseriesinsights.ReferenceDataSetKeyProperty{
	// 			{
	// 				Name: to.Ptr("DeviceId1"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeString),
	// 			},
	// 			{
	// 				Name: to.Ptr("DeviceFloor"),
	// 				Type: to.Ptr(armtimeseriesinsights.ReferenceDataKeyPropertyTypeDouble),
	// 		}},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-18T19:20:33.228Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armtimeseriesinsights.ProvisioningStateSucceeded),
	// 	},
	// }
}
