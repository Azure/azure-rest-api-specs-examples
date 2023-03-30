package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listjobs.json
func ExampleResourceClient_NewListJobsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.JobResponseListResult = armiothub.JobResponseListResult{
		// 	Value: []*armiothub.JobResponse{
		// 		{
		// 			Type: to.Ptr(armiothub.JobTypeUnknown),
		// 			EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 15 Jun 2017 19:20:58 GMT"); return t}()),
		// 			JobID: to.Ptr("test"),
		// 			StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "Thu, 15 Jun 2017 19:20:58 GMT"); return t}()),
		// 			Status: to.Ptr(armiothub.JobStatusUnknown),
		// 	}},
		// }
	}
}
