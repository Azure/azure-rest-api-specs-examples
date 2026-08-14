package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub/v2"
)

// Generated from example definition: 2026-05-01-preview/iothub_listjobs.json
func ExampleResourceClient_NewListJobsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("91d12660-3dec-467a-be2a-213b5544ddc0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceClient().NewListJobsPager("myResourceGroup", "testHub", nil)
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
		// page = armiothub.ResourceClientListJobsResponse{
		// 	JobResponseListResult: armiothub.JobResponseListResult{
		// 		Value: []*armiothub.JobResponse{
		// 			{
		// 				Type: to.Ptr(armiothub.JobTypeUnknown),
		// 				EndTimeUTC: to.Ptr(time.Date(2017, time.June, 15, 19, 20, 58, 0, time.UTC)),
		// 				JobID: to.Ptr("test"),
		// 				StartTimeUTC: to.Ptr(time.Date(2017, time.June, 15, 19, 20, 58, 0, time.UTC)),
		// 				Status: to.Ptr(armiothub.JobStatusUnknown),
		// 			},
		// 		},
		// 	},
		// }
	}
}
