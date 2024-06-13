package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/incidents/CreateTeam.json
func ExampleIncidentsClient_CreateTeam() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().CreateTeam(ctx, "ambawolvese5resourcegroup", "AmbaE5WestCentralUS", "69a30280-6a4c-4aa7-9af0-5d63f335d600", armsecurityinsights.TeamProperties{
		TeamDescription: to.Ptr("Team description"),
		TeamName:        to.Ptr("Team name"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TeamInformation = armsecurityinsights.TeamInformation{
	// 	Name: to.Ptr("Team name"),
	// 	Description: to.Ptr("Team description"),
	// 	PrimaryChannelURL: to.Ptr("https://teams.microsoft.com/l/team/19:80bf3b25485b4067b7d2dc4eec9e1578%40thread.tacv2/conversations?groupId=99978838-9bda-4ad4-8f93-4cf7ebc50ca5&tenantId=5b5a146c-eba8-46af-96f8-e31b50d15a3f"),
	// 	TeamCreationTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-15T15:08:21.995Z"); return t}()),
	// 	TeamID: to.Ptr("99978838-9bda-4ad4-8f93-4cf7ebc50ca5"),
	// }
}
