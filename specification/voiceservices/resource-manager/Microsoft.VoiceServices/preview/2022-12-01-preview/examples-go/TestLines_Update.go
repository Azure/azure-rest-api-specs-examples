package armvoiceservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/voiceservices/armvoiceservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a7b696c2c73218fbca91c7c3bb625ee0f0bbea0/specification/voiceservices/resource-manager/Microsoft.VoiceServices/preview/2022-12-01-preview/examples/TestLines_Update.json
func ExampleTestLinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvoiceservices.NewTestLinesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "testrg", "myname", "myline", armvoiceservices.TestLineUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TestLine = armvoiceservices.TestLine{
	// 	Name: to.Ptr("myline"),
	// 	Type: to.Ptr("Microsoft.Voiceservice/communicationsGateways/testLines"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.VoiceService/communicationsGateway/myname/TestLines/myline"),
	// 	Location: to.Ptr("useast"),
	// 	Properties: &armvoiceservices.TestLineProperties{
	// 		PhoneNumber: to.Ptr("+1-555-1234"),
	// 		Purpose: to.Ptr(armvoiceservices.TestLinePurposeAutomated),
	// 	},
	// }
}
