package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseAutomaticTuningUpdateMin.json
func ExampleDatabaseAutomaticTuningClient_Update_updatesDatabaseAutomaticTuningSettingsWithMinimalProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseAutomaticTuningClient().Update(ctx, "default-sql-onebox", "testsvr11", "db1", armsql.DatabaseAutomaticTuning{
		Properties: &armsql.DatabaseAutomaticTuningProperties{
			DesiredState: to.Ptr(armsql.AutomaticTuningModeAuto),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseAutomaticTuning = armsql.DatabaseAutomaticTuning{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/automaticTuning"),
	// 	ID: to.Ptr("/subscriptions/c3aa9078-0000-0000-0000-e36f151182d7/resourceGroups/default-sql-onebox/providers/Microsoft.Sql/servers/testsvr11/databases/db1/automaticTuning/current"),
	// 	Properties: &armsql.DatabaseAutomaticTuningProperties{
	// 		ActualState: to.Ptr(armsql.AutomaticTuningModeAuto),
	// 		DesiredState: to.Ptr(armsql.AutomaticTuningModeAuto),
	// 		Options: map[string]*armsql.AutomaticTuningOptions{
	// 			"createIndex": &armsql.AutomaticTuningOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOn),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningDisabledReasonAutoConfigured),
	// 			},
	// 			"dropIndex": &armsql.AutomaticTuningOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOff),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningDisabledReasonAutoConfigured),
	// 			},
	// 			"forceLastGoodPlan": &armsql.AutomaticTuningOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOn),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningDisabledReasonAutoConfigured),
	// 			},
	// 			"maintainIndex": &armsql.AutomaticTuningOptions{
	// 				ActualState: to.Ptr(armsql.AutomaticTuningOptionModeActualOff),
	// 				DesiredState: to.Ptr(armsql.AutomaticTuningOptionModeDesiredDefault),
	// 				ReasonCode: to.Ptr[int32](2),
	// 				ReasonDesc: to.Ptr(armsql.AutomaticTuningDisabledReasonAutoConfigured),
	// 			},
	// 		},
	// 	},
	// }
}
