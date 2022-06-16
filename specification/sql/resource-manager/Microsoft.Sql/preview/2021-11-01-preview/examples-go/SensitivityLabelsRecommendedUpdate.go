package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/SensitivityLabelsRecommendedUpdate.json
func ExampleRecommendedSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewRecommendedSensitivityLabelsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Update(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.RecommendedSensitivityLabelUpdateList{
			Operations: []*armsql.RecommendedSensitivityLabelUpdate{
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.Ptr("<schema>"),
						Column: to.Ptr("<column>"),
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindEnable),
						Table:  to.Ptr("<table>"),
					},
				},
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.Ptr("<schema>"),
						Column: to.Ptr("<column>"),
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindEnable),
						Table:  to.Ptr("<table>"),
					},
				},
				{
					Properties: &armsql.RecommendedSensitivityLabelUpdateProperties{
						Schema: to.Ptr("<schema>"),
						Column: to.Ptr("<column>"),
						Op:     to.Ptr(armsql.RecommendedSensitivityLabelUpdateKindDisable),
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
