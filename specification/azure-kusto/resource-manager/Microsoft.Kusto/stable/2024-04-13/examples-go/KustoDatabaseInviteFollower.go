package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8adce17dc500f338f86f18af30aac61dcb71c5f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDatabaseInviteFollower.json
func ExampleDatabaseClient_InviteFollower() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseClient().InviteFollower(ctx, "kustorptest", "kustoCluster", "database", armkusto.DatabaseInviteFollowerRequest{
		InviteeEmail: to.Ptr("invitee@contoso.com"),
		TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
			ExternalTablesToExclude: []*string{},
			ExternalTablesToInclude: []*string{
				to.Ptr("ExternalTable*")},
			FunctionsToExclude: []*string{
				to.Ptr("functionsToExclude2")},
			FunctionsToInclude: []*string{
				to.Ptr("functionsToInclude1")},
			MaterializedViewsToExclude: []*string{
				to.Ptr("MaterializedViewTable2")},
			MaterializedViewsToInclude: []*string{
				to.Ptr("MaterializedViewTable1")},
			TablesToExclude: []*string{
				to.Ptr("Table2")},
			TablesToInclude: []*string{
				to.Ptr("Table1")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseInviteFollowerResult = armkusto.DatabaseInviteFollowerResult{
	// 	GeneratedInvitation: to.Ptr("eyJ0eXAiOiJKVInvitationToken"),
	// }
}
