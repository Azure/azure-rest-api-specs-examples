package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ChildAvailabilityStatuses_List.json
func ExampleChildAvailabilityStatusesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChildAvailabilityStatusesClient().NewListPager("subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4", &armresourcehealth.ChildAvailabilityStatusesClientListOptions{Filter: nil,
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
		// 			ID: to.Ptr("/subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4/providers/Microsoft.ResourceHealth/availabilityStatuses/current"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesAvailable),
		// 				DetailedStatus: to.Ptr(""),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T23:37:44.000Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesTransient),
		// 				ReasonType: to.Ptr(""),
		// 				RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 					{
		// 						Action: to.Ptr("If you're having problems, use the Troubleshoot tool to get recommended solutions."),
		// 						ActionURL: to.Ptr("<#TroubleshootV2Blade>"),
		// 						ActionURLText: to.Ptr("Troubleshoot tool"),
		// 					},
		// 					{
		// 						Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
		// 						ActionURL: to.Ptr("<#SupportCase>"),
		// 						ActionURLText: to.Ptr("contact support"),
		// 				}},
		// 				ReportedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-05T19:45:45.664Z"); return t}()),
		// 				Summary: to.Ptr("There aren't any known Azure platform problems affecting this virtual machine"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-11-30+23%3a36%3a03Z"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 			ID: to.Ptr("/subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4/providers/Microsoft.ResourceHealth/availabilityStatuses/2018-11-30+23%3a36%3a03Z"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesUnavailable),
		// 				DetailedStatus: to.Ptr(""),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T23:36:03.000Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesTransient),
		// 				ReasonType: to.Ptr("Customer Initiated"),
		// 				RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 					{
		// 						Action: to.Ptr("Check back here for status updates"),
		// 						ActionURL: to.Ptr(""),
		// 						ActionURLText: to.Ptr(""),
		// 					},
		// 					{
		// 						Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
		// 						ActionURL: to.Ptr("<#SupportCase>"),
		// 						ActionURLText: to.Ptr("contact support"),
		// 				}},
		// 				ResolutionETA: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T23:56:03.000Z"); return t}()),
		// 				RootCauseAttributionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T23:36:02.514Z"); return t}()),
		// 				Summary: to.Ptr("This virtual machine is rebooting as requested by an authorized user or process. It will be back online after the reboot completes."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-11-30+22%3a32%3a12Z"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 			ID: to.Ptr("/subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4/providers/Microsoft.ResourceHealth/availabilityStatuses/2018-11-30+22%3a32%3a12Z"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesAvailable),
		// 				DetailedStatus: to.Ptr(""),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T22:32:12.000Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesTransient),
		// 				ReasonType: to.Ptr(""),
		// 				RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 					{
		// 						Action: to.Ptr("If you're having problems, use the Troubleshoot tool to get recommended solutions."),
		// 						ActionURL: to.Ptr("<#TroubleshootV2Blade>"),
		// 						ActionURLText: to.Ptr("Troubleshoot tool"),
		// 					},
		// 					{
		// 						Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
		// 						ActionURL: to.Ptr("<#SupportCase>"),
		// 						ActionURLText: to.Ptr("contact support"),
		// 				}},
		// 				Summary: to.Ptr("There aren't any known Azure platform problems affecting this virtual machine"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-11-30+22%3a30%3a23Z"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 			ID: to.Ptr("/subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4/providers/Microsoft.ResourceHealth/availabilityStatuses/2018-11-30+22%3a30%3a23Z"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesUnavailable),
		// 				DetailedStatus: to.Ptr(""),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T22:30:23.000Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesTransient),
		// 				ReasonType: to.Ptr("Customer Initiated"),
		// 				RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 					{
		// 						Action: to.Ptr("Check back here for status updates"),
		// 						ActionURL: to.Ptr(""),
		// 						ActionURLText: to.Ptr(""),
		// 					},
		// 					{
		// 						Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
		// 						ActionURL: to.Ptr("<#SupportCase>"),
		// 						ActionURLText: to.Ptr("contact support"),
		// 				}},
		// 				ResolutionETA: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T22:50:23.000Z"); return t}()),
		// 				RootCauseAttributionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-30T22:30:22.256Z"); return t}()),
		// 				Summary: to.Ptr("This virtual machine is rebooting as requested by an authorized user or process. It will be back online after the reboot completes."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2018-11-21+00%3a00%3a00Z"),
		// 			Type: to.Ptr("Microsoft.ResourceHealth/AvailabilityStatuses"),
		// 			ID: to.Ptr("/subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4/providers/Microsoft.ResourceHealth/availabilityStatuses/2018-11-21+00%3a00%3a00Z"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armresourcehealth.AvailabilityStatusProperties{
		// 				AvailabilityState: to.Ptr(armresourcehealth.AvailabilityStateValuesAvailable),
		// 				DetailedStatus: to.Ptr(""),
		// 				OccurredTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-21T00:00:00.000Z"); return t}()),
		// 				ReasonChronicity: to.Ptr(armresourcehealth.ReasonChronicityTypesTransient),
		// 				ReasonType: to.Ptr(""),
		// 				RecommendedActions: []*armresourcehealth.RecommendedAction{
		// 					{
		// 						Action: to.Ptr("If you're having problems, use the Troubleshoot tool to get recommended solutions."),
		// 						ActionURL: to.Ptr("<#TroubleshootV2Blade>"),
		// 						ActionURLText: to.Ptr("Troubleshoot tool"),
		// 					},
		// 					{
		// 						Action: to.Ptr("If you are experiencing problems you believe are caused by Azure, contact support"),
		// 						ActionURL: to.Ptr("<#SupportCase>"),
		// 						ActionURLText: to.Ptr("contact support"),
		// 				}},
		// 				Summary: to.Ptr("There aren't any known Azure platform problems affecting this virtual machine"),
		// 			},
		// 	}},
		// }
	}
}
