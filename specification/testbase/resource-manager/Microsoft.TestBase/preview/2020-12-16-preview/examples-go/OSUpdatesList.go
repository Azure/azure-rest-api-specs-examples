package armtestbase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/testbase/resource-manager/Microsoft.TestBase/preview/2020-12-16-preview/examples/OSUpdatesList.json
func ExampleOSUpdatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtestbase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOSUpdatesClient().NewListPager("contoso-rg1", "contoso-testBaseAccount1", "contoso-package2", armtestbase.OsUpdateTypeSecurityUpdate, nil)
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
		// page.OSUpdateListResult = armtestbase.OSUpdateListResult{
		// 	Value: []*armtestbase.OSUpdateResource{
		// 		{
		// 			Name: to.Ptr("Windows-10-2004-2020-12-B-505"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages/osUpdates"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/osUpdates/Windows-10-2004-2020-12-B-505"),
		// 			Properties: &armtestbase.OSUpdateProperties{
		// 				Type: to.Ptr(armtestbase.TypeSecurityUpdate),
		// 				BuildRevision: to.Ptr("505"),
		// 				OSName: to.Ptr("Windows 10 2004"),
		// 				Release: to.Ptr("2020.12B"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Windows-10-2004-2020-10-B-413"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages/osUpdates"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/osUpdates/Windows-10-2004-2020-10-B-413"),
		// 			Properties: &armtestbase.OSUpdateProperties{
		// 				Type: to.Ptr(armtestbase.TypeSecurityUpdate),
		// 				BuildRevision: to.Ptr("413"),
		// 				OSName: to.Ptr("Windows 10 2004"),
		// 				Release: to.Ptr("2020.10B"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Windows-10-2004-2020-08-B-226"),
		// 			Type: to.Ptr("Microsoft.TestBase/testBaseAccounts/packages/osUpdates"),
		// 			ID: to.Ptr("/subscriptions/476f61a4-952c-422a-b4db-568a828f35df/resourceGroups/contoso-rg1/providers/Microsoft.TestBase/testBaseAccounts/contoso-testBaseAccount1/packages/contoso-package2/osUpdates/Windows-10-2004-2020-08-B-226"),
		// 			Properties: &armtestbase.OSUpdateProperties{
		// 				Type: to.Ptr(armtestbase.TypeSecurityUpdate),
		// 				BuildRevision: to.Ptr("226"),
		// 				OSName: to.Ptr("Windows 10 2004"),
		// 				Release: to.Ptr("2020.08B"),
		// 			},
		// 	}},
		// }
	}
}
