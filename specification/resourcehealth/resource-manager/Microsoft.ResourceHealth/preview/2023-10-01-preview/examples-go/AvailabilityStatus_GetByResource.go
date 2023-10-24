package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/AvailabilityStatus_GetByResource.json
func ExampleAvailabilityStatusesClient_GetByResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilityStatusesClient().GetByResource(ctx, "resourceUri", &armresourcehealth.AvailabilityStatusesClientGetByResourceOptions{Filter: nil,
		Expand: to.Ptr("recommendedactions"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilityStatus = armresourcehealth.AvailabilityStatus{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
	// 	ID: to.Ptr("/subscriptions/4abcdefgh-ijkl-mnop-qrstuvwxyz/resourceGroups/rhctestenv/providers/Microsoft.ClassicCompute/virtualMachines /rhctestenvV1PI/providers/Microsoft.ResourceHealth/availabilityStatuses/current"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armresourcehealth.AvailabilityStatusProperties{
	// 		AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesUnavailable),
	// 		Category: to.Ptr("Unplanned"),
	// 		Context: to.Ptr("Platform Initiated"),
	// 		DetailedStatus: to.Ptr("Disk problems are preventing us from automatically recovering your virtual machine"),
	// 		OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:12:00Z"); return t}()),
	// 		ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesPersistent),
	// 		ReasonType: to.Ptr("Unplanned"),
	// 		RecommendedActions: []*armresourcehealth.RecommendedAction{
	// 			{
	// 				Action: to.Ptr("To start this virtual machine, open the resource blade and click Start"),
	// 				ActionURL: to.Ptr("<#ResourceBlade>"),
	// 				ActionURLText: to.Ptr("resource blade"),
	// 			},
	// 			{
	// 				Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
	// 				ActionURL: to.Ptr("<#SupportCase>"),
	// 				ActionURLText: to.Ptr("contact support"),
	// 		}},
	// 		ReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-04T14:11:29.7598931Z"); return t}()),
	// 		ResolutionETA: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:37:00Z"); return t}()),
	// 		RootCauseAttributionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:13:00Z"); return t}()),
	// 		Summary: to.Ptr("We're sorry, we couldn't automatically recover your virtual machine"),
	// 		Title: to.Ptr("Unavailable"),
	// 	},
	// }
}
