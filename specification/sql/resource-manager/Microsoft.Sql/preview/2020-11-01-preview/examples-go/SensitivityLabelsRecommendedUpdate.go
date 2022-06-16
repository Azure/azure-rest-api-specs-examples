package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SensitivityLabelsRecommendedUpdate.json
func ExampleRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewRecommendedSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx,
		"myRG",
		"myServer",
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
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindEnable),
						Table:  to.Ptr("table2"),
					},
				},
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("column3"),
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindDisable),
						Table:  to.Ptr("table1"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
