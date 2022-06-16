package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_LogTrackingEvents.json
func ExampleIntegrationAccountsClient_LogTrackingEvents() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationAccountsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.LogTrackingEvents(ctx,
		"testResourceGroup",
		"testIntegrationAccount",
		armlogic.TrackingEventsDefinition{
			Events: []*armlogic.TrackingEvent{
				{
					Error: &armlogic.TrackingEventErrorInfo{
						Code:    to.Ptr("NotFound"),
						Message: to.Ptr("Some error occurred"),
					},
					EventLevel: to.Ptr(armlogic.EventLevelInformational),
					EventTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T01:54:49.505567Z"); return t }()),
					Record: map[string]interface{}{
						"agreementProperties": map[string]interface{}{
							"agreementName":       "testAgreement",
							"as2From":             "testas2from",
							"as2To":               "testas2to",
							"receiverPartnerName": "testPartner2",
							"senderPartnerName":   "testPartner1",
						},
						"messageProperties": map[string]interface{}{
							"IsMessageEncrypted":   false,
							"IsMessageSigned":      false,
							"correlationMessageId": "Unique message identifier",
							"direction":            "Receive",
							"dispositionType":      "received-success",
							"fileName":             "test",
							"isMdnExpected":        true,
							"isMessageCompressed":  false,
							"isMessageFailed":      false,
							"isNrrEnabled":         true,
							"mdnType":              "Async",
							"messageId":            "12345",
						},
					},
					RecordType: to.Ptr(armlogic.TrackingRecordTypeAS2Message),
				}},
			SourceType: to.Ptr("Microsoft.Logic/workflows"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
