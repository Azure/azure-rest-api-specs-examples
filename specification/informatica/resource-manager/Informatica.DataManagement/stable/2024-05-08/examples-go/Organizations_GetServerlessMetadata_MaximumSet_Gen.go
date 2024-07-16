package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/56537883b7cdb95618c3d1ec1c0ee37b59d88d72/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_GetServerlessMetadata_MaximumSet_Gen.json
func ExampleOrganizationsClient_GetServerlessMetadata_organizationsGetServerlessMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().GetServerlessMetadata(ctx, "rgopenapi", "3_UC", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerlessMetadataResponse = arminformaticadatamgmt.ServerlessMetadataResponse{
	// 	Type: to.Ptr(arminformaticadatamgmt.RuntimeTypeSERVERLESS),
	// 	ServerlessConfigProperties: &arminformaticadatamgmt.ServerlessConfigProperties{
	// 		ApplicationTypes: []*arminformaticadatamgmt.ApplicationTypeMetadata{
	// 			{
	// 				Name: to.Ptr("kdbgedaiazvoakoheguldnmhadcs"),
	// 				Value: to.Ptr("ipbrpyvxilvqfnwktnq"),
	// 		}},
	// 		ComputeUnits: []*arminformaticadatamgmt.ComputeUnitsMetadata{
	// 			{
	// 				Name: to.Ptr("dvvnxwibhl"),
	// 				Value: []*string{
	// 					to.Ptr("msdbcuekw")},
	// 			}},
	// 			ExecutionTimeout: to.Ptr("zmypifivoqmdpzfjhtpqzvlxol"),
	// 			Platform: to.Ptr(arminformaticadatamgmt.PlatformTypeAZURE),
	// 			Regions: []*arminformaticadatamgmt.RegionsMetadata{
	// 				{
	// 					Name: to.Ptr("wwnertzjidsshgoesojbnaopb"),
	// 					ID: to.Ptr("mqkbutwvzbnnisnen"),
	// 			}},
	// 		},
	// 		ServerlessRuntimeConfigProperties: &arminformaticadatamgmt.ServerlessRuntimeConfigProperties{
	// 			CdiConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 			CdieConfigProps: []*arminformaticadatamgmt.CdiConfigProps{
	// 				{
	// 					ApplicationConfigs: []*arminformaticadatamgmt.ApplicationConfigs{
	// 						{
	// 							Name: to.Ptr("upfvjrqcrwwedfujkmsodeinw"),
	// 							Type: to.Ptr("lw"),
	// 							Customized: to.Ptr("j"),
	// 							DefaultValue: to.Ptr("zvgkqwmi"),
	// 							Platform: to.Ptr("dixfyeobngivyvf"),
	// 							Value: to.Ptr("mozgsetpwjmtyl"),
	// 					}},
	// 					EngineName: to.Ptr("hngsdqvtjdhwqlbqfotipaiwjuys"),
	// 					EngineVersion: to.Ptr("zlrlbg"),
	// 			}},
	// 		},
	// 	}
}
