package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: 2025-05-01/AvailabilityStatuses_ListBySubscriptionId.json
func ExampleAvailabilityStatusesClient_NewListBySubscriptionIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilityStatusesClient().NewListBySubscriptionIDPager(&armresourcehealth.AvailabilityStatusesClientListBySubscriptionIDOptions{
		Expand: to.Ptr("recommendedactions")})
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
		// page = armresourcehealth.AvailabilityStatusesClientListBySubscriptionIDResponse{
		// 	AvailabilityStatusListResult: armresourcehealth.AvailabilityStatusListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionId/providers/Microsoft.ResourceHealth/availabilityStatuses?api-version=2025-05-01&$skipToken=OpaquePageNumber"),
		// 		Value: []*armresourcehealth.AvailabilityStatus{
		// 			{
		// 				Name: to.Ptr("current"),
		// 				Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 				ID: to.Ptr("/subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/resourceGroupName/providers/Microsoft.Compute/virtualMachines/virtualMachineName/providers/Microsoft.ResourceHealth/availabilityStatuses/current"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 					AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesAvailable),
		// 					Category: to.Ptr("Unplanned"),
		// 					Context: to.Ptr("Platform Initiated"),
		// 					DetailedStatus: to.Ptr("We have not seen any issues with your virtual machine"),
		// 					OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:12:00Z"); return t}()),
		// 					ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesPersistent),
		// 					ReasonType: to.Ptr("Unplanned"),
		// 					RecentlyResolved: &armresourcehealth.AvailabilityStatusPropertiesRecentlyResolved{
		// 						ResolvedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-28T00:49:00Z"); return t}()),
		// 						UnavailableOccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-28T00:48:00Z"); return t}()),
		// 						UnavailabilitySummary: to.Ptr("We are sorry your SQL database is unavailable"),
		// 					},
		// 					RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 						{
		// 							Action: to.Ptr("To start this virtualmachine, open the resource blade and click Start"),
		// 							ActionURL: to.Ptr("<#ResourceBlade>"),
		// 							ActionURLText: to.Ptr("resourceblade"),
		// 						},
		// 					},
		// 					ReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-04T14:11:29.7598931Z"); return t}()),
		// 					Summary: to.Ptr("Vm is available"),
		// 					Title: to.Ptr("Available"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("current"),
		// 				Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 				ID: to.Ptr("/subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/resourceGroupName/providers/Microsoft.Compute/virtualMachines/virtualMachineName/providers/Microsoft.ResourceHealth/availabilityStatuses/current"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 					AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesUnavailable),
		// 					DetailedStatus: to.Ptr("Diskproblemsarepreventingusfromautomaticallyrecoveringyourvirtualmachine"),
		// 					OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:12:00Z"); return t}()),
		// 					ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesPersistent),
		// 					ReasonType: to.Ptr("Unplanned"),
		// 					RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 						{
		// 							Action: to.Ptr("To start this virtualmachine, open the resource blade"),
		// 							ActionURL: to.Ptr("<#ResourceBlade>"),
		// 							ActionURLText: to.Ptr("resourceblade"),
		// 						},
		// 						{
		// 							Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
		// 							ActionURL: to.Ptr("<#SupportCase>"),
		// 							ActionURLText: to.Ptr("contactsupport"),
		// 						},
		// 					},
		// 					ReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-04T14:11:29.7598931Z"); return t}()),
		// 					ResolutionETA: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:37:00Z"); return t}()),
		// 					RootCauseAttributionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T09:13:00Z"); return t}()),
		// 					Summary: to.Ptr("We are sorry, we couldn't automatically recovery our virtualmachine"),
		// 					Title: to.Ptr("Unavailable"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
