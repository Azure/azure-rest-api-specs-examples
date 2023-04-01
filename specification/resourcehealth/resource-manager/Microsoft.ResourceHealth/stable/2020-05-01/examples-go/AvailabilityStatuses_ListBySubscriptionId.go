package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/AvailabilityStatuses_ListBySubscriptionId.json
func ExampleAvailabilityStatusesClient_NewListBySubscriptionIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilityStatusesClient().NewListBySubscriptionIDPager(&armresourcehealth.AvailabilityStatusesClientListBySubscriptionIDOptions{Filter: nil,
		Expand: to.Ptr("recommendedactions"),
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
		// 			ID: to.Ptr("<resourceId>/providers/Microsoft.ResourceHealth/AvailabilityStatueses/current"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesAvailable),
		// 				DetailedStatus: to.Ptr("We have not seen any issues with your virtual machine"),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:12:00Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesPersistent),
		// 				ReasonType: to.Ptr("Unplanned"),
		// 				RecentlyResolved: &armresourcehealth.AvailabilityStatusPropertiesRecentlyResolved{
		// 					ResolvedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-28T00:49:00Z"); return t}()),
		// 					UnavailabilitySummary: to.Ptr("We are sorry your SQL database is unavailable"),
		// 					UnavailableOccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-28T00:48:00Z"); return t}()),
		// 				},
		// 				RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 					{
		// 						Action: to.Ptr("To start this virtualmachine, open the resource blade and click Start"),
		// 						ActionURL: to.Ptr("<#ResourceBlade>"),
		// 						ActionURLText: to.Ptr("resourceblade"),
		// 				}},
		// 				ReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-04T14:11:29.7598931Z"); return t}()),
		// 				Summary: to.Ptr("Vm is available"),
		// 				Title: to.Ptr("Available"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 			ID: to.Ptr("<resourceId>/providers/Microsoft.ResourceHealth/AvailabilityStatueses/current"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesUnavailable),
		// 				DetailedStatus: to.Ptr("Diskproblemsarepreventingusfromautomaticallyrecoveringyourvirtualmachine"),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:12:00Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesPersistent),
		// 				ReasonType: to.Ptr("Unplanned"),
		// 				RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 					{
		// 						Action: to.Ptr("To start this virtualmachine, open the resource blade"),
		// 						ActionURL: to.Ptr("<#ResourceBlade>"),
		// 						ActionURLText: to.Ptr("resourceblade"),
		// 					},
		// 					{
		// 						Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
		// 						ActionURL: to.Ptr("<#SupportCase>"),
		// 						ActionURLText: to.Ptr("contactsupport"),
		// 				}},
		// 				ReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-04T14:11:29.7598931Z"); return t}()),
		// 				ResolutionETA: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:37:00Z"); return t}()),
		// 				RootCauseAttributionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:13:00Z"); return t}()),
		// 				Summary: to.Ptr("We are sorry, we couldn't automatically recovery our virtualmachine"),
		// 				Title: to.Ptr("Unavailable"),
		// 			},
		// 	}},
		// }
	}
}
