package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationJobs_Get.json
func ExampleReplicationJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationJobsClient().Get(ctx, "vault1", "resourceGroupPS1", "58776d0b-3141-48b2-a377-9ad863eb160d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Job = armrecoveryservicessiterecovery.Job{
	// 	Name: to.Ptr("32ea4b9e-de62-49a1-b062-7864d5c3b897"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/32ea4b9e-de62-49a1-b062-7864d5c3b897"),
	// 	Properties: &armrecoveryservicessiterecovery.JobProperties{
	// 		ActivityID: to.Ptr("fc8e9c8f-0e76-4b6b-8e7e-d37c1b31eba0 ActivityId: c506b6ba-0711-411e-8b09-1f3f4dcb824b"),
	// 		AllowedActions: []*string{
	// 		},
	// 		CustomDetails: &armrecoveryservicessiterecovery.AsrJobDetails{
	// 			AffectedObjectDetails: map[string]*string{
	// 				"PrimaryFabricProviderId": to.Ptr("HyperVSite"),
	// 				"PrimaryVmmId": to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 				"PrimaryVmmName": to.Ptr("cloud1"),
	// 				"RecoveryFabricProviderId": to.Ptr("Azure"),
	// 				"RecoveryVmmId": to.Ptr("21a9403c-6ec1-44f2-b744-b4e50b792387"),
	// 			},
	// 			InstanceType: to.Ptr("AsrJobDetails"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T11:25:57.000Z"); return t}()),
	// 		Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
	// 		},
	// 		FriendlyName: to.Ptr("Delete a recovery plan"),
	// 		ScenarioName: to.Ptr("DeleteRecoveryPlan"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T11:25:56.800Z"); return t}()),
	// 		State: to.Ptr("Succeeded"),
	// 		StateDescription: to.Ptr("Completed"),
	// 		TargetInstanceType: to.Ptr("RecoveryPlan"),
	// 		TargetObjectID: to.Ptr("966c33bb-66e7-4567-9786-f80b0694f5f9"),
	// 		TargetObjectName: to.Ptr("RPtest1"),
	// 		Tasks: []*armrecoveryservicessiterecovery.ASRTask{
	// 			{
	// 				Name: to.Ptr("DeleteRecoveryPlanTask"),
	// 				AllowedActions: []*string{
	// 				},
	// 				CustomDetails: &armrecoveryservicessiterecovery.ManualActionTaskDetails{
	// 					InstanceType: to.Ptr("ManualActionTaskDetails"),
	// 				},
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T11:25:57.318Z"); return t}()),
	// 				Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
	// 				},
	// 				FriendlyName: to.Ptr("Delete a recovery plan task"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T11:25:57.302Z"); return t}()),
	// 				State: to.Ptr("Succeeded"),
	// 				StateDescription: to.Ptr("Completed"),
	// 				TaskID: to.Ptr("763326a2-01c9-4257-b2a1-0aac56465014"),
	// 				TaskType: to.Ptr("TaskDetails"),
	// 		}},
	// 	},
	// }
}
