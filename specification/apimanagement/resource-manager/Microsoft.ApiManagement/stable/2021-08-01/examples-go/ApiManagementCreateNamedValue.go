package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateNamedValue.json
func ExampleNamedValueClient_BeginCreateOrUpdate_apiManagementCreateNamedValue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamedValueClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "testprop2", armapimanagement.NamedValueCreateContract{
		Properties: &armapimanagement.NamedValueCreateContractProperties{
			Secret: to.Ptr(false),
			Tags: []*string{
				to.Ptr("foo"),
				to.Ptr("bar")},
			DisplayName: to.Ptr("prop3name"),
			Value:       to.Ptr("propValue"),
		},
	}, &armapimanagement.NamedValueClientBeginCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NamedValueContract = armapimanagement.NamedValueContract{
	// 	Name: to.Ptr("testprop2"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/namedValues"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/namedValues/testprop2"),
	// 	Properties: &armapimanagement.NamedValueContractProperties{
	// 		Secret: to.Ptr(false),
	// 		Tags: []*string{
	// 			to.Ptr("foo"),
	// 			to.Ptr("bar")},
	// 			DisplayName: to.Ptr("prop3name"),
	// 			Value: to.Ptr("propValue"),
	// 		},
	// 	}
}
