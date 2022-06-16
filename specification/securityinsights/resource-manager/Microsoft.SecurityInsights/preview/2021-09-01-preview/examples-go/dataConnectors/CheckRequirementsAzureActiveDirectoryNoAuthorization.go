package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/CheckRequirementsAzureActiveDirectoryNoAuthorization.json
func ExampleDataConnectorsCheckRequirementsClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewDataConnectorsCheckRequirementsClient("<subscription-id>", cred, nil)
	res, err := client.Post(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		&armsecurityinsight.AADCheckRequirements{
			Kind: armsecurityinsight.DataConnectorKind("AzureActiveDirectory").ToPtr(),
			Properties: &armsecurityinsight.AADCheckRequirementsProperties{
				TenantID: to.StringPtr("<tenant-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataConnectorsCheckRequirementsClientPostResult)
}
