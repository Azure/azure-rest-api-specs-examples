package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabaseAddPrincipals.json
func ExampleDatabasesClient_AddPrincipals() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewDatabasesClient("<subscription-id>", cred, nil)
	res, err := client.AddPrincipals(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<database-name>",
		armkusto.DatabasePrincipalListRequest{
			Value: []*armkusto.DatabasePrincipal{
				{
					Name:  to.StringPtr("<name>"),
					Type:  armkusto.DatabasePrincipalType("User").ToPtr(),
					AppID: to.StringPtr("<app-id>"),
					Email: to.StringPtr("<email>"),
					Fqn:   to.StringPtr("<fqn>"),
					Role:  armkusto.DatabasePrincipalRole("Admin").ToPtr(),
				},
				{
					Name:  to.StringPtr("<name>"),
					Type:  armkusto.DatabasePrincipalType("Group").ToPtr(),
					AppID: to.StringPtr("<app-id>"),
					Email: to.StringPtr("<email>"),
					Fqn:   to.StringPtr("<fqn>"),
					Role:  armkusto.DatabasePrincipalRole("Viewer").ToPtr(),
				},
				{
					Name:  to.StringPtr("<name>"),
					Type:  armkusto.DatabasePrincipalType("App").ToPtr(),
					AppID: to.StringPtr("<app-id>"),
					Email: to.StringPtr("<email>"),
					Fqn:   to.StringPtr("<fqn>"),
					Role:  armkusto.DatabasePrincipalRole("Admin").ToPtr(),
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatabasesClientAddPrincipalsResult)
}
