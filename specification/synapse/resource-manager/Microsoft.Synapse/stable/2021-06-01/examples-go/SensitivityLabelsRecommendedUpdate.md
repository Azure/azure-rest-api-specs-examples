Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.1.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsRecommendedUpdate.json
func ExampleSQLPoolRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolRecommendedSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.RecommendedSensitivityLabelUpdateList{
			Operations: []*armsynapse.RecommendedSensitivityLabelUpdate{
				{
					Properties: &armsynapse.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsynapse.RecommendedSensitivityLabelUpdateKindEnable.ToPtr(),
						Table:  to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsynapse.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsynapse.RecommendedSensitivityLabelUpdateKindEnable.ToPtr(),
						Table:  to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsynapse.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsynapse.RecommendedSensitivityLabelUpdateKindDisable.ToPtr(),
						Table:  to.StringPtr("<table>"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
