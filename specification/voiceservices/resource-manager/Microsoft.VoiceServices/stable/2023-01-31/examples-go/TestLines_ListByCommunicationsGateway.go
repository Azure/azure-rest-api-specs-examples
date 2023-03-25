package armvoiceservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/voiceservices/armvoiceservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aa85f59e259c4b12197b57b221067c40fa2fe3f1/specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_ListByCommunicationsGateway.json
func ExampleTestLinesClient_NewListByCommunicationsGatewayPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvoiceservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTestLinesClient().NewListByCommunicationsGatewayPager("testrg", "myname", nil)
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
		// page.TestLineListResult = armvoiceservices.TestLineListResult{
		// 	Value: []*armvoiceservices.TestLine{
		// 		{
		// 			Name: to.Ptr("myline"),
		// 			Type: to.Ptr("Microsoft.Voiceservice/communicationsGateways/testLines"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.VoiceService/communicationsGateway/myname/TestLines/myline"),
		// 			Location: to.Ptr("useast"),
		// 			Properties: &armvoiceservices.TestLineProperties{
		// 				PhoneNumber: to.Ptr("+1-555-1234"),
		// 				Purpose: to.Ptr(armvoiceservices.TestLinePurposeAutomated),
		// 			},
		// 	}},
		// }
	}
}
