package armsitemanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sitemanager/armsitemanager"
)

// Generated from example definition: 2025-06-01/SitesByServiceGroup_ListByServiceGroup_MaximumSet_Gen.json
func ExampleSitesByServiceGroupClient_NewListByServiceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsitemanager.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSitesByServiceGroupClient().NewListByServiceGroupPager("string", nil)
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
		// page = armsitemanager.SitesByServiceGroupClientListByServiceGroupResponse{
		// 	SiteListResult: armsitemanager.SiteListResult{
		// 		Value: []*armsitemanager.Site{
		// 			{
		// 				Properties: &armsitemanager.SiteProperties{
		// 					DisplayName: to.Ptr("string"),
		// 					ProvisioningState: to.Ptr(armsitemanager.ResourceProvisioningStateSucceeded),
		// 					Description: to.Ptr("mazbpkzbkvvntk"),
		// 					SiteAddress: &armsitemanager.SiteAddressProperties{
		// 						StreetAddress1: to.Ptr("fodimymrxbhrfslsmzfhmitn"),
		// 						StreetAddress2: to.Ptr("widjg"),
		// 						City: to.Ptr("zkcbzjkatafo"),
		// 						StateOrProvince: to.Ptr("wk"),
		// 						Country: to.Ptr("xeevcfvimlfzsfuxtyujw"),
		// 						PostalCode: to.Ptr("qbrhqk"),
		// 					},
		// 					Labels: map[string]*string{
		// 						"key8188": to.Ptr("mcgnu"),
		// 					},
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/SGSites/providers/Microsoft.Edge/Sites/Rome"),
		// 				Name: to.Ptr("string"),
		// 				Type: to.Ptr("string"),
		// 				SystemData: &armsitemanager.SystemData{
		// 					CreatedBy: to.Ptr("julxbiyjzi"),
		// 					CreatedByType: to.Ptr(armsitemanager.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-30T07:53:03.972Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("bceneuzzvzqmiocbrfef"),
		// 					LastModifiedByType: to.Ptr(armsitemanager.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-30T07:53:03.972Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
