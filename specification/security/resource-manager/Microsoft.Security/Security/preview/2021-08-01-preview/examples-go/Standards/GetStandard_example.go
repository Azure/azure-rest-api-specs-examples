package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2021-08-01-preview/Standards/GetStandard_example.json
func ExampleStandardsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandardsClient().Get(ctx, "myResourceGroup", "21300918-b2e3-0346-785f-c77ff57d243b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurity.StandardsClientGetResponse{
	// 	Standard: armsecurity.Standard{
	// 		Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 		Type: to.Ptr("Microsoft.Security/standards"),
	// 		Etag: to.Ptr("etag value"),
	// 		ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myResourceGroup/providers/Microsoft.Security/standards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armsecurity.StandardProperties{
	// 			Description: to.Ptr("Controls and security data implementing security recommendations defined in Azure Security"),
	// 			Category: to.Ptr("SecurityCenter"),
	// 			Components: []*armsecurity.StandardComponentProperties{
	// 				{
	// 					Key: to.Ptr("1195afff-c881-495e-9bc5-1486211ae03f"),
	// 				},
	// 				{
	// 					Key: to.Ptr("dbd0cb49-b563-45e7-9724-889e799fa648"),
	// 				},
	// 			},
	// 			DisplayName: to.Ptr("Cross cloud standard 1"),
	// 			StandardType: to.Ptr("Custom"),
	// 			SupportedClouds: []*armsecurity.StandardSupportedClouds{
	// 				to.Ptr(armsecurity.StandardSupportedCloudsGCP),
	// 			},
	// 		},
	// 		SystemData: &armsecurity.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
