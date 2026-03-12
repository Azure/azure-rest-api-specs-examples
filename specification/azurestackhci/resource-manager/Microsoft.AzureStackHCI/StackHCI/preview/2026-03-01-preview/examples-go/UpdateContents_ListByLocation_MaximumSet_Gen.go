package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/UpdateContents_ListByLocation_MaximumSet_Gen.json
func ExampleUpdateContentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("2886575D-173A-44A0-80E2-7DBA57F18B46", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUpdateContentsClient().NewListPager("westus2", nil)
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
		// page = armazurestackhci.UpdateContentsClientListResponse{
		// 	UpdateContentListResult: armazurestackhci.UpdateContentListResult{
		// 		Value: []*armazurestackhci.UpdateContent{
		// 			{
		// 				Properties: &armazurestackhci.UpdateContentProperties{
		// 					UpdatePayloads: []*armazurestackhci.ContentPayload{
		// 						{
		// 							URL: to.Ptr("https://microsoft.com/a"),
		// 							Hash: to.Ptr("F51FFEEB50A5C51B930F107132DDFA389DC8983313438764204A6F70F4BEFC60"),
		// 							HashAlgorithm: to.Ptr("SHA256"),
		// 							Identifier: to.Ptr("SolutionUpdate_zip"),
		// 							PackageSizeInBytes: to.Ptr("9367"),
		// 							Group: to.Ptr("Update"),
		// 							FileName: to.Ptr("SolutionUpdate.12.2510.1002.84.zip"),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/2886575D-173A-44A0-80E2-7DBA57F18B46/providers/Microsoft.AzureStackHCI/locations/westus2/updateContents/12.2510.0.1"),
		// 				Name: to.Ptr("12.2510.0.1"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/UpdateContents"),
		// 			},
		// 			{
		// 				Properties: &armazurestackhci.UpdateContentProperties{
		// 					UpdatePayloads: []*armazurestackhci.ContentPayload{
		// 						{
		// 							URL: to.Ptr("https://microsoft.com/b"),
		// 							Hash: to.Ptr("A62CCFED60B6C61C950F208243EEGD490ED9094424549875315B7G81G5CFGD71"),
		// 							HashAlgorithm: to.Ptr("SHA256"),
		// 							Identifier: to.Ptr("SolutionUpdate_zip"),
		// 							PackageSizeInBytes: to.Ptr("10456"),
		// 							Group: to.Ptr("Update"),
		// 							FileName: to.Ptr("SolutionUpdate.12.2511.1003.85.zip"),
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/2886575D-173A-44A0-80E2-7DBA57F18B46/providers/Microsoft.AzureStackHCI/locations/westus2/updateContents/12.2511.0.1"),
		// 				Name: to.Ptr("12.2511.0.1"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/UpdateContents"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/2886575D-173A-44A0-80E2-7DBA57F18B46/providers/Microsoft.AzureStackHCI/locations/westus2/updateContents?api-version=2025-12-01-preview&$skiptoken=X'12345'"),
		// 	},
		// }
	}
}
