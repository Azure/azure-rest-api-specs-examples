package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAutomaticTuningUpdateMin.json
func ExampleServerAutomaticTuningClient_Update_updatesServerAutomaticTuningSettingsWithMinimalProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerAutomaticTuningClient().Update(ctx, "default-sql-onebox", "testsvr11", armsql.ServerAutomaticTuning{
		Properties: &armsql.AutomaticTuningServerProperties{
			DesiredState: to.Ptr(armsql.AutomaticTuningServerModeAuto),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerAutomaticTuning = armsql.ServerAutomaticTuning{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/automaticTuning"),
	// 	ID: to.Ptr("/subscriptions/c3aa9078-0000-0000-0000-e36f151182d7/resourceGroups/default-sql-onebox/providers/Microsoft.Sql/servers/testsvr11/automaticTuning/current"),
	// 	Properties: &armsql.AutomaticTuningServerProperties{
	// 		ActualState: to.Ptr(armsql.AutomaticTuningServerModeAuto),
	// 		DesiredState: to.Ptr(armsql.AutomaticTuningServerModeAuto),
	// 		Options: map[string]*armsql.AutomaticTuningServerOptions{
	// 			"createIndex": &armsql.AutomaticTuningServerOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOn),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningServerReasonAutoConfigured),
	// 			},
	// 			"dropIndex": &armsql.AutomaticTuningServerOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOff),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningServerReasonAutoConfigured),
	// 			},
	// 			"forceLastGoodPlan": &armsql.AutomaticTuningServerOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOn),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningServerReasonAutoConfigured),
	// 			},
	// 			"maintainIndex": &armsql.AutomaticTuningServerOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOff),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningServerReasonAutoConfigured),
	// 			},
	// 		},
	// 	},
	// }
}
