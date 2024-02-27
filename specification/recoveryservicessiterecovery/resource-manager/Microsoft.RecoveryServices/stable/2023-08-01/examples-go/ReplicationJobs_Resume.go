package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationJobs_Resume.json
func ExampleReplicationJobsClient_BeginResume() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationJobsClient().BeginResume(ctx, "vault1", "resourceGroupPS1", "58776d0b-3141-48b2-a377-9ad863eb160d", armrecoveryservicessiterecovery.ResumeJobParams{
		Properties: &armrecoveryservicessiterecovery.ResumeJobParamsProperties{
			Comments: to.Ptr(" "),
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
	// res.Job = armrecoveryservicessiterecovery.Job{
	// 	Name: to.Ptr("58776d0b-3141-48b2-a377-9ad863eb160d"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/58776d0b-3141-48b2-a377-9ad863eb160d"),
	// 	Properties: &armrecoveryservicessiterecovery.JobProperties{
	// 		ActivityID: to.Ptr("1b808dfe-0451-44ac-894c-c7270711cd8c ActivityId: 9f6f849e-922a-43ec-a7a6-0be45fc85c56"),
	// 		AllowedActions: []*string{
	// 			to.Ptr("Cancel"),
	// 			to.Ptr("Resume")},
	// 			CustomDetails: &armrecoveryservicessiterecovery.TestFailoverJobDetails{
	// 				AffectedObjectDetails: map[string]*string{
	// 					"PrimaryCloudId": to.Ptr("cloud_6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 					"PrimaryCloudName": to.Ptr("cloud1"),
	// 					"PrimaryFabricProviderId": to.Ptr("HyperVSite"),
	// 					"PrimaryVmId": to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 					"PrimaryVmName": to.Ptr("vm1"),
	// 					"PrimaryVmmId": to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 					"PrimaryVmmName": to.Ptr("cloud1"),
	// 					"ProtectionProfileId": to.Ptr("af095a1e-1f1b-5365-87c9-99162ebcfaf0"),
	// 					"RecoveryCloudId": to.Ptr("d38048d4-b460-4791-8ece-108395ee8478"),
	// 					"RecoveryCloudName": to.Ptr("Microsoft Azure"),
	// 					"RecoveryFabricProviderId": to.Ptr("Azure"),
	// 					"RecoveryVmId": to.Ptr(""),
	// 					"RecoveryVmName": to.Ptr("vm1"),
	// 					"RecoveryVmmId": to.Ptr("21a9403c-6ec1-44f2-b744-b4e50b792387"),
	// 					"RecoveryVmmName": to.Ptr("Microsoft Azure"),
	// 				},
	// 				InstanceType: to.Ptr("TestFailoverJobDetails"),
	// 				Comments: to.Ptr(" "),
	// 				NetworkFriendlyName: to.Ptr("vnetavrai"),
	// 				NetworkName: to.Ptr("vnetavrai"),
	// 				NetworkType: to.Ptr("VmNetworkAsInput"),
	// 				ProtectedItemDetails: []*armrecoveryservicessiterecovery.FailoverReplicationProtectedItemDetails{
	// 					{
	// 						Name: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 						FriendlyName: to.Ptr("vm1"),
	// 						NetworkConnectionStatus: to.Ptr("Connected"),
	// 						NetworkFriendlyName: to.Ptr("vnetavrai"),
	// 						Subnet: to.Ptr("Subnet1"),
	// 						TestVMFriendlyName: to.Ptr("vm1-test"),
	// 						TestVMName: to.Ptr("vm1-test"),
	// 				}},
	// 				TestFailoverStatus: to.Ptr("Completed"),
	// 			},
	// 			Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
	// 			},
	// 			FriendlyName: to.Ptr("Test failover"),
	// 			ScenarioName: to.Ptr("TestFailover"),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-25T09:57:57.035Z"); return t}()),
	// 			State: to.Ptr("Suspended"),
	// 			StateDescription: to.Ptr("WaitingForStopTestFailover"),
	// 			TargetInstanceType: to.Ptr("ProtectionEntity"),
	// 			TargetObjectID: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 			TargetObjectName: to.Ptr("vm1"),
	// 			Tasks: []*armrecoveryservicessiterecovery.ASRTask{
	// 			},
	// 		},
	// 	}
}
