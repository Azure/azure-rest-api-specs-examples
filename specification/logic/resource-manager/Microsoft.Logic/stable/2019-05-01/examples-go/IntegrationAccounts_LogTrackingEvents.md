Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.1/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_LogTrackingEvents.json
func ExampleIntegrationAccountsClient_LogTrackingEvents() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewIntegrationAccountsClient("<subscription-id>", cred, nil)
	_, err = client.LogTrackingEvents(ctx,
		"<resource-group-name>",
		"<integration-account-name>",
		armlogic.TrackingEventsDefinition{
			Events: []*armlogic.TrackingEvent{
				{
					Error: &armlogic.TrackingEventErrorInfo{
						Code:    to.StringPtr("<code>"),
						Message: to.StringPtr("<message>"),
					},
					EventLevel: armlogic.EventLevelInformational.ToPtr(),
					EventTime:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-05T01:54:49.505567Z"); return t }()),
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
					RecordType: armlogic.TrackingRecordType("AS2Message").ToPtr(),
				}},
			SourceType: to.StringPtr("<source-type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
