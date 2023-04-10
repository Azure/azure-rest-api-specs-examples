package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/GetAllIncidentEntities.json
func ExampleIncidentsClient_ListEntities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().ListEntities(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IncidentEntitiesResponse = armsecurityinsights.IncidentEntitiesResponse{
	// 	Entities: []armsecurityinsights.EntityClassification{
	// 		&armsecurityinsights.AccountEntity{
	// 			Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindEnumAccount),
	// 			Properties: &armsecurityinsights.AccountEntityProperties{
	// 				FriendlyName: to.Ptr("administrator"),
	// 				AccountName: to.Ptr("administrator"),
	// 				NtDomain: to.Ptr("domain"),
	// 			},
	// 	}},
	// 	MetaData: []*armsecurityinsights.IncidentEntitiesResultsMetadata{
	// 		{
	// 			Count: to.Ptr[int32](1),
	// 			EntityKind: to.Ptr(armsecurityinsights.EntityKindEnumAccount),
	// 	}},
	// }
}
