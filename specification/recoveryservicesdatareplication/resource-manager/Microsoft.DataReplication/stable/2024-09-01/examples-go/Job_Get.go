package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/Job_Get.json
func ExampleJobClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobClient().Get(ctx, "rgrecoveryservicesdatareplication", "4", "ZGH4y", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesdatareplication.JobClientGetResponse{
	// 	JobModel: &armrecoveryservicesdatareplication.JobModel{
	// 		Properties: &armrecoveryservicesdatareplication.JobModelProperties{
	// 			DisplayName: to.Ptr("dhopzytkd"),
	// 			State: to.Ptr(armrecoveryservicesdatareplication.JobStatePending),
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
	// 			ObjectID: to.Ptr("wvtmwiyxqrpqvljzn"),
	// 			ObjectName: to.Ptr("ieieqaw"),
	// 			ObjectInternalID: to.Ptr("ahbtlwmbeivmlbj"),
	// 			ObjectInternalName: to.Ptr("mxikyrinkeyj"),
	// 			ObjectType: to.Ptr(armrecoveryservicesdatareplication.JobObjectTypeAvsDiskPool),
	// 			ReplicationProviderID: to.Ptr("ghxsbnvdkx"),
	// 			SourceFabricProviderID: to.Ptr("yqlertkmzdsgsplzgkmwcttsiagsa"),
	// 			TargetFabricProviderID: to.Ptr("stjlmqzpgnrug"),
	// 			AllowedActions: []*string{
	// 				to.Ptr("mfsyvxzgmcpgdzfbbhoxrzhya"),
	// 			},
	// 			ActivityID: to.Ptr("esjvxsa"),
	// 			Tasks: []*armrecoveryservicesdatareplication.TaskModel{
	// 				{
	// 					TaskName: to.Ptr("flkzfbbpngqbbjsdqysqfon"),
	// 					State: to.Ptr(armrecoveryservicesdatareplication.TaskStatePending),
	// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
	// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
	// 					CustomProperties: &armrecoveryservicesdatareplication.TaskModelCustomProperties{
	// 						InstanceType: to.Ptr("aaqgqvnhskxpsbnrdekxaghweon"),
	// 					},
	// 				},
	// 			},
	// 			Errors: []*armrecoveryservicesdatareplication.ErrorModel{
	// 				{
	// 					Code: to.Ptr("ndcxzieiuwoxoklilcvjmglml"),
	// 					Type: to.Ptr("iyktxohrtrkshbjdhboscsu"),
	// 					Severity: to.Ptr("ldgmfmbzmjtqjg"),
	// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
	// 					Message: to.Ptr("ltuufmxomfecqeyvzrfjqlelytkdwr"),
	// 					Causes: to.Ptr("iffxig"),
	// 					Recommendation: to.Ptr("brridkskflo"),
	// 				},
	// 			},
	// 			CustomProperties: &armrecoveryservicesdatareplication.FailoverJobModelCustomProperties{
	// 				InstanceType: to.Ptr("JobModelCustomProperties"),
	// 				AffectedObjectDetails: &armrecoveryservicesdatareplication.AffectedObjectDetails{
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/jobs/job1"),
	// 		Name: to.Ptr("ukii"),
	// 		Type: to.Ptr("cswros"),
	// 		SystemData: &armrecoveryservicesdatareplication.SystemData{
	// 			CreatedBy: to.Ptr("jurgsdagntjg"),
	// 			CreatedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("zowfl")),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.270Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("vuw"),
	// 			LastModifiedByType: to.Ptr(armrecoveryservicesdatareplication.CreatedByType("h")),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.270Z"); return t}()),
	// 		},
	// 	},
	// }
}
