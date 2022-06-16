package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAutomaticTuningUpdateMax.json
func ExampleServerAutomaticTuningClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewServerAutomaticTuningClient("c3aa9078-0000-0000-0000-e36f151182d7", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"default-sql-onebox",
		"testsvr11",
		armsql.ServerAutomaticTuning{
			Properties: &armsql.AutomaticTuningServerProperties{
				DesiredState: to.Ptr(armsql.AutomaticTuningServerModeAuto),
				Options: map[string]*armsql.AutomaticTuningServerOptions{
					"createIndex": {
						DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredOff),
					},
					"dropIndex": {
						DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredOn),
					},
					"forceLastGoodPlan": {
						DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
