package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_GetAllServerlessRuntimes_MaximumSet_Gen.json
func ExampleOrganizationsClient_GetAllServerlessRuntimes_organizationsGetAllServerlessRuntimes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().GetAllServerlessRuntimes(ctx, "rgopenapi", "t", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformaticaServerlessRuntimeResourceList = arminformaticadatamgmt.InformaticaServerlessRuntimeResourceList{
	// 	InformaticaRuntimeResources: []*arminformaticadatamgmt.InfaRuntimeResourceFetchMetaData{
	// 		{
	// 			Name: to.Ptr("wwhnptecwbip"),
	// 			Type: to.Ptr(arminformaticadatamgmt.RuntimeTypeSERVERLESS),
	// 			Description: to.Ptr("kozzyfee"),
	// 			CreatedBy: to.Ptr("sxizcppjasypefvhacvzfhzn"),
	// 			CreatedTime: to.Ptr("wfbezwbdyszdvwlzvthrxi"),
	// 			ID: to.Ptr("hgnqusidmutarhmwnswhalpjzievdm"),
	// 			ServerlessConfigProperties: &arminformaticadatamgmt.InfaServerlessFetchConfigProperties{
	// 				AdvancedCustomProperties: to.Ptr("pvusvtqtkhozfuqtokw"),
	// 				ApplicationType: to.Ptr("qnrsngoycualsvmmsu"),
	// 				ComputeUnits: to.Ptr("ao"),
	// 				ExecutionTimeout: to.Ptr("pbnxiricrjifreoiazmzeqiwu"),
	// 				Platform: to.Ptr("p"),
	// 				Region: to.Ptr("lrudpuvzcuh"),
	// 				ResourceGroupName: to.Ptr("ly"),
	// 				ServerlessArmResourceID: to.Ptr("kmviaxayvasaceptsuxkzhqjzjur"),
	// 				Subnet: to.Ptr("mmhpqazmxniguwzmqiu"),
	// 				SubscriptionID: to.Ptr("zl"),
	// 				SupplementaryFileLocation: to.Ptr("wylmhc"),
	// 				Tags: to.Ptr("tevuusglxuivd"),
	// 				TenantID: to.Ptr("elbephnucclhjihcj"),
	// 				Vnet: to.Ptr("vvs"),
	// 			},
	// 			Status: to.Ptr("exvtblad"),
	// 			StatusLocalized: to.Ptr("otg"),
	// 			StatusMessage: to.Ptr("bnosdq"),
	// 			UpdatedBy: to.Ptr("uhsmrmagqrtfd"),
	// 			UpdatedTime: to.Ptr("mjiscnhzzutsrsffhdqsl"),
	// 	}},
	// }
}
