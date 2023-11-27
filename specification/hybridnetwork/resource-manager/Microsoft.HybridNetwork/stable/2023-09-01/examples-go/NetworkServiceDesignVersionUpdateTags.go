package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkServiceDesignVersionUpdateTags.json
func ExampleNetworkServiceDesignVersionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkServiceDesignVersionsClient().Update(ctx, "rg", "TestPublisher", "TestNetworkServiceDesignGroupName", "1.0.0", armhybridnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkServiceDesignVersion = armhybridnetwork.NetworkServiceDesignVersion{
	// 	Name: to.Ptr("TestVersion"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/publishers/networkServiceDesignGroups/networkServiceDesignVersions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/TestPublisher/networkServiceDesignGroups/TestNetworkServiceDesignGroupName/networkServiceDesignVersions/1.0.0"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armhybridnetwork.NetworkServiceDesignVersionPropertiesFormat{
	// 		ConfigurationGroupSchemaReferences: map[string]*armhybridnetwork.ReferencedResource{
	// 			"MyVM_Configuration": &armhybridnetwork.ReferencedResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/publishers/contosoGroup/networkServiceDesignGroups/NSD_contoso/configurationGroupSchemas/MyVM_Configuration_Schema"),
	// 			},
	// 		},
	// 		ResourceElementTemplates: []armhybridnetwork.ResourceElementTemplateClassification{
	// 			&armhybridnetwork.ArmResourceDefinitionResourceElementTemplateDetails{
	// 				Name: to.Ptr("MyVM"),
	// 				ResourceElementType: to.Ptr(armhybridnetwork.TypeArmResourceDefinition),
	// 				DependsOnProfile: &armhybridnetwork.DependsOnProfile{
	// 					InstallDependsOn: []*string{
	// 					},
	// 				},
	// 				Configuration: &armhybridnetwork.ArmResourceDefinitionResourceElementTemplate{
	// 					ArtifactProfile: &armhybridnetwork.NSDArtifactProfile{
	// 						ArtifactName: to.Ptr("MyVMArmTemplate"),
	// 						ArtifactStoreReference: &armhybridnetwork.ReferencedResource{
	// 							ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/publishers/contosoGroup/ArtifactStore/store1"),
	// 						},
	// 						ArtifactVersion: to.Ptr("1.0.0"),
	// 					},
	// 					ParameterValues: to.Ptr("\"publisherName\": \"{configurationparameters('MyVM_Configuration').publisherName}\",\r\n                \"skuGroupName\": \"{configurationparameters('MyVM_Configuration').skuGroupName}\",\r\n                \"skuVersion\": \"{configurationparameters('MyVM_Configuration').skuVersion}\",\r\n                \"skuOfferingLocation\": \"{configurationparameters('MyVM_Configuration').skuOfferingLocation}\",\r\n                \"nfviType\": \"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviType}\",\r\n                \"nfviId\": \"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviId}\",\r\n                \"allowSoftwareUpdates\": \"{configurationparameters('MyVM_Configuration').allowSoftwareUpdates}\",\r\n                \"virtualNetworkName\": \"{configurationparameters('MyVM_Configuration').vnetName}\",\r\n                \"subnetName\": \"{configurationparameters('MyVM_Configuration').subnetName}\",\r\n                \"subnetAddressPrefix\": \"{configurationparameters('MyVM_Configuration').subnetAddressPrefix}\",\r\n                \"managedResourceGroup\": \"{configurationparameters('SNSSelf').managedResourceGroupName}\"\r\n  "),
	// 					TemplateType: to.Ptr(armhybridnetwork.TemplateTypeArmTemplate),
	// 				},
	// 		}},
	// 		VersionState: to.Ptr(armhybridnetwork.VersionStateActive),
	// 	},
	// }
}
