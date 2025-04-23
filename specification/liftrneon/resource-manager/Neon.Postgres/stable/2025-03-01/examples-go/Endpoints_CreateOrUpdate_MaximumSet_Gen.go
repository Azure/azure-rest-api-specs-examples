package armneonpostgres_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
)

// Generated from example definition: 2025-03-01/Endpoints_CreateOrUpdate_MaximumSet_Gen.json
func ExampleEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEndpointsClient().BeginCreateOrUpdate(ctx, "rgneon", "test-org", "entity-name", "entity-name", "entity-name", armneonpostgres.Endpoint{
		Properties: &armneonpostgres.EndpointProperties{
			EntityName: to.Ptr("entity-name"),
			Attributes: []*armneonpostgres.Attributes{
				{
					Name:  to.Ptr("trhvzyvaqy"),
					Value: to.Ptr("evpkgsskyavybxwwssm"),
				},
			},
			ProjectID:    to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
			BranchID:     to.Ptr("rzsyrhpfbydxtfkpaa"),
			EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
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
	// res = armneonpostgres.EndpointsClientCreateOrUpdateResponse{
	// 	Endpoint: &armneonpostgres.Endpoint{
	// 		Properties: &armneonpostgres.EndpointProperties{
	// 			EntityID: to.Ptr("entity-id"),
	// 			EntityName: to.Ptr("entity-name"),
	// 			CreatedAt: to.Ptr("vhcilurdd"),
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			Attributes: []*armneonpostgres.Attributes{
	// 				{
	// 					Name: to.Ptr("trhvzyvaqy"),
	// 					Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 				},
	// 			},
	// 			ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 			BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 			EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 		},
	// 		ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org/projects/entity-name/branches/entity-name/endpoints/entity-name"),
	// 		Name: to.Ptr("btfgrpluxpodftoy"),
	// 		Type: to.Ptr("ldmicxcfahaewxfelpamtgb"),
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
