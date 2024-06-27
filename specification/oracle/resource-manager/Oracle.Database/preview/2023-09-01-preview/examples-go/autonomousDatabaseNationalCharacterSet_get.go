package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/autonomousDatabaseNationalCharacterSet_get.json
func ExampleAutonomousDatabaseNationalCharacterSetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutonomousDatabaseNationalCharacterSetsClient().Get(ctx, "eastus", "NATIONAL", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AutonomousDatabaseNationalCharacterSet = armoracledatabase.AutonomousDatabaseNationalCharacterSet{
	// 	Type: to.Ptr("Oracle.Database/locations/autonomousDatabaseNationalCharacterSets"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/locations/eastus/autonomousDatabaseNationalCharacterSets/NATIONAL"),
	// 	Properties: &armoracledatabase.AutonomousDatabaseNationalCharacterSetProperties{
	// 		CharacterSet: to.Ptr("AL16UTF16"),
	// 	},
	// }
}
