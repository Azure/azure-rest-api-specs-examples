package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/EnvironmentsPatchTags.json
func ExampleEnvironmentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEnvironmentsClient().BeginUpdate(ctx, "rg1", "env1", &armtimeseriesinsights.EnvironmentUpdateParameters{
		Tags: map[string]*string{
			"someTag": to.Ptr("someTagValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armtimeseriesinsights.EnvironmentsClientUpdateResponse{
	// 	                            EnvironmentResourceClassification: &armtimeseriesinsights.Gen1EnvironmentResource{
	// 		Name: to.Ptr("env1"),
	// 		Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 			"someTag": to.Ptr("someTagValue"),
	// 		},
	// 		Kind: to.Ptr(armtimeseriesinsights.EnvironmentResourceKindGen1),
	// 		SKU: &armtimeseriesinsights.SKU{
	// 			Name: to.Ptr(armtimeseriesinsights.SKUNameS1),
	// 			Capacity: to.Ptr[int32](10),
	// 		},
	// 		Properties: &armtimeseriesinsights.Gen1EnvironmentResourceProperties{
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-18T19:20:33.2288820Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armtimeseriesinsights.ProvisioningStateSucceeded),
	// 			DataRetentionTime: to.Ptr("P31D"),
	// 		},
	// 	},
	// 	                        }
}
