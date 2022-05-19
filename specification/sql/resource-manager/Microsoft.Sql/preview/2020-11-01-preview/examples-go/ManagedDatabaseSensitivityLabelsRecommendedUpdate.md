Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.6.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsRecommendedUpdate.json
func ExampleManagedDatabaseRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseRecommendedSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx,
		"myRG",
		"myManagedInstanceName",
		"myDatabase",
		armsql.RecommendedSensitivityLabelUpdateList{
			Operations: []*armsql.RecommendedSensitivityLabelUpdate{
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("column1"),
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindEnable),
						Table:  to.Ptr("table1"),
					},
				},
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("column2"),
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindDisable),
						Table:  to.Ptr("table2"),
					},
				},
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("Column3"),
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindDisable),
						Table:  to.Ptr("Table1"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
