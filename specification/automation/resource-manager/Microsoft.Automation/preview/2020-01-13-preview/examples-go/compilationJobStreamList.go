package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/compilationJobStreamList.json
func ExampleDscCompilationJobStreamClient_ListByJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscCompilationJobStreamClient().ListByJob(ctx, "rg", "myAutomationAccount33", "836d4e06-2d88-46b4-8500-7febd4906838", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobStreamListResult = armautomation.JobStreamListResult{
	// 	Value: []*armautomation.JobStream{
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062355996678_00000000000000000001"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062355996678_00000000000000000001"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeWarning),
	// 				Summary: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838:[localhost]:The 'Microsoft.PowerShell.Management' module was not imported because the 'Microsoft.PowerShell.Management' snap-in was already imported."),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:35.599Z"); return t}()),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062380840740_00000000000000000002"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062380840740_00000000000000000002"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeWarning),
	// 				Summary: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838:[localhost]:The configuration 'NewDscConfiguration' is loading one or more built-in resources without explicitly importing associated modules. Add Import-DscResource –ModuleName 'PSDesiredStateConfiguration' to your configuration to avoid this message."),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:38.084Z"); return t}()),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062384590127_00000000000000000003"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062384590127_00000000000000000003"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeWarning),
	// 				Summary: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838:[localhost]:The 'Microsoft.PowerShell.Management' module was not imported because the 'Microsoft.PowerShell.Management' snap-in was already imported."),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:38.459Z"); return t}()),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062387245395_00000000000000000004"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062387245395_00000000000000000004"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeWarning),
	// 				Summary: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838:[localhost]:The configuration 'NewDscConfiguration' is loading one or more built-in resources without explicitly importing associated modules. Add Import-DscResource –ModuleName 'PSDesiredStateConfiguration' to your configuration to avoid this message."),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:38.724Z"); return t}()),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062417091181_00000000000000000005"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062417091181_00000000000000000005"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
	// 				Summary: to.Ptr(""),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:41.709Z"); return t}()),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062418809632_00000000000000000006"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062418809632_00000000000000000006"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
	// 				Summary: to.Ptr(""),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:41.880Z"); return t}()),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062420371712_00000000000000000007"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062420371712_00000000000000000007"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
	// 				Summary: to.Ptr(""),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.037Z"); return t}()),
	// 			},
	// 		},
	// 		{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062421684835_00000000000000000008"),
	// 			Properties: &armautomation.JobStreamProperties{
	// 				JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838_00636481062421684835_00000000000000000008"),
	// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
	// 				Summary: to.Ptr(""),
	// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.168Z"); return t}()),
	// 			},
	// 	}},
	// }
}
