package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Organization_ListEnvironments_MaximumSet_Gen.json
func ExampleOrganizationClient_NewListEnvironmentsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationClient().NewListEnvironmentsPager("rgconfluent", "zgvcszgobzkrvomvhkabzamqincp", &armconfluent.OrganizationClientListEnvironmentsOptions{
		PageSize:  to.Ptr[int32](21),
		PageToken: to.Ptr("e")})
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
		// page = armconfluent.OrganizationClientListEnvironmentsResponse{
		// 	GetEnvironmentsResponse: armconfluent.GetEnvironmentsResponse{
		// 		Value: []*armconfluent.SCEnvironmentRecord{
		// 			{
		// 				Kind: to.Ptr("qhwbkvelujjbojvhrgiikildjdrqox"),
		// 				Properties: &armconfluent.EnvironmentProperties{
		// 					StreamGovernanceConfig: &armconfluent.StreamGovernanceConfig{
		// 						Package: to.Ptr(armconfluent.PackageESSENTIALS),
		// 					},
		// 					Metadata: &armconfluent.SCMetadataEntity{
		// 						Self: to.Ptr("bnbnbarlsvfifpzcnsnplf"),
		// 						ResourceName: to.Ptr("ciadqmxlpgllibvkz"),
		// 						CreatedTimestamp: to.Ptr("ouqjivxfggaxzrsmxm"),
		// 						UpdatedTimestamp: to.Ptr("ctrngbppcxdpzmp"),
		// 						DeletedTimestamp: to.Ptr("gn"),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
		// 				Name: to.Ptr("tmbxsqjrvmryomopsflpdvaygfzy"),
		// 				Type: to.Ptr("chfikcsvqvoiedtcgd"),
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
