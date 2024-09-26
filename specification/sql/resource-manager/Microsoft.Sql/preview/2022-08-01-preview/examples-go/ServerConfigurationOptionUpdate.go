package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ServerConfigurationOptionUpdate.json
func ExampleServerConfigurationOptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerConfigurationOptionsClient().BeginCreateOrUpdate(ctx, "testrg", "testinstance", armsql.ServerConfigurationOptionNameAllowPolybaseExport, armsql.ServerConfigurationOption{
		Properties: &armsql.ServerConfigurationOptionProperties{
			ServerConfigurationOptionValue: to.Ptr[int32](1),
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
	// res.ServerConfigurationOption = armsql.ServerConfigurationOption{
	// 	Name: to.Ptr("allowPolybaseExport"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/serverConfigurationOptions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/serverConfigurationOptions/allowPolybaseExport"),
	// 	Properties: &armsql.ServerConfigurationOptionProperties{
	// 		ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 		ServerConfigurationOptionValue: to.Ptr[int32](1),
	// 	},
	// }
}
