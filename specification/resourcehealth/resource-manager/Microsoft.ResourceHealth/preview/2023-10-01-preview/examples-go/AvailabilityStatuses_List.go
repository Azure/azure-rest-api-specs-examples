package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/AvailabilityStatuses_List.json
func ExampleAvailabilityStatusesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilityStatusesClient().NewListPager("resourceUri", &armresourcehealth.AvailabilityStatusesClientListOptions{Filter: nil,
		Expand: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AvailabilityStatusListResult = armresourcehealth.AvailabilityStatusListResult{
		// 	Value: []*armresourcehealth.AvailabilityStatus{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 			ID: to.Ptr("/subscriptions/4abcdefgh-ijkl-mnop-qrstuvwxyz/resourceGroups/rhctestenv/providers/Microsoft.ClassicCompute/virtualMachines/rhctes tenvV1PI/providers/Microsoft.ResourceHealth/availabilityStatuses/current"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesUnavailable),
		// 				Category: to.Ptr("Unplanned"),
		// 				Context: to.Ptr("Platform Initiated"),
		// 				DetailedStatus: to.Ptr("Disk problems are preventing us from automatically recovering your virtual machine"),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:12:00.000Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesPersistent),
		// 				ReasonType: to.Ptr("Unplanned"),
		// 				ReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-04T14:11:29.759Z"); return t}()),
		// 				ResolutionETA: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:37:00.000Z"); return t}()),
		// 				RootCauseAttributionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:13:00.000Z"); return t}()),
		// 				ServiceImpactingEvents: []*armresourcehealth.ServiceImpactingEvent{
		// 					{
		// 						CorrelationID: to.Ptr("b56d0180-2d6c-4f7b-b750-c1eca681874c"),
		// 						EventStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-02T19:23:13.711Z"); return t}()),
		// 						EventStatusLastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-02T19:27:04.954Z"); return t}()),
		// 						IncidentProperties: &armresourcehealth.ServiceImpactingEventIncidentProperties{
		// 							IncidentType: to.Ptr("outage"),
		// 							Region: to.Ptr("East US"),
		// 							Service: to.Ptr("Virtual Machines"),
		// 							Title: to.Ptr("Virtual Machines - West Europe [West Europe]"),
		// 						},
		// 						Status: &armresourcehealth.ServiceImpactingEventStatus{
		// 							Value: to.Ptr("Resolved"),
		// 						},
		// 				}},
		// 				Summary: to.Ptr("We're sorry, we couldn't automatically recover your virtual machine"),
		// 				Title: to.Ptr("Unavailable"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2016-03-28+16%3a23%3a00Z"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 			ID: to.Ptr("/subscriptions/4abcdefgh-ijkl-mnop-qrstuvwxyz/resourceGroups/rhctestenv/providers/Microsoft.ClassicCompute/virtualMachines/rhctes tenvV1PI/providers/Microsoft.ResourceHealth/availabilityStatuses/2016-03-28+16%3a23%3a00Z"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesAvailable),
		// 				DetailedStatus: to.Ptr("There aren’t any known Azure platform problems affecting this virtual machine"),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-28T16:23:00.000Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesPersistent),
		// 				Summary: to.Ptr("This virtual machine is running normally"),
		// 			},
		// 	}},
		// }
	}
}
