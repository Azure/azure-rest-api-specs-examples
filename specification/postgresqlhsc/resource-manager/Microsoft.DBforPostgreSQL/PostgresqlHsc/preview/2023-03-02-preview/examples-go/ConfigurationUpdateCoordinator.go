package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: 2023-03-02-preview/ConfigurationUpdateCoordinator.json
func ExampleConfigurationsClient_BeginUpdateOnCoordinator() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConfigurationsClient().BeginUpdateOnCoordinator(ctx, "TestResourceGroup", "testcluster", "array_nulls", armpostgresqlhsc.ServerConfiguration{
		Properties: &armpostgresqlhsc.ServerConfigurationProperties{
			Value: to.Ptr("on"),
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
	// res = armpostgresqlhsc.ConfigurationsClientUpdateOnCoordinatorResponse{
	// 	ServerConfiguration: &armpostgresqlhsc.ServerConfiguration{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/configurations/array_nulls"),
	// 		Name: to.Ptr("array_nulls"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/configurations/coordinator"),
	// 		SystemData: &armpostgresqlhsc.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 		},
	// 		Properties: &armpostgresqlhsc.ServerConfigurationProperties{
	// 			AllowedValues: to.Ptr("on,off"),
	// 			DataType: to.Ptr(armpostgresqlhsc.ConfigurationDataTypeBoolean),
	// 			DefaultValue: to.Ptr("on"),
	// 			Description: to.Ptr("Enables input of NULL elements in arrays."),
	// 			ProvisioningState: to.Ptr(armpostgresqlhsc.ProvisioningStateInProgress),
	// 			RequiresRestart: to.Ptr(false),
	// 			Source: to.Ptr("user-override"),
	// 			Value: to.Ptr("on"),
	// 		},
	// 	},
	// }
}
