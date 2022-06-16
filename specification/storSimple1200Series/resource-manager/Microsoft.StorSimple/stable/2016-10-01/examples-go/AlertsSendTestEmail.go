package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/AlertsSendTestEmail.json
func ExampleAlertsClient_SendTestEmail() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewAlertsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.SendTestEmail(ctx,
		"is2-hlmdhdgu1et",
		"ResourceGroupForSDKTest",
		"hManagerForSDKTest4",
		armstorsimple1200series.SendTestAlertEmailRequest{
			EmailList: []*string{
				to.Ptr("testemailid@contoso.com")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
