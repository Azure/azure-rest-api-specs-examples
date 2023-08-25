package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiVersionSets.json
func ExampleAPIVersionSetClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPIVersionSetClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.APIVersionSetClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
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
		// page.APIVersionSetCollection = armapimanagement.APIVersionSetCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.APIVersionSetContract{
		// 		{
		// 			Name: to.Ptr("vs1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/api-version-sets"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/vs1"),
		// 			Properties: &armapimanagement.APIVersionSetContractProperties{
		// 				Description: to.Ptr("Version configuration"),
		// 				DisplayName: to.Ptr("api set 1"),
		// 				VersioningScheme: to.Ptr(armapimanagement.VersioningSchemeSegment),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("vs2"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/api-version-sets"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apiVersionSets/vs2"),
		// 			Properties: &armapimanagement.APIVersionSetContractProperties{
		// 				Description: to.Ptr("Version configuration 2"),
		// 				DisplayName: to.Ptr("api set 2"),
		// 				VersioningScheme: to.Ptr(armapimanagement.VersioningSchemeQuery),
		// 			},
		// 	}},
		// }
	}
}
