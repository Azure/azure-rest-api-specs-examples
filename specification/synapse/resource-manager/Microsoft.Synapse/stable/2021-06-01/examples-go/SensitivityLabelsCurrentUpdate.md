Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.2.1/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsCurrentUpdate.json
func ExampleSQLPoolSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolSensitivityLabelsClient("<subscription-id>", cred, nil)
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		armsynapse.SensitivityLabelUpdateList{
			Operations: []*armsynapse.SensitivityLabelUpdate{
				{
					Properties: &armsynapse.SensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsynapse.SensitivityLabelUpdateKindSet.ToPtr(),
						SensitivityLabel: &armsynapse.SensitivityLabel{
							Properties: &armsynapse.SensitivityLabelProperties{
								InformationType:   to.StringPtr("<information-type>"),
								InformationTypeID: to.StringPtr("<information-type-id>"),
								LabelID:           to.StringPtr("<label-id>"),
								LabelName:         to.StringPtr("<label-name>"),
								Rank:              armsynapse.SensitivityLabelRankLow.ToPtr(),
							},
						},
						Table: to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsynapse.SensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsynapse.SensitivityLabelUpdateKindSet.ToPtr(),
						SensitivityLabel: &armsynapse.SensitivityLabel{
							Properties: &armsynapse.SensitivityLabelProperties{
								InformationType:   to.StringPtr("<information-type>"),
								InformationTypeID: to.StringPtr("<information-type-id>"),
								LabelID:           to.StringPtr("<label-id>"),
								LabelName:         to.StringPtr("<label-name>"),
								Rank:              armsynapse.SensitivityLabelRankCritical.ToPtr(),
							},
						},
						Table: to.StringPtr("<table>"),
					},
				},
				{
					Properties: &armsynapse.SensitivityLabelUpdateProperties{
						Schema: to.StringPtr("<schema>"),
						Column: to.StringPtr("<column>"),
						Op:     armsynapse.SensitivityLabelUpdateKindRemove.ToPtr(),
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
