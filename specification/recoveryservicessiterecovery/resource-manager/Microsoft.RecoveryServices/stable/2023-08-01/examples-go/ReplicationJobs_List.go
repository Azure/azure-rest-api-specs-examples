package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationJobs_List.json
func ExampleReplicationJobsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationJobsClient().NewListPager("vault1", "resourceGroupPS1", &armrecoveryservicessiterecovery.ReplicationJobsClientListOptions{Filter: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.JobCollection = armrecoveryservicessiterecovery.JobCollection{
		// 	Value: []*armrecoveryservicessiterecovery.Job{
		// 		{
		// 			Name: to.Ptr("1557d73f-6244-491e-8f0b-d300f752240b"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/1557d73f-6244-491e-8f0b-d300f752240b"),
		// 			Properties: &armrecoveryservicessiterecovery.JobProperties{
		// 				AllowedActions: []*string{
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T12:46:11.000Z"); return t}()),
		// 				Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
		// 				},
		// 				FriendlyName: to.Ptr("Create replication policy"),
		// 				ScenarioName: to.Ptr("AddProtectionProfile"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T12:46:04.641Z"); return t}()),
		// 				State: to.Ptr("Succeeded"),
		// 				StateDescription: to.Ptr("Completed"),
		// 				TargetInstanceType: to.Ptr("ProtectionProfile"),
		// 				TargetObjectID: to.Ptr("af095a1e-1f1b-5365-87c9-99162ebcfaf0"),
		// 				TargetObjectName: to.Ptr("protectionprofile1"),
		// 				Tasks: []*armrecoveryservicessiterecovery.ASRTask{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("0236416a-7573-4913-a4a1-6a286fbb1ceb"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/0236416a-7573-4913-a4a1-6a286fbb1ceb"),
		// 			Properties: &armrecoveryservicessiterecovery.JobProperties{
		// 				AllowedActions: []*string{
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T14:07:21.000Z"); return t}()),
		// 				Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
		// 				},
		// 				FriendlyName: to.Ptr("Register the Azure Site Recovery Provider"),
		// 				ScenarioName: to.Ptr("RegisterDra"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T14:07:19.278Z"); return t}()),
		// 				State: to.Ptr("Succeeded"),
		// 				StateDescription: to.Ptr("Completed"),
		// 				TargetInstanceType: to.Ptr("Server"),
		// 				TargetObjectID: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
		// 				TargetObjectName: to.Ptr("CP-B3L40406-12.ntdev.corp.microsoft.com"),
		// 				Tasks: []*armrecoveryservicessiterecovery.ASRTask{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("2438d560-80f0-420b-839e-5c8ee0af90a1"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/2438d560-80f0-420b-839e-5c8ee0af90a1"),
		// 			Properties: &armrecoveryservicessiterecovery.JobProperties{
		// 				AllowedActions: []*string{
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T05:56:16.000Z"); return t}()),
		// 				Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
		// 				},
		// 				FriendlyName: to.Ptr("Create a site"),
		// 				ScenarioName: to.Ptr("CreateSite"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T05:56:14.569Z"); return t}()),
		// 				State: to.Ptr("Succeeded"),
		// 				StateDescription: to.Ptr("Completed"),
		// 				TargetInstanceType: to.Ptr("Server"),
		// 				TargetObjectID: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
		// 				TargetObjectName: to.Ptr("cloud1"),
		// 				Tasks: []*armrecoveryservicessiterecovery.ASRTask{
		// 				},
		// 			},
		// 	}},
		// }
	}
}
