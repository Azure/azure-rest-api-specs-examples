package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Organization_ListSchemaRegistryClusters_MaximumSet_Gen.json
func ExampleOrganizationClient_NewListSchemaRegistryClustersPager_organizationListSchemaRegistryClustersMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationClient().NewListSchemaRegistryClustersPager("rgconfluent", "vkzifcygqhoewuixdmmg", "psxriyxxbjnctgeohah", &armconfluent.OrganizationClientListSchemaRegistryClustersOptions{
		PageSize:  to.Ptr[int32](3),
		PageToken: to.Ptr("npqeazvityguunrpgbumrqivvq")})
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
		// page = armconfluent.OrganizationClientListSchemaRegistryClustersResponse{
		// 	ListSchemaRegistryClustersResponse: armconfluent.ListSchemaRegistryClustersResponse{
		// 		Value: []*armconfluent.SchemaRegistryClusterRecord{
		// 			{
		// 				Kind: to.Ptr("vlwwdyfvefryzszqpimeyghwbg"),
		// 				ID: to.Ptr("gpfxwmqatynchrxjdgvr"),
		// 				Properties: &armconfluent.SchemaRegistryClusterProperties{
		// 					Metadata: &armconfluent.SCMetadataEntity{
		// 						Self: to.Ptr("bnbnbarlsvfifpzcnsnplf"),
		// 						ResourceName: to.Ptr("ciadqmxlpgllibvkz"),
		// 						CreatedTimestamp: to.Ptr("ouqjivxfggaxzrsmxm"),
		// 						UpdatedTimestamp: to.Ptr("ctrngbppcxdpzmp"),
		// 						DeletedTimestamp: to.Ptr("gn"),
		// 					},
		// 					Spec: &armconfluent.SchemaRegistryClusterSpecEntity{
		// 						Name: to.Ptr("ofhyvugq"),
		// 						HTTPEndpoint: to.Ptr("jsszdvdfopfoaixx"),
		// 						Package: to.Ptr("yruygcrfip"),
		// 						Region: &armconfluent.SchemaRegistryClusterEnvironmentRegionEntity{
		// 							ID: to.Ptr("k"),
		// 							Related: to.Ptr("inuetawa"),
		// 							ResourceName: to.Ptr("hmjptvxt"),
		// 						},
		// 						Environment: &armconfluent.SchemaRegistryClusterEnvironmentRegionEntity{
		// 							ID: to.Ptr("k"),
		// 							Related: to.Ptr("inuetawa"),
		// 							ResourceName: to.Ptr("hmjptvxt"),
		// 						},
		// 						Cloud: to.Ptr("jcangjhteaadfyipxxso"),
		// 					},
		// 					Status: &armconfluent.SchemaRegistryClusterStatusEntity{
		// 						Phase: to.Ptr("nipb"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
