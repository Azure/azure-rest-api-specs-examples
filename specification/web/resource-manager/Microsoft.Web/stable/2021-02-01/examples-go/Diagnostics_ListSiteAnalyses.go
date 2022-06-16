package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/Diagnostics_ListSiteAnalyses.json
func ExampleDiagnosticsClient_ListSiteAnalysesSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewDiagnosticsClient("<subscription-id>", cred, nil)
	pager := client.ListSiteAnalysesSlot("<resource-group-name>",
		"<site-name>",
		"<diagnostic-category>",
		"<slot>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("AnalysisDefinition.ID: %s\n", *v.ID)
		}
	}
}
