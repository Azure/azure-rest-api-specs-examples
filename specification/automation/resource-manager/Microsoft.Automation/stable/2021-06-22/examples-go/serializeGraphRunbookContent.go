package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/serializeGraphRunbookContent.json
func ExampleClient_ConvertGraphRunbookContent_getGraphicalRawRunbookContentFromGraphicalRunbookJsonObject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ConvertGraphRunbookContent(ctx, "rg", "MyAutomationAccount", armautomation.GraphicalRunbookContent{
		GraphRunbookJSON: to.Ptr("<GraphRunbookJSON>"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GraphicalRunbookContent = armautomation.GraphicalRunbookContent{
	// 	GraphRunbookJSON: to.Ptr("<GraphRunbookJSON>"),
	// 	RawContent: &armautomation.RawGraphicalRunbookContent{
	// 		RunbookDefinition: to.Ptr("AAEAAADAQAAAAAAAAAMAgAAAGJPcmNoZXN0cmF0b3IuR3JhcGhSdW5ib29rLk1vZGVsLCBWZXJzaW9uPTcuMy4wLjAsIEN1bHR...."),
	// 		RunbookType: to.Ptr(armautomation.GraphRunbookTypeGraphPowerShell),
	// 		SchemaVersion: to.Ptr("1.10"),
	// 	},
	// }
}
