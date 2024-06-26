package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e79a5aa63c0551c1b5af1d2853cceb495283d334/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/TroubleshootSqlVirtualMachine.json
func ExampleTroubleshootClient_BeginTroubleshoot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTroubleshootClient().BeginTroubleshoot(ctx, "testrg", "testvm", armsqlvirtualmachine.SQLVMTroubleshooting{
		EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-09T22:10:00.000Z"); return t }()),
		Properties: &armsqlvirtualmachine.TroubleshootingAdditionalProperties{
			UnhealthyReplicaInfo: &armsqlvirtualmachine.UnhealthyReplicaInfo{
				AvailabilityGroupName: to.Ptr("AG1"),
			},
		},
		StartTimeUTC:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-09T17:10:00.000Z"); return t }()),
		TroubleshootingScenario: to.Ptr(armsqlvirtualmachine.TroubleshootingScenarioUnhealthyReplica),
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
	// res.SQLVMTroubleshooting = armsqlvirtualmachine.SQLVMTroubleshooting{
	// 	EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-09T22:10:00.000Z"); return t}()),
	// 	Properties: &armsqlvirtualmachine.TroubleshootingAdditionalProperties{
	// 		UnhealthyReplicaInfo: &armsqlvirtualmachine.UnhealthyReplicaInfo{
	// 			AvailabilityGroupName: to.Ptr("AG1"),
	// 		},
	// 	},
	// 	StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-09T17:10:00.000Z"); return t}()),
	// 	TroubleshootingScenario: to.Ptr(armsqlvirtualmachine.TroubleshootingScenarioUnhealthyReplica),
	// }
}
