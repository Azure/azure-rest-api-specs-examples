package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Connector_Get_MaximumSet_Gen.json
func ExampleConnectorClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectorClient().Get(ctx, "rgconfluent", "pgwuoi", "rxbrvvdnplvbedrzwbgtwhbdm", "eknmpvbhtvwxdxddkos", "zakwjragxeiur", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconfluent.ConnectorClientGetResponse{
	// 	ConnectorResource: &armconfluent.ConnectorResource{
	// 		Properties: &armconfluent.ConnectorResourceProperties{
	// 			ConnectorBasicInfo: &armconfluent.ConnectorInfoBase{
	// 				ConnectorType: to.Ptr(armconfluent.ConnectorTypeSINK),
	// 				ConnectorClass: to.Ptr(armconfluent.ConnectorClassAZUREBLOBSOURCE),
	// 				ConnectorName: to.Ptr("gxad"),
	// 				ConnectorID: to.Ptr("qlrrqyekgitbbes"),
	// 				ConnectorState: to.Ptr(armconfluent.ConnectorStatusPROVISIONING),
	// 			},
	// 			ConnectorServiceTypeInfo: &armconfluent.ConnectorServiceTypeInfoBase{
	// 				ConnectorServiceType: to.Ptr(armconfluent.ConnectorServiceType("ConnectorServiceTypeInfoBase")),
	// 			},
	// 			PartnerConnectorInfo: &armconfluent.PartnerInfoBase{
	// 				PartnerConnectorType: to.Ptr(armconfluent.PartnerConnectorType("PartnerInfoBase")),
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
	// 		Name: to.Ptr("fnjrvjq"),
	// 		Type: to.Ptr("xxmdgmmqfhsvkaxpvoermvn"),
	// 		SystemData: &armconfluent.SystemData{
	// 			CreatedBy: to.Ptr("lfskmafvssxoohhokqsa"),
	// 			CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("txubvkbhgirdizxd"),
	// 			LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t}()),
	// 		},
	// 	},
	// }
}
