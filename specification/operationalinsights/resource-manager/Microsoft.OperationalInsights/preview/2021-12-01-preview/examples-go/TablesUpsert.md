Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.3.0/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/TablesUpsert.json
func ExampleTablesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armoperationalinsights.NewTablesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<table-name>",
		armoperationalinsights.Table{
			Properties: &armoperationalinsights.TableProperties{
				Schema: &armoperationalinsights.Schema{
					Name: to.StringPtr("<name>"),
					Columns: []*armoperationalinsights.Column{
						{
							Name: to.StringPtr("<name>"),
							Type: armoperationalinsights.ColumnTypeEnum("guid").ToPtr(),
						}},
				},
				RetentionInDays:      to.Int32Ptr(45),
				TotalRetentionInDays: to.Int32Ptr(70),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TablesClientUpdateResult)
}
```
