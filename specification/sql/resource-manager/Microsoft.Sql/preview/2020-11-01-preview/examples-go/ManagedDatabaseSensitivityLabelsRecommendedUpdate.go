package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsRecommendedUpdate.json
func ExampleManagedDatabaseRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagedDatabaseRecommendedSensitivityLabelsClient().Update(ctx, "myRG", "myManagedInstanceName", "myDatabase", armsql.RecommendedSensitivityLabelUpdateList{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
