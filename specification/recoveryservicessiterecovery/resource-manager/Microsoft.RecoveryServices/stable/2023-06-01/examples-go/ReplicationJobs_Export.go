package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationJobs_Export.json
func ExampleReplicationJobsClient_BeginExport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationJobsClient().BeginExport(ctx, "vault1", "resourceGroupPS1", armrecoveryservicessiterecovery.JobQueryParameter{
		AffectedObjectTypes: to.Ptr(""),
		EndTime:             to.Ptr("2017-05-04T14:26:51.9161395Z"),
		JobStatus:           to.Ptr(""),
		StartTime:           to.Ptr("2017-04-27T14:26:51.9161395Z"),
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
	// 	Name: to.Ptr("37e0fc2b-13f2-4817-aafa-0cd807d46842"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/37e0fc2b-13f2-4817-aafa-0cd807d46842"),
	// 	Properties: &armrecoveryservicessiterecovery.JobProperties{
	// 		ActivityID: to.Ptr("36841d27-34f6-49ad-b572-e7dc263f100b-2017-05-04 14:26:47Z-Ibz ActivityId: c124df21-7661-4541-b32a-3c723ebbb045"),
	// 		AllowedActions: []*string{
	// 		},
	// 		CustomDetails: &armrecoveryservicessiterecovery.ExportJobDetails{
	// 			AffectedObjectDetails: map[string]*string{
	// 			},
	// 			InstanceType: to.Ptr("ExportJobDetails"),
	// 			BlobURI: to.Ptr("<blobUri>"),
	// 			SasToken: to.Ptr("<sasToken>"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-04T14:33:43.000Z"); return t}()),
	// 		Errors: []*armrecoveryservicessiterecovery.JobErrorDetails{
	// 		},
	// 		ScenarioName: to.Ptr("ExportJobs"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-04T14:33:42.276Z"); return t}()),
	// 		State: to.Ptr("Succeeded"),
	// 		StateDescription: to.Ptr("Completed"),
	// 		TargetInstanceType: to.Ptr("Other"),
	// 		TargetObjectID: to.Ptr(""),
	// 		TargetObjectName: to.Ptr(""),
	// 		Tasks: []*armrecoveryservicessiterecovery.ASRTask{
	// 		},
	// 	},
	// }
}
