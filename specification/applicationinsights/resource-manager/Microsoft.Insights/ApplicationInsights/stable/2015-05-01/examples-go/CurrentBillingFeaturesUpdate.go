package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: 2015-05-01/CurrentBillingFeaturesUpdate.json
func ExampleComponentCurrentBillingFeaturesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentCurrentBillingFeaturesClient().Update(ctx, "my-resource-group", "my-component", armapplicationinsights.ComponentBillingFeatures{
		CurrentBillingFeatures: []*string{
			to.Ptr("Basic"),
			to.Ptr("Application Insights Enterprise"),
		},
		DataVolumeCap: &armapplicationinsights.ComponentDataVolumeCap{
			Cap:                            to.Ptr[float32](100),
			StopSendNotificationWhenHitCap: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapplicationinsights.ComponentCurrentBillingFeaturesClientUpdateResponse{
	// 	ComponentBillingFeatures: armapplicationinsights.ComponentBillingFeatures{
	// 		CurrentBillingFeatures: []*string{
	// 			to.Ptr("Basic"),
	// 			to.Ptr("Application Insights Enterprise"),
	// 		},
	// 		DataVolumeCap: &armapplicationinsights.ComponentDataVolumeCap{
	// 			Cap: to.Ptr[float32](100),
	// 			MaxHistoryCap: to.Ptr[float32](500),
	// 			ResetTime: to.Ptr[int32](16),
	// 			StopSendNotificationWhenHitCap: to.Ptr(true),
	// 			StopSendNotificationWhenHitThreshold: to.Ptr(false),
	// 			WarningThreshold: to.Ptr[int32](90),
	// 		},
	// 	},
	// }
}
