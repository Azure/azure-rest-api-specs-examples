package armneonpostgres_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
)

// Generated from example definition: 2025-03-01/NeonDatabases_Update_MaximumSet_Gen.json
func ExampleNeonDatabasesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNeonDatabasesClient().BeginUpdate(ctx, "rgneon", "test-org", "entity-name", "entity-name", "entity-name", armneonpostgres.NeonDatabase{
		Properties: &armneonpostgres.NeonDatabaseProperties{
			EntityName: to.Ptr("entity-name"),
			Attributes: []*armneonpostgres.Attributes{
				{
					Name:  to.Ptr("trhvzyvaqy"),
					Value: to.Ptr("evpkgsskyavybxwwssm"),
				},
			},
			BranchID:  to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
			OwnerName: to.Ptr("odmbeg"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.NeonDatabasesClientUpdateResponse{
	// 	NeonDatabase: &armneonpostgres.NeonDatabase{
	// 		Properties: &armneonpostgres.NeonDatabaseProperties{
	// 			EntityID: to.Ptr("entity-id"),
	// 			EntityName: to.Ptr("entity-name"),
	// 			CreatedAt: to.Ptr("wgdmylla"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("trhvzyvaqy"),
	// 					Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 				},
	// 			},
	// 			BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 			OwnerName: to.Ptr("odmbeg"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org/projects/entity-name/branches/entity-name/neonDatabases/entity-name"),
	// 		Name: to.Ptr("awzpsbzkfarwnmi"),
	// 		Type: to.Ptr("eltqxh"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("hnyidmqyvvtsddrwkmrqlwtlew"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("szuncyyauzxhpzlbcvjkeamp"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 		},
	// 	},
	// }
}
