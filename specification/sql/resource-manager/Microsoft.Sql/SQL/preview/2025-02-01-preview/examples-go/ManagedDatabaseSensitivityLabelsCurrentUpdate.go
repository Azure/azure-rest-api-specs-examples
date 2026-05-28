package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/ManagedDatabaseSensitivityLabelsCurrentUpdate.json
func ExampleManagedDatabaseSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseSensitivityLabelsClient().Update(ctx, "myRG", "myManagedInstanceName", "myDatabase", armsql.SensitivityLabelUpdateList{
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
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.ManagedDatabaseSensitivityLabelsClientUpdateResponse{
	// }
}
