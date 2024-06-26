package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkServiceDesignVersionGet.json
func ExampleNetworkServiceDesignVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkServiceDesignVersionsClient().Get(ctx, "rg", "TestPublisher", "TestNetworkServiceDesignGroupName", "1.0.0", nil)
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
	// 	Properties: &armhybridnetwork.NetworkServiceDesignVersionPropertiesFormat{
	// 		ConfigurationGroupSchemaReferences: map[string]*armhybridnetwork.ReferencedResource{
	// 			"MyVM_Configuration": &armhybridnetwork.ReferencedResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/publishers/contosoGroup/networkServiceDesignGroups/NSD_contoso/configurationGroupSchemas/MyVM_Configuration_Schema"),
	// 			},
	// 		},
	// 		NfvisFromSite: map[string]*armhybridnetwork.NfviDetails{
	// 			"nfvi1": &armhybridnetwork.NfviDetails{
	// 				Name: to.Ptr("{configurationparameter('MyVM_Configuration').nfviNameFromSite}"),
	// 				Type: to.Ptr("AzureCore"),
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
	// 					ParameterValues: to.Ptr("{\"publisherName\":\"{configurationparameters('MyVM_Configuration').publisherName}\",\"skuGroupName\":\"{configurationparameters('MyVM_Configuration').skuGroupName}\",\"skuVersion\":\"{configurationparameters('MyVM_Configuration').skuVersion}\",\"skuOfferingLocation\":\"{configurationparameters('MyVM_Configuration').skuOfferingLocation}\",\"nfviType\":\"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviType}\",\"nfviId\":\"{nfvis().nfvisFromSitePerNfviType.AzureCore.nfviAlias1.nfviId}\",\"allowSoftwareUpdates\":\"{configurationparameters('MyVM_Configuration').allowSoftwareUpdates}\",\"virtualNetworkName\":\"{configurationparameters('MyVM_Configuration').vnetName}\",\"subnetName\":\"{configurationparameters('MyVM_Configuration').subnetName}\",\"subnetAddressPrefix\":\"{configurationparameters('MyVM_Configuration').subnetAddressPrefix}\",\"managedResourceGroup\":\"{configurationparameters('SNSSelf').managedResourceGroupName}\",\"adminPassword\":\"{secretparameters('MyVM_Configuration').adminPassword}\"}"),
	// 					TemplateType: to.Ptr(armhybridnetwork.TemplateTypeArmTemplate),
	// 				},
	// 		}},
	// 		VersionState: to.Ptr(armhybridnetwork.VersionStateActive),
	// 	},
	// }
}
