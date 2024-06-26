package armdomainservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/GetOperations.json
func ExampleDomainServiceOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDomainServiceOperationsClient().NewListPager(nil)
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
		// page.OperationEntityListResult = armdomainservices.OperationEntityListResult{
		// 	Value: []*armdomainservices.OperationEntity{
		// 		{
		// 			Name: to.Ptr("Microsoft.AAD/unregister/action"),
		// 			Display: &armdomainservices.OperationDisplayInfo{
		// 				Description: to.Ptr("Unregisters Domain Services"),
		// 				Operation: to.Ptr("Unregister Domain Service"),
		// 				Provider: to.Ptr("Domain Services Resource Provider"),
		// 				Resource: to.Ptr("Domain Service Type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AAD/register/action"),
		// 			Display: &armdomainservices.OperationDisplayInfo{
		// 				Description: to.Ptr("Registers Domain Services"),
		// 				Operation: to.Ptr("Register Domain Service"),
		// 				Provider: to.Ptr("Domain Services Resource Provider"),
		// 				Resource: to.Ptr("Domain Service Type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AAD/domainServices/read"),
		// 			Display: &armdomainservices.OperationDisplayInfo{
		// 				Description: to.Ptr("Reads Domain Services"),
		// 				Operation: to.Ptr("Read Domain Service"),
		// 				Provider: to.Ptr("Domain Services Resource Provider"),
		// 				Resource: to.Ptr("Domain Service Type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AAD/domainServices/write"),
		// 			Display: &armdomainservices.OperationDisplayInfo{
		// 				Description: to.Ptr("Writes Domain Services"),
		// 				Operation: to.Ptr("Write Domain Service"),
		// 				Provider: to.Ptr("Domain Services Resource Provider"),
		// 				Resource: to.Ptr("Domain Service Type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AAD/domainServices/delete"),
		// 			Display: &armdomainservices.OperationDisplayInfo{
		// 				Description: to.Ptr("Deletes Domain Services"),
		// 				Operation: to.Ptr("Delete Domain Service"),
		// 				Provider: to.Ptr("Domain Services Resource Provider"),
		// 				Resource: to.Ptr("Domain Service Type"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
