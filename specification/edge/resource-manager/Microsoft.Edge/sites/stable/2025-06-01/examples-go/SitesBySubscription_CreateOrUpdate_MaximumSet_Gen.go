package armsitemanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sitemanager/armsitemanager"
)

// Generated from example definition: 2025-06-01/SitesBySubscription_CreateOrUpdate_MaximumSet_Gen.json
func ExampleSitesBySubscriptionClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsitemanager.NewClientFactory("0154f7fe-df09-4981-bf82-7ad5c1f596eb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSitesBySubscriptionClient().BeginCreateOrUpdate(ctx, "string", armsitemanager.Site{
		Properties: &armsitemanager.SiteProperties{
			DisplayName: to.Ptr("string"),
			Labels: map[string]*string{
				"key8188": to.Ptr("mcgnu"),
			},
			Description: to.Ptr("enxcmpvfvadbapo"),
			SiteAddress: &armsitemanager.SiteAddressProperties{
				StreetAddress1:  to.Ptr("fodimymrxbhrfslsmzfhmitn"),
				StreetAddress2:  to.Ptr("widjg"),
				City:            to.Ptr("zkcbzjkatafo"),
				StateOrProvince: to.Ptr("wk"),
				Country:         to.Ptr("xeevcfvimlfzsfuxtyujw"),
				PostalCode:      to.Ptr("qbrhqk"),
			},
		},
	}, nil)
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
	// res = armsitemanager.SitesBySubscriptionClientCreateOrUpdateResponse{
	// 	Site: &armsitemanager.Site{
	// 		Properties: &armsitemanager.SiteProperties{
	// 			DisplayName: to.Ptr("string"),
	// 			Labels: map[string]*string{
	// 				"key8188": to.Ptr("mcgnu"),
	// 			},
	// 			Description: to.Ptr("enxcmpvfvadbapo"),
	// 			SiteAddress: &armsitemanager.SiteAddressProperties{
	// 				StreetAddress1: to.Ptr("fodimymrxbhrfslsmzfhmitn"),
	// 				StreetAddress2: to.Ptr("widjg"),
	// 				City: to.Ptr("zkcbzjkatafo"),
	// 				StateOrProvince: to.Ptr("wk"),
	// 				Country: to.Ptr("xeevcfvimlfzsfuxtyujw"),
	// 				PostalCode: to.Ptr("qbrhqk"),
	// 			},
	// 			ProvisioningState: to.Ptr(armsitemanager.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/SGSites/providers/Microsoft.Edge/Sites/Rome"),
	// 		Name: to.Ptr("string"),
	// 		Type: to.Ptr("string"),
	// 		SystemData: &armsitemanager.SystemData{
	// 			CreatedBy: to.Ptr("julxbiyjzi"),
	// 			CreatedByType: to.Ptr(armsitemanager.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-30T07:53:03.972Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("bceneuzzvzqmiocbrfef"),
	// 			LastModifiedByType: to.Ptr(armsitemanager.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-30T07:53:03.972Z"); return t}()),
	// 		},
	// 	},
	// }
}
