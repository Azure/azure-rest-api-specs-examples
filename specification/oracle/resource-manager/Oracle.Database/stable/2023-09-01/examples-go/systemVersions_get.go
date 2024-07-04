package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1c63635d66ae38cff18045ab416a6572d3e15f6e/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/systemVersions_get.json
func ExampleSystemVersionsClient_Get_getExadataSystemVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSystemVersionsClient().Get(ctx, "eastus", "22.x", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SystemVersion = armoracledatabase.SystemVersion{
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/systemVersion/22.1.7.0.0.230113"),
	// 	Properties: &armoracledatabase.SystemVersionProperties{
	// 		SystemVersion: to.Ptr("22.1.7.0.0.230113"),
	// 	},
	// }
}
