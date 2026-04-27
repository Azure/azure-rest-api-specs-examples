package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Environment_CreateOrUpdate_MaximumSet_Gen.json
func ExampleEnvironmentClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentClient().CreateOrUpdate(ctx, "rgconfluent", "uf", "diycvbfypirqvomdkt", armconfluent.SCEnvironmentRecord{
		Kind: to.Ptr("qhwbkvelujjbojvhrgiikildjdrqox"),
		Properties: &armconfluent.EnvironmentProperties{
			StreamGovernanceConfig: &armconfluent.StreamGovernanceConfig{
				Package: to.Ptr(armconfluent.PackageESSENTIALS),
			},
			Metadata: &armconfluent.SCMetadataEntity{
				Self:             to.Ptr("bnbnbarlsvfifpzcnsnplf"),
				ResourceName:     to.Ptr("ciadqmxlpgllibvkz"),
				CreatedTimestamp: to.Ptr("ouqjivxfggaxzrsmxm"),
				UpdatedTimestamp: to.Ptr("ctrngbppcxdpzmp"),
				DeletedTimestamp: to.Ptr("gn"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconfluent.EnvironmentClientCreateOrUpdateResponse{
	// 	SCEnvironmentRecord: &armconfluent.SCEnvironmentRecord{
	// 		Kind: to.Ptr("qhwbkvelujjbojvhrgiikildjdrqox"),
	// 		Properties: &armconfluent.EnvironmentProperties{
	// 			StreamGovernanceConfig: &armconfluent.StreamGovernanceConfig{
	// 				Package: to.Ptr(armconfluent.PackageESSENTIALS),
	// 			},
	// 			Metadata: &armconfluent.SCMetadataEntity{
	// 				Self: to.Ptr("bnbnbarlsvfifpzcnsnplf"),
	// 				ResourceName: to.Ptr("ciadqmxlpgllibvkz"),
	// 				CreatedTimestamp: to.Ptr("ouqjivxfggaxzrsmxm"),
	// 				UpdatedTimestamp: to.Ptr("ctrngbppcxdpzmp"),
	// 				DeletedTimestamp: to.Ptr("gn"),
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
	// 		Name: to.Ptr("tmbxsqjrvmryomopsflpdvaygfzy"),
	// 		Type: to.Ptr("chfikcsvqvoiedtcgd"),
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
