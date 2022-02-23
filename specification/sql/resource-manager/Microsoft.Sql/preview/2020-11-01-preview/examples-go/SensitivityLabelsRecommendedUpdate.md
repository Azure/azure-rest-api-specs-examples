Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsRecommendedUpdate.json
func ExampleRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewRecommendedSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.RecommendedSensitivityLabelUpdateList{
			Operations: []*armsql.RecommendedSensitivityLabelUpdate{
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsql.RecommendedSensitivityLabelUpdateKindEnable.ToPtr(),
						Table:  to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsql.RecommendedSensitivityLabelUpdateKindEnable.ToPtr(),
						Table:  to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsql.RecommendedSensitivityLabelUpdateKindDisable.ToPtr(),
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
