package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/PrivateAccesses_Get_Get_A_Private_Access_Resource.json
func ExamplePrivateAccessesClient_Get_getAPrivateAccessResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateAccessesClient().Get(ctx, "myResourceGroup", "myPrivateAccess", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.PrivateAccessesClientGetResponse{
	// 	PrivateAccess: armchaos.PrivateAccess{
	// 		Name: to.Ptr("myPrivateAccess"),
	// 		Type: to.Ptr("Microsoft.Chaos/privateAccesses"),
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/myResourcegroup/providers/Microsoft.Chaos/privateAccesses/myPrivateAccess"),
	// 		Location: to.Ptr("centraluseuap"),
	// 		Properties: &armchaos.PrivateAccessProperties{
	// 		},
	// 		SystemData: &armchaos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
	// 		},
	// 	},
	// }
}
