package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/incidents/CreateTeam.json
func ExampleIncidentsClient_CreateTeam() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewIncidentsClient("9023f5b5-df22-4313-8fbf-b4b75af8a6d9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateTeam(ctx,
		"ambawolvese5resourcegroup",
		"AmbaE5WestCentralUS",
		"69a30280-6a4c-4aa7-9af0-5d63f335d600",
		armsecurityinsights.TeamProperties{
			TeamDescription: to.Ptr("Team description"),
			TeamName:        to.Ptr("Team name"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
