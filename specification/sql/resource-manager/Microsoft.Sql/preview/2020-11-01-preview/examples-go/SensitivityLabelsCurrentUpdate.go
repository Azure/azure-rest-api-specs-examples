package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsCurrentUpdate.json
func ExampleSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSensitivityLabelsClient().Update(ctx, "myRG", "myServer", "myDatabase", armsql.SensitivityLabelUpdateList{
		Operations: []*armsql.SensitivityLabelUpdate{
			{
				Properties: &armsql.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column1"),
					Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
					SensitivityLabel: &armsql.SensitivityLabel{
						Properties: &armsql.SensitivityLabelProperties{
							InformationType:   to.Ptr("Financial"),
							InformationTypeID: to.Ptr("1D3652D6-422C-4115-82F1-65DAEBC665C8"),
							LabelID:           to.Ptr("3A477B16-9423-432B-AA97-6069B481CEC3"),
							LabelName:         to.Ptr("Highly Confidential"),
							Rank:              to.Ptr(armsql.SensitivityLabelRankLow),
						},
					},
					Table: to.Ptr("table1"),
				},
			},
			{
				Properties: &armsql.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column2"),
					Op:     to.Ptr(armsql.SensitivityLabelUpdateKindSet),
					SensitivityLabel: &armsql.SensitivityLabel{
						Properties: &armsql.SensitivityLabelProperties{
							InformationType:   to.Ptr("PhoneNumber"),
							InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
							LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
							LabelName:         to.Ptr("PII"),
							Rank:              to.Ptr(armsql.SensitivityLabelRankCritical),
						},
					},
					Table: to.Ptr("table2"),
				},
			},
			{
				Properties: &armsql.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("Column3"),
					Op:     to.Ptr(armsql.SensitivityLabelUpdateKindRemove),
					Table:  to.Ptr("Table1"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
