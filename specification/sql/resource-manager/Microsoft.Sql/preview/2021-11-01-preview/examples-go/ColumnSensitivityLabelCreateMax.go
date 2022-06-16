package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ColumnSensitivityLabelCreateMax.json
func ExampleSensitivityLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsql.NewSensitivityLabelsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		"<schema-name>",
		"<table-name>",
		"<column-name>",
		armsql.SensitivityLabel{
			Properties: &armsql.SensitivityLabelProperties{
				InformationType:   to.Ptr("<information-type>"),
				InformationTypeID: to.Ptr("<information-type-id>"),
				LabelID:           to.Ptr("<label-id>"),
				LabelName:         to.Ptr("<label-name>"),
				Rank:              to.Ptr(armsql.SensitivityLabelRankLow),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
