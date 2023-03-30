package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/chaos/resource-manager/Microsoft.Chaos/preview/2021-09-15-preview/examples/ListExperimentStatuses.json
func ExampleExperimentsClient_NewListAllStatusesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExperimentsClient().NewListAllStatusesPager("exampleRG", "exampleExperiment", nil)
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
		// page.ExperimentStatusListResult = armchaos.ExperimentStatusListResult{
		// 	Value: []*armchaos.ExperimentStatus{
		// 		{
		// 			Name: to.Ptr("50734542-2e64-4e08-814c-cc0e7475f7e4"),
		// 			Type: to.Ptr("Microsoft.Chaos/experiments/statuses"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/experiments/exampleExperiment/statuses/50734542-2e64-4e08-814c-cc0e7475f7e4"),
		// 			Properties: &armchaos.ExperimentStatusProperties{
		// 				CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
		// 				EndDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T01:00:00.0Z"); return t}()),
		// 				Status: to.Ptr("Successful"),
		// 			},
		// 	}},
		// }
	}
}
