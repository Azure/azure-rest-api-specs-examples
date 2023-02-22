package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsCurrentUpdate.json
func ExampleSQLPoolSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx, "myRG", "myWorkspace", "mySqlPool", armsynapse.SensitivityLabelUpdateList{
		Operations: []*armsynapse.SensitivityLabelUpdate{
			{
				Properties: &armsynapse.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column1"),
					Op:     to.Ptr(armsynapse.SensitivityLabelUpdateKindSet),
					SensitivityLabel: &armsynapse.SensitivityLabel{
						Properties: &armsynapse.SensitivityLabelProperties{
							InformationType:   to.Ptr("Financial"),
							InformationTypeID: to.Ptr("1D3652D6-422C-4115-82F1-65DAEBC665C8"),
							LabelID:           to.Ptr("3A477B16-9423-432B-AA97-6069B481CEC3"),
							LabelName:         to.Ptr("Highly Confidential"),
							Rank:              to.Ptr(armsynapse.SensitivityLabelRankLow),
						},
					},
					Table: to.Ptr("table1"),
				},
			},
			{
				Properties: &armsynapse.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("column2"),
					Op:     to.Ptr(armsynapse.SensitivityLabelUpdateKindSet),
					SensitivityLabel: &armsynapse.SensitivityLabel{
						Properties: &armsynapse.SensitivityLabelProperties{
							InformationType:   to.Ptr("PhoneNumber"),
							InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
							LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
							LabelName:         to.Ptr("PII"),
							Rank:              to.Ptr(armsynapse.SensitivityLabelRankCritical),
						},
					},
					Table: to.Ptr("table2"),
				},
			},
			{
				Properties: &armsynapse.SensitivityLabelUpdateProperties{
					Schema: to.Ptr("dbo"),
					Column: to.Ptr("Column3"),
					Op:     to.Ptr(armsynapse.SensitivityLabelUpdateKindRemove),
					Table:  to.Ptr("Table1"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
