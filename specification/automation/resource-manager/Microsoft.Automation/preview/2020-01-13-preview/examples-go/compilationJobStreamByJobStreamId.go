package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/compilationJobStreamByJobStreamId.json
func ExampleDscCompilationJobClient_GetStream() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscCompilationJobClient().GetStream(ctx, "rg", "myAutomationAccount33", "836d4e06-2d88-46b4-8500-7febd4906838", "836d4e06-2d88-46b4-8500-7febd4906838_00636481062421684835_00000000000000000008", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobStream = armautomation.JobStream{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062421684835_00000000000000000008"),
	// 	Properties: &armautomation.JobStreamProperties{
	// 		JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838:00636481062421684835:00000000000000000001"),
	// 		StreamText: to.Ptr(""),
	// 		StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
	// 		Summary: to.Ptr(""),
	// 		Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.1684835+00:00"); return t}()),
	// 		Value: map[string]any{
	// 			"PSComputerName": "localhost",
	// 			"PSShowComputerName": true,
	// 			"PSSourceJobInstanceId": "836d4e06-2d88-46b4-8500-7febd4906838",
	// 			"value": "",
	// 		},
	// 	},
	// }
}
