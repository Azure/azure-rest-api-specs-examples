package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/ParameterGroups_Update.json
func ExampleParameterGroupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewParameterGroupsClient().BeginUpdate(ctx, "exampleresourcegroup", "exampleparametergroup", armhorizondb.ParameterGroupForPatchUpdate{
		Tags: map[string]*string{
			"team": to.Ptr("updated-data-platform"),
		},
		Properties: &armhorizondb.ParameterGroupPropertiesForPatchUpdate{
			Parameters: []*armhorizondb.ParameterProperties{
				{
					Name:  to.Ptr("max_connections"),
					Value: to.Ptr("300"),
				},
				{
					Name:  to.Ptr("log_min_error_statement"),
					Value: to.Ptr("warning"),
				},
			},
			Description:      to.Ptr("Updated parameter group for high-throughput workloads"),
			ApplyImmediately: to.Ptr(true),
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
	// res = armhorizondb.ParameterGroupsClientUpdateResponse{
	// 	ParameterGroup: &armhorizondb.ParameterGroup{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/parameterGroups/exampleparametergroup"),
	// 		Name: to.Ptr("exampleparametergroup"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/parameterGroups"),
	// 		Location: to.Ptr("westus2"),
	// 		Tags: map[string]*string{
	// 			"team": to.Ptr("updated-data-platform"),
	// 		},
	// 		Properties: &armhorizondb.ParameterGroupProperties{
	// 			Parameters: []*armhorizondb.ParameterProperties{
	// 				{
	// 					Name: to.Ptr("max_connections"),
	// 					Description: to.Ptr("Sets the maximum number of concurrent connections"),
	// 					Value: to.Ptr("300"),
	// 					DataType: to.Ptr("integer"),
	// 					AllowedValues: to.Ptr("100,200,500,1000"),
	// 					IsDynamic: to.Ptr(true),
	// 					IsReadOnly: to.Ptr(false),
	// 					DocumentationLink: to.Ptr("https://docs.example.com/max_connections"),
	// 					Unit: to.Ptr("connections"),
	// 				},
	// 				{
	// 					Name: to.Ptr("log_min_error_statement"),
	// 					Description: to.Ptr("Controls which statements are logged"),
	// 					Value: to.Ptr("warning"),
	// 					DataType: to.Ptr("string"),
	// 					AllowedValues: to.Ptr("debug,info,notice,warning,error,log,fatal,panic"),
	// 					IsDynamic: to.Ptr(true),
	// 					IsReadOnly: to.Ptr(false),
	// 					DocumentationLink: to.Ptr("https://docs.example.com/log_min_error_statement"),
	// 				},
	// 			},
	// 			Description: to.Ptr("Updated parameter group for high-throughput workloads"),
	// 			ApplyImmediately: to.Ptr(true),
	// 			PgVersion: to.Ptr[int32](17),
	// 			Version: to.Ptr[int32](22),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
