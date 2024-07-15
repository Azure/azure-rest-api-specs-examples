package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/56537883b7cdb95618c3d1ec1c0ee37b59d88d72/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_ListByInformaticaOrganizationResource_MaximumSet_Gen.json
func ExampleServerlessRuntimesClient_NewListByInformaticaOrganizationResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerlessRuntimesClient().NewListByInformaticaOrganizationResourcePager("rgopenapi", "orgName", nil)
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
		// page.InformaticaServerlessRuntimeResourceListResult = arminformaticadatamgmt.InformaticaServerlessRuntimeResourceListResult{
		// 	Value: []*arminformaticadatamgmt.InformaticaServerlessRuntimeResource{
		// 		{
		// 			Name: to.Ptr("byzccgftqjthb"),
		// 			Type: to.Ptr("due"),
		// 			ID: to.Ptr("vcdjzfgqjv"),
		// 			SystemData: &arminformaticadatamgmt.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
		// 				CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
		// 				CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
		// 				LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
		// 			},
		// 			Properties: &arminformaticadatamgmt.InformaticaServerlessRuntimeProperties{
		// 				Description: to.Ptr("mqkaenjmxakvzrwmirelmhgiedto"),
		// 				AdvancedCustomProperties: []*arminformaticadatamgmt.AdvancedCustomProperties{
		// 					{
		// 						Key: to.Ptr("qcmc"),
		// 						Value: to.Ptr("unraxmnohdmvutt"),
		// 				}},
		// 				ApplicationType: to.Ptr(arminformaticadatamgmt.ApplicationTypeCDI),
		// 				ComputeUnits: to.Ptr("bsctukmndvowse"),
		// 				ExecutionTimeout: to.Ptr("ruiougpypny"),
		// 				Platform: to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
		// 				ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
		// 				ServerlessAccountLocation: to.Ptr("bkxdfopapbqucyhduewrubjpaei"),
		// 				ServerlessRuntimeConfig: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
		// 					CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
		// 						{
		// 							ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
		// 								{
		// 									Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
		// 									Type: to.Ptr("lw"),
		// 									Customized: to.Ptr("j"),
		// 									DefaultValue: to.Ptr("zvgkqwmi"),
		// 									Platform: to.Ptr("dixfyeobngivyvf"),
		// 									Value: to.Ptr("mozgsetpwjmtyl"),
		// 							}},
		// 							EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
		// 							EngineVersion: to.Ptr("zlrlbg"),
		// 					}},
		// 					CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
		// 						{
		// 							ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
		// 								{
		// 									Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
		// 									Type: to.Ptr("lw"),
		// 									Customized: to.Ptr("j"),
		// 									DefaultValue: to.Ptr("zvgkqwmi"),
		// 									Platform: to.Ptr("dixfyeobngivyvf"),
		// 									Value: to.Ptr("mozgsetpwjmtyl"),
		// 							}},
		// 							EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
		// 							EngineVersion: to.Ptr("zlrlbg"),
		// 					}},
		// 				},
		// 				ServerlessRuntimeNetworkProfile: &arminformaticadatamgmt.ServerlessRuntimeNetworkProfile{
		// 					NetworkInterfaceConfiguration: &arminformaticadatamgmt.NetworkInterfaceConfiguration{
		// 						SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/subnet1"),
		// 						VnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1"),
		// 					},
		// 				},
		// 				ServerlessRuntimeTags: []*arminformaticadatamgmt.ServerlessRuntimeTag{
		// 					{
		// 						Name: to.Ptr("korveuycuwhs"),
		// 						Value: to.Ptr("uyiuegxnkgp"),
		// 				}},
		// 				ServerlessRuntimeUserContextProperties: &arminformaticadatamgmt.ServerlessRuntimeUserContextProperties{
		// 					UserContextToken: to.Ptr("oludf"),
		// 				},
		// 				SupplementaryFileLocation: to.Ptr("zmlqtkncwgqhhupsnqluumz"),
		// 			},
		// 	}},
		// }
	}
}
