package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControlSyncJobStreams/getSourceControlSyncJobStreams.json
func ExampleSourceControlSyncJobStreamsClient_NewListBySyncJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSourceControlSyncJobStreamsClient().NewListBySyncJobPager("rg", "myAutomationAccount33", "MySourceControl", "ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2b", &armautomation.SourceControlSyncJobStreamsClientListBySyncJobOptions{Filter: nil})
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
		// page.SourceControlSyncJobStreamsListBySyncJob = armautomation.SourceControlSyncJobStreamsListBySyncJob{
		// 	Value: []*armautomation.SourceControlSyncJobStream{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2b/streams/b86c5c31-e9fd-4734-8764-ddd6c101e706_00636596855134810785_00000000000000000005"),
		// 			Properties: &armautomation.SourceControlSyncJobStreamProperties{
		// 				SourceControlSyncJobStreamID: to.Ptr("b86c5c31-e9fd-4734-8764-ddd6c101e706_00636596855134810785_00000000000000000005"),
		// 				StreamType: to.Ptr(armautomation.StreamTypeError),
		// 				Summary: to.Ptr("ForbiddenError: The server failed to authenticate the request. Verify that the certificate is valid and is associated with this subscription."),
		// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2b/streams/b86c5c31-e9fd-4734-8764-ddd6c101e706_00636596855136998262_00000000000000000006"),
		// 			Properties: &armautomation.SourceControlSyncJobStreamProperties{
		// 				SourceControlSyncJobStreamID: to.Ptr("b86c5c31-e9fd-4734-8764-ddd6c101e706_00636596855136998262_00000000000000000006"),
		// 				StreamType: to.Ptr(armautomation.StreamTypeError),
		// 				Summary: to.Ptr("System.Management.Automation.RuntimeException: Cannot index into a null array.\r\n   at CallSite.Target(Closure , CallSite , Object , Int32 )\r\n   at System.Dynamic.UpdateDelegates.UpdateAndExecute2[T0,T1,TRet](CallSite site, T0 arg0, T1 arg1)\r\n   at System.Management.Automation.Interpreter.DynamicInstruction`3.Run(InterpretedFrame frame)\r\n   at System.Management.Automation.Interpreter.EnterTryCatchFinallyInstruction.Run(InterpretedFrame frame)"),
		// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:27.903+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/sourceControls/MySourceControl/sourceControlSyncJobs/ce6fe3e3-9db3-4096-a6b4-82bfb4c10a2b/streams/b86c5c31-e9fd-4734-8764-ddd6c101e706_00636596855139029522_00000000000000000007"),
		// 			Properties: &armautomation.SourceControlSyncJobStreamProperties{
		// 				SourceControlSyncJobStreamID: to.Ptr("b86c5c31-e9fd-4734-8764-ddd6c101e706_00636596855139029522_00000000000000000007"),
		// 				StreamType: to.Ptr(armautomation.StreamTypeError),
		// 				Summary: to.Ptr("System.Management.Automation.ParameterBindingValidationException: Cannot validate argument on parameter 'Location'. The argument is null or empty. Provide an argument that is not null or empty, and then try the command again. ---> System.Management.Automation.ValidationMetadataException: The argument is null or empty. Provide an argument that is not null or empty, and then try the command again.\r\n   at System.Management.Automation.ValidateNotNullOrEmptyAttribute.Validate(Object arguments, EngineIntrinsics engineIntrinsics)"),
		// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:28.903+00:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
