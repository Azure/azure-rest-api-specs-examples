package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/AvailableOSsList.json
func ExampleAvailableOSClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableOSClient().NewListPager("contoso-rg", "contoso-testBaseAccount", armtestbase.OsUpdateTypeSecurityUpdate, nil)
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
		// page.AvailableOSListResult = armtestbase.AvailableOSListResult{
		// 	Value: []*armtestbase.AvailableOSResource{
		// 		{
		// 			Name: to.Ptr("Windows-10-2004"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/availableOSs"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/availableOSs/Windows-10-2004"),
		// 			Properties: &armtestbase.AvailableOSProperties{
		// 				InsiderChannel: to.Ptr("Channel 1"),
		// 				OSID: to.Ptr("Windows 10 2004"),
		// 				OSName: to.Ptr("Windows 10 2004"),
		// 				OSPlatform: to.Ptr("Client"),
		// 				OSUpdateType: to.Ptr("Feature update"),
		// 				OSVersion: to.Ptr("Windows 10"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Windows-10-1909"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/availableOSs"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/availableOSs/Windows-10-1909"),
		// 			Properties: &armtestbase.AvailableOSProperties{
		// 				OSID: to.Ptr("Windows 10 1909"),
		// 				OSName: to.Ptr("Windows 10 1909"),
		// 				OSPlatform: to.Ptr("Client"),
		// 				OSUpdateType: to.Ptr("Security update"),
		// 				OSVersion: to.Ptr("Windows 10"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Windows-10-1903"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/availableOSs"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/availableOSs/Windows-10-1903"),
		// 			Properties: &armtestbase.AvailableOSProperties{
		// 				OSID: to.Ptr("Windows 10 1903"),
		// 				OSName: to.Ptr("Windows 10 1903"),
		// 				OSPlatform: to.Ptr("Client"),
		// 				OSUpdateType: to.Ptr("Security update"),
		// 				OSVersion: to.Ptr("Windows 10"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Windows-10-1809"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/availableOSs"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/availableOSs/Windows-10-1809"),
		// 			Properties: &armtestbase.AvailableOSProperties{
		// 				InsiderChannel: to.Ptr("Channel 3"),
		// 				OSID: to.Ptr("Windows 10 1809"),
		// 				OSName: to.Ptr("Windows 10 1809"),
		// 				OSPlatform: to.Ptr("Client"),
		// 				OSUpdateType: to.Ptr("Feature update"),
		// 				OSVersion: to.Ptr("Windows 10"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Windows-Server-2016"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/availableOSs"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount/availableOSs/Windows-Server-2016"),
		// 			Properties: &armtestbase.AvailableOSProperties{
		// 				OSID: to.Ptr("Windows Server 2016"),
		// 				OSName: to.Ptr("Windows Server 2016"),
		// 				OSPlatform: to.Ptr("Server"),
		// 				OSUpdateType: to.Ptr("Security update"),
		// 				OSVersion: to.Ptr("Windows Server 2016"),
		// 			},
		// 	}},
		// }
	}
}
