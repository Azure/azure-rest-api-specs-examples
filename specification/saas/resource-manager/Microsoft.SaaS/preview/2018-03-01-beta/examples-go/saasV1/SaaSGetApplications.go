package armsaas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/saas/resource-manager/Microsoft.SaaS/preview/2018-03-01-beta/examples/saasV1/SaaSGetApplications.json
func ExampleApplicationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsaas.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationsClient().NewListPager("myResourceGroup", nil)
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
		// page.AppResponseWithContinuation = armsaas.AppResponseWithContinuation{
		// 	Value: []*armsaas.App{
		// 		{
		// 			Name: to.Ptr("myapp"),
		// 			Type: to.Ptr("Microsoft.SaaS/applications"),
		// 			ID: to.Ptr("/subscriptions/bc6c2f82-a39d-41b8-a648-71527498a23e/resourceGroups/myResourceGroup/providers/Microsoft.SaaS/applications/myapp"),
		// 			Location: to.Ptr("location"),
		// 			Properties: &armsaas.AppProperties{
		// 				SaasAppPlan: &armsaas.AppPlan{
		// 					Name: to.Ptr("myPlan1"),
		// 					Product: to.Ptr("myOffer"),
		// 					Publisher: to.Ptr("contoso"),
		// 				},
		// 				Status: to.Ptr(armsaas.SaasAppStatusSubscribed),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myapp"),
		// 			Type: to.Ptr("Microsoft.SaaS/applications"),
		// 			ID: to.Ptr("/subscriptions/bc6c2f82-a39d-41b8-a648-71527498a23e/resourceGroups/myResourceGroup/providers/Microsoft.SaaS/applications/myapp"),
		// 			Location: to.Ptr("location"),
		// 			Properties: &armsaas.AppProperties{
		// 				SaasAppPlan: &armsaas.AppPlan{
		// 					Name: to.Ptr("myPlan2"),
		// 					Product: to.Ptr("myOffer"),
		// 					Publisher: to.Ptr("contoso"),
		// 				},
		// 				Status: to.Ptr(armsaas.SaasAppStatusSubscribed),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 	}},
		// }
	}
}
