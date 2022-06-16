package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/incidents/CreateTeam.json
func ExampleIncidentsClient_CreateTeam() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewIncidentsClient("<subscription-id>", cred, nil)
	res, err := client.CreateTeam(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<incident-id>",
		armsecurityinsight.TeamProperties{
			TeamDescription: to.StringPtr("<team-description>"),
			TeamName:        to.StringPtr("<team-name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.IncidentsClientCreateTeamResult)
}
