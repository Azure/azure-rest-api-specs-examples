package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b74978708bb95475562412d4654c00fbcedd9f89/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/Events_ListBySingleResource.json
func ExampleEventsClient_NewListBySingleResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventsClient().NewListBySingleResourcePager("subscriptions/4abcdefgh-ijkl-mnop-qrstuvwxyz/resourceGroups/rhctestenv/providers/Microsoft.Compute/virtualMachines/rhctestenvV1PI", &armresourcehealth.EventsClientListBySingleResourceOptions{Filter: nil})
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
		// page.Events = armresourcehealth.Events{
		// 	Value: []*armresourcehealth.Event{
		// 		{
		// 			Name: to.Ptr("BC_1-FXZ"),
		// 			Type: to.Ptr("/providers/Microsoft.ResourceHealth/events"),
		// 			ID: to.Ptr("/providers/Microsoft.ResourceHealth/events/BC_1-FXZ"),
		// 			Properties: &armresourcehealth.EventProperties{
		// 				Article: &armresourcehealth.EventPropertiesArticle{
		// 					ArticleContent: to.Ptr("<html>An outage alert is being investigated. More information will be provided as it is known</html>"),
		// 				},
		// 				EnableChatWithUs: to.Ptr(false),
		// 				EnableMicrosoftSupport: to.Ptr(true),
		// 				EventLevel: to.Ptr(armresourcehealth.EventLevelValuesInformational),
		// 				EventSource: to.Ptr(armresourcehealth.EventSourceValuesResourceHealth),
		// 				EventType: to.Ptr(armresourcehealth.EventTypeValuesServiceIssue),
		// 				Faqs: []*armresourcehealth.Faq{
		// 					{
		// 						Answer: to.Ptr("This is an answer"),
		// 						LocaleCode: to.Ptr("en"),
		// 						Question: to.Ptr("This is a question"),
		// 				}},
		// 				Header: to.Ptr("Your service might have been impacted by an Azure service issue"),
		// 				HirStage: to.Ptr("active"),
		// 				Impact: []*armresourcehealth.Impact{
		// 					{
		// 						ImpactedRegions: []*armresourcehealth.ImpactedServiceRegion{
		// 							{
		// 								ImpactedRegion: to.Ptr("West US"),
		// 								ImpactedSubscriptions: []*string{
		// 									to.Ptr("{subscriptionId}")},
		// 									LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:05:00Z"); return t}()),
		// 									Status: to.Ptr(armresourcehealth.EventStatusValuesActive),
		// 							}},
		// 							ImpactedService: to.Ptr("Virtual Machines"),
		// 					}},
		// 					ImpactMitigationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
		// 					ImpactStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-07T00:00:00Z"); return t}()),
		// 					IsHIR: to.Ptr(false),
		// 					LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-08T00:00:00Z"); return t}()),
		// 					Level: to.Ptr(armresourcehealth.LevelValuesWarning),
		// 					Links: []*armresourcehealth.Link{
		// 						{
		// 							Type: to.Ptr(armresourcehealth.LinkTypeValuesHyperlink),
		// 							BladeName: to.Ptr("RequestRCABlade"),
		// 							DisplayText: &armresourcehealth.LinkDisplayText{
		// 								LocalizedValue: to.Ptr("Request RCA"),
		// 								Value: to.Ptr("Request RCA"),
		// 							},
		// 							ExtensionName: to.Ptr("Microsoft_Azure_Health"),
		// 							Parameters: map[string]any{
		// 								"rcaRequested": "False",
		// 								"trackingId": "BC_1-FXZ",
		// 							},
		// 						},
		// 						{
		// 							Type: to.Ptr(armresourcehealth.LinkTypeValuesButton),
		// 							BladeName: to.Ptr("AzureHealthBrowseBlade"),
		// 							DisplayText: &armresourcehealth.LinkDisplayText{
		// 								LocalizedValue: to.Ptr("Sign up for updates"),
		// 								Value: to.Ptr("Sign up for updates"),
		// 							},
		// 							ExtensionName: to.Ptr("Microsoft_Azure_Health"),
		// 							Parameters: map[string]any{
		// 								"trackingId": "BC_1-FXZ",
		// 							},
		// 					}},
		// 					Priority: to.Ptr[int32](2),
		// 					RecommendedActions: &armresourcehealth.EventPropertiesRecommendedActions{
		// 						Actions: []*armresourcehealth.EventPropertiesRecommendedActionsItem{
		// 							{
		// 								ActionText: to.Ptr("action 1"),
		// 								GroupID: to.Ptr[int32](23243),
		// 							},
		// 							{
		// 								ActionText: to.Ptr("action 2"),
		// 								GroupID: to.Ptr[int32](23432),
		// 						}},
		// 						LocaleCode: to.Ptr("en"),
		// 						Message: to.Ptr("Recommended actions title"),
		// 					},
		// 					Status: to.Ptr(armresourcehealth.EventStatusValuesActive),
		// 					Summary: to.Ptr("An outage alert is being investigated. More information will be provided as it is known."),
		// 					Title: to.Ptr("ACTIVE: Virtual machines in West US"),
		// 				},
		// 		}},
		// 	}
	}
}
