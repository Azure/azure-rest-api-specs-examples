package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Workflow_List.json
func ExampleWorkflowClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowClient().NewListPager("rgrecoveryservicesdatareplication", "4", &armrecoveryservicesdatareplication.WorkflowClientListOptions{Filter: to.Ptr("mnebpgmjcitjleipnttx"),
		ContinuationToken: to.Ptr("rdavrzbethhslmkqgajontnxsue"),
	})
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
		// page.WorkflowModelCollection = armrecoveryservicesdatareplication.WorkflowModelCollection{
		// 	Value: []*armrecoveryservicesdatareplication.WorkflowModel{
		// 		{
		// 			Name: to.Ptr("ukii"),
		// 			Type: to.Ptr("cswros"),
		// 			ID: to.Ptr("zsyhbwznbkaia"),
		// 			Properties: &armrecoveryservicesdatareplication.WorkflowModelProperties{
		// 				ActivityID: to.Ptr("esjvxsa"),
		// 				AllowedActions: []*string{
		// 					to.Ptr("mfsyvxzgmcpgdzfbbhoxrzhya")},
		// 					CustomProperties: &armrecoveryservicesdatareplication.WorkflowModelCustomProperties{
		// 						AffectedObjectDetails: map[string]*string{
		// 							"key7245": to.Ptr("yllr"),
		// 						},
		// 						InstanceType: to.Ptr("WorkflowModelCustomProperties"),
		// 					},
		// 					DisplayName: to.Ptr("dhopzytkd"),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
		// 					Errors: []*armrecoveryservicesdatareplication.ErrorModel{
		// 						{
		// 							Type: to.Ptr("iyktxohrtrkshbjdhboscsu"),
		// 							Causes: to.Ptr("iffxig"),
		// 							Code: to.Ptr("ndcxzieiuwoxoklilcvjmglml"),
		// 							CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
		// 							Message: to.Ptr("ltuufmxomfecqeyvzrfjqlelytkdwr"),
		// 							Recommendation: to.Ptr("brridkskflo"),
		// 							Severity: to.Ptr("ldgmfmbzmjtqjg"),
		// 					}},
		// 					ObjectID: to.Ptr("wvtmwiyxqrpqvljzn"),
		// 					ObjectInternalID: to.Ptr("ahbtlwmbeivmlbj"),
		// 					ObjectInternalName: to.Ptr("mxikyrinkeyj"),
		// 					ObjectName: to.Ptr("ieieqaw"),
		// 					ObjectType: to.Ptr(armrecoveryservicesdatareplication.WorkflowObjectTypeAvsDiskPool),
		// 					ReplicationProviderID: to.Ptr("ghxsbnvdkx"),
		// 					SourceFabricProviderID: to.Ptr("yqlertkmzdsgsplzgkmwcttsiagsa"),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
		// 					State: to.Ptr(armrecoveryservicesdatareplication.WorkflowStatePending),
		// 					TargetFabricProviderID: to.Ptr("stjlmqzpgnrug"),
		// 					Tasks: []*armrecoveryservicesdatareplication.TaskModel{
		// 						{
		// 							ChildrenWorkflows: []*armrecoveryservicesdatareplication.WorkflowModel{
		// 							},
		// 							CustomProperties: &armrecoveryservicesdatareplication.TaskModelCustomProperties{
		// 								InstanceType: to.Ptr("aaqgqvnhskxpsbnrdekxaghweon"),
		// 							},
		// 							EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.269Z"); return t}()),
		// 							State: to.Ptr(armrecoveryservicesdatareplication.TaskStatePending),
		// 							TaskName: to.Ptr("flkzfbbpngqbbjsdqysqfon"),
		// 					}},
		// 				},
		// 				SystemData: &armrecoveryservicesdatareplication.WorkflowModelSystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.270Z"); return t}()),
		// 					CreatedBy: to.Ptr("jurgsdagntjg"),
		// 					CreatedByType: to.Ptr("zowfl"),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:59.270Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("vuw"),
		// 					LastModifiedByType: to.Ptr("h"),
		// 				},
		// 		}},
		// 	}
	}
}
