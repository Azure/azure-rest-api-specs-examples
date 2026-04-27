package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Connector_List_MaximumSet_Gen.json
func ExampleConnectorClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectorClient().NewListPager("rgconfluent", "ygxwgulsjztjoxuhmegodplubt", "mmxahiyh", "rslbzgqdgsnwzsqhlhethe", &armconfluent.ConnectorClientListOptions{
		PageSize:  to.Ptr[int32](18),
		PageToken: to.Ptr("spklebovnebppxshqcmkyundbw")})
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
		// page = armconfluent.ConnectorClientListResponse{
		// 	ListConnectorsSuccessResponse: armconfluent.ListConnectorsSuccessResponse{
		// 		Value: []*armconfluent.ConnectorResource{
		// 			{
		// 				Properties: &armconfluent.ConnectorResourceProperties{
		// 					ConnectorBasicInfo: &armconfluent.ConnectorInfoBase{
		// 						ConnectorType: to.Ptr(armconfluent.ConnectorTypeSINK),
		// 						ConnectorClass: to.Ptr(armconfluent.ConnectorClassAZUREBLOBSOURCE),
		// 						ConnectorName: to.Ptr("gxad"),
		// 						ConnectorID: to.Ptr("qlrrqyekgitbbes"),
		// 						ConnectorState: to.Ptr(armconfluent.ConnectorStatusPROVISIONING),
		// 					},
		// 					ConnectorServiceTypeInfo: &armconfluent.ConnectorServiceTypeInfoBase{
		// 						ConnectorServiceType: to.Ptr(armconfluent.ConnectorServiceType("ConnectorServiceTypeInfoBase")),
		// 					},
		// 					PartnerConnectorInfo: &armconfluent.PartnerInfoBase{
		// 						PartnerConnectorType: to.Ptr(armconfluent.PartnerConnectorType("PartnerInfoBase")),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
		// 				Name: to.Ptr("fnjrvjq"),
		// 				Type: to.Ptr("xxmdgmmqfhsvkaxpvoermvn"),
		// 				SystemData: &armconfluent.SystemData{
		// 					CreatedBy: to.Ptr("lfskmafvssxoohhokqsa"),
		// 					CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("txubvkbhgirdizxd"),
		// 					LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
