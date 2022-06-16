package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/Diagnostics_ExecuteSiteAnalysis.json
func ExampleDiagnosticsClient_ExecuteSiteAnalysisSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewDiagnosticsClient("<subscription-id>", cred, nil)
	res, err := client.ExecuteSiteAnalysisSlot(ctx,
		"<resource-group-name>",
		"<site-name>",
		"<diagnostic-category>",
		"<analysis-name>",
		"<slot>",
		&armappservice.DiagnosticsExecuteSiteAnalysisSlotOptions{StartTime: nil,
			EndTime:   nil,
			TimeGrain: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DiagnosticAnalysis.ID: %s\n", *res.ID)
}
