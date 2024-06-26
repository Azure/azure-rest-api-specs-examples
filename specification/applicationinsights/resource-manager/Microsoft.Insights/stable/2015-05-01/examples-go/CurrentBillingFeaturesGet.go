package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/CurrentBillingFeaturesGet.json
func ExampleComponentCurrentBillingFeaturesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentCurrentBillingFeaturesClient().Get(ctx, "my-resource-group", "my-component", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentBillingFeatures = armapplicationinsights.ComponentBillingFeatures{
	// 	CurrentBillingFeatures: []*string{
	// 		to.Ptr("Basic")},
	// 		DataVolumeCap: &armapplicationinsights.ComponentDataVolumeCap{
	// 			Cap: to.Ptr[float32](500),
	// 			MaxHistoryCap: to.Ptr[float32](500),
	// 			ResetTime: to.Ptr[int32](16),
	// 			StopSendNotificationWhenHitCap: to.Ptr(false),
	// 			StopSendNotificationWhenHitThreshold: to.Ptr(false),
	// 			WarningThreshold: to.Ptr[int32](90),
	// 		},
	// 	}
}
