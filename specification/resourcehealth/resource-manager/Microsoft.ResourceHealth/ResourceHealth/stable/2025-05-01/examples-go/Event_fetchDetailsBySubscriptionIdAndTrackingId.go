package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: 2025-05-01/Event_fetchDetailsBySubscriptionIdAndTrackingId.json
func ExampleEventClient_FetchDetailsBySubscriptionIDAndTrackingID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventClient().FetchDetailsBySubscriptionIDAndTrackingID(ctx, "eventTrackingId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresourcehealth.EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse{
	// 	Event: armresourcehealth.Event{
	// 		Name: to.Ptr("{eventTrackingId}"),
	// 		Type: to.Ptr("/providers/Microsoft.ResourceHealth/events"),
	// 		ID: to.Ptr("/providers/Microsoft.ResourceHealth/events/{eventTrackingId}"),
	// 		Properties: &armresourcehealth.EventProperties{
	// 			Article: &armresourcehealth.EventPropertiesArticle{
	// 				ArticleContent: to.Ptr("<html>An outage alert is being investigated. More information will be provided as it is known</html>"),
	// 			},
	// 			EnableChatWithUs: to.Ptr(false),
	// 			EnableMicrosoftSupport: to.Ptr(true),
	// 			EventLevel: to.Ptr(armresourcehealth.EventLevelValuesWarning),
	// 			EventSource: to.Ptr(armresourcehealth.EventSourceValuesResourceHealth),
	// 			EventTags: []*string{
	// 				to.Ptr("Action Recommended"),
	// 				to.Ptr("False Positive"),
	// 				to.Ptr("Preliminary PIR"),
	// 				to.Ptr("Final PIR"),
	// 			},
	// 			EventType: to.Ptr(armresourcehealth.EventTypeValuesServiceIssue),
	// 			Faqs: []*armresourcehealth.Faq{
	// 				{
	// 					Answer: to.Ptr("This is an answer"),
	// 					LocaleCode: to.Ptr("en"),
	// 					Question: to.Ptr("This is a question"),
	// 				},
	// 			},
	// 			Header: to.Ptr("Your service might have been impacted by an Azure service issue"),
	// 			HirStage: to.Ptr("resolved"),
	// 			Impact: []*armresourcehealth.Impact{
	// 				{
	// 					ImpactedRegions: []*armresourcehealth.ImpactedServiceRegion{
	// 						{
	// 							ImpactedRegion: to.Ptr("West US"),
	// 							ImpactedSubscriptions: []*string{
	// 								to.Ptr("{subscriptionId}"),
	// 							},
	// 							ImpactedTenants: []*string{
	// 							},
	// 							LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-13T15:43:48.1203530Z"); return t}()),
	// 							Status: to.Ptr(armresourcehealth.EventStatusValuesActive),
	// 							Updates: []*armresourcehealth.Update{
	// 								{
	// 									EventTags: []*string{
	// 										to.Ptr("Final PIR"),
	// 									},
	// 									Summary: to.Ptr("Update 3 - An outage alert is being investigated. More information will be provided as it is known."),
	// 									UpdateDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-13T15:43:48.1203530Z"); return t}()),
	// 								},
	// 								{
	// 									EventTags: []*string{
	// 										to.Ptr("False Positive"),
	// 										to.Ptr("Preliminary PIR"),
	// 									},
	// 									Summary: to.Ptr("Update 2 - An outage alert is being investigated. More information will be provided as it is known."),
	// 									UpdateDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-13T10:32:48.1203530Z"); return t}()),
	// 								},
	// 								{
	// 									EventTags: []*string{
	// 										to.Ptr("Action Recommended"),
	// 									},
	// 									Summary: to.Ptr("Update 1 - An outage alert is being investigated. More information will be provided as it is known."),
	// 									UpdateDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-12T15:00:48.1203530Z"); return t}()),
	// 								},
	// 							},
	// 						},
	// 					},
	// 					ImpactedService: to.Ptr("Virtual Machines"),
	// 					ImpactedServiceGUID: to.Ptr("fd8065f5-ffd0-4756-8788-e6a11bf36257"),
	// 				},
	// 			},
	// 			ImpactMitigationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-14T15:43:48.1203530Z"); return t}()),
	// 			ImpactStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-12T14:45:48.1203530Z"); return t}()),
	// 			IsEventSensitive: to.Ptr(false),
	// 			IsHIR: to.Ptr(false),
	// 			LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-05-13T15:43:48.1203530Z"); return t}()),
	// 			Level: to.Ptr(armresourcehealth.LevelValuesWarning),
	// 			Links: []*armresourcehealth.Link{
	// 				{
	// 					Type: to.Ptr(armresourcehealth.LinkTypeValuesHyperlink),
	// 					BladeName: to.Ptr("RequestRCABlade"),
	// 					DisplayText: &armresourcehealth.LinkDisplayText{
	// 						LocalizedValue: to.Ptr("Request RCA"),
	// 						Value: to.Ptr("Request RCA"),
	// 					},
	// 					ExtensionName: to.Ptr("Microsoft_Azure_Health"),
	// 					Parameters: map[string]any{
	// 						"rcaRequested": "False",
	// 						"trackingId": "{eventTrackingId}",
	// 					},
	// 				},
	// 				{
	// 					Type: to.Ptr(armresourcehealth.LinkTypeValuesButton),
	// 					BladeName: to.Ptr("AzureHealthBrowseBlade"),
	// 					DisplayText: &armresourcehealth.LinkDisplayText{
	// 						LocalizedValue: to.Ptr("Sign up for updates"),
	// 						Value: to.Ptr("Sign up for updates"),
	// 					},
	// 					ExtensionName: to.Ptr("Microsoft_Azure_Health"),
	// 					Parameters: map[string]any{
	// 						"trackingId": "{eventTrackingId}",
	// 					},
	// 				},
	// 			},
	// 			Priority: to.Ptr[int32](2),
	// 			RecommendedActions: &armresourcehealth.EventPropertiesRecommendedActions{
	// 				Actions: []*armresourcehealth.EventPropertiesRecommendedActionsItem{
	// 					{
	// 						ActionText: to.Ptr("action 1"),
	// 						GroupID: to.Ptr[int32](23243),
	// 					},
	// 					{
	// 						ActionText: to.Ptr("action 2"),
	// 						GroupID: to.Ptr[int32](23432),
	// 					},
	// 				},
	// 				LocaleCode: to.Ptr("en"),
	// 				Message: to.Ptr("Recommended actions title"),
	// 			},
	// 			Status: to.Ptr(armresourcehealth.EventStatusValuesActive),
	// 			Summary: to.Ptr("An outage alert is being investigated. More information will be provided as it is known."),
	// 			Title: to.Ptr("ACTIVE: Virtual machines in West US"),
	// 		},
	// 	},
	// }
}
