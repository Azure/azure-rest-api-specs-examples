Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.5.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedDatabaseSensitivityLabelsCurrentUpdate.json
func ExampleManagedDatabaseSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewManagedDatabaseSensitivityLabelsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<managed-instance-name>",
		"<database-name>",
		armsql.SensitivityLabelUpdateList{
			Operations: []*armsql.SensitivityLabelUpdate{
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("<schema>"),
						Column: to.Ptr("<column>"),
						Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
						SensitivityLabel: &armsql.SensitivityLabel{
							Properties: &armsql.SensitivityLabelProperties{
								InformationType:   to.Ptr("<information-type>"),
								InformationTypeID: to.Ptr("<information-type-id>"),
								LabelID:           to.Ptr("<label-id>"),
								LabelName:         to.Ptr("<label-name>"),
							},
						},
						Table: to.Ptr("<table>"),
					},
				},
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("<schema>"),
						Column: to.Ptr("<column>"),
						Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
						SensitivityLabel: &armsql.SensitivityLabel{
							Properties: &armsql.SensitivityLabelProperties{
								InformationType:   to.Ptr("<information-type>"),
								InformationTypeID: to.Ptr("<information-type-id>"),
								LabelID:           to.Ptr("<label-id>"),
								LabelName:         to.Ptr("<label-name>"),
							},
						},
						Table: to.Ptr("<table>"),
					},
				},
				{
					Properties: &armsql.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("<schema>"),
						Column: to.Ptr("<column>"),
						Op:     to.Ptr(armsql.SensitivityLabelUpdateKindRemove),
						Table:  to.Ptr("<table>"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
```
