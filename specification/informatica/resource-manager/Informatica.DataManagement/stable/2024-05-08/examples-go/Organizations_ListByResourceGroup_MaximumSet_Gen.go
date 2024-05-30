package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_ListByResourceGroup_MaximumSet_Gen.json
func ExampleOrganizationsClient_NewListByResourceGroupPager_organizationsListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationsClient().NewListByResourceGroupPager("rgopenapi", nil)
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
		// page.InformaticaOrganizationResourceListResult = arminformaticadatamgmt.InformaticaOrganizationResourceListResult{
		// 	Value: []*arminformaticadatamgmt.InformaticaOrganizationResource{
		// 		{
		// 			Name: to.Ptr("qmlpllxohjomejbeylyhlqwt"),
		// 			Type: to.Ptr("korjyotq"),
		// 			ID: to.Ptr("cadokiejnrth"),
		// 			SystemData: &arminformaticadatamgmt.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
		// 				CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
		// 				CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
		// 				LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("pamjoudtssthlbhrnfjidr"),
		// 			Tags: map[string]*string{
		// 				"key8430": to.Ptr("cagshqtjlxtqqhdwtchokvxszybp"),
		// 			},
		// 			Properties: &arminformaticadatamgmt.OrganizationProperties{
		// 				CompanyDetails: &arminformaticadatamgmt.CompanyDetails{
		// 					Business: to.Ptr("pucosrtjv"),
		// 					CompanyName: to.Ptr("xszcggknokhw"),
		// 					Country: to.Ptr("gwkcpnwyaqc"),
		// 					Domain: to.Ptr("utcxetzzpmbvwmjrvphqngvp"),
		// 					NumberOfEmployees: to.Ptr[int32](25),
		// 					OfficeAddress: to.Ptr("sbttzwyajgdbsvipuiclbzvkcvwyil"),
		// 				},
		// 				InformaticaProperties: &arminformaticadatamgmt.InformaticaProperties{
		// 					InformaticaRegion: to.Ptr("zfqodqpbeflhedypiijdkc"),
		// 					OrganizationID: to.Ptr("wtdmhlwhkvgqdumaehgfgiqcxgnqpx"),
		// 					OrganizationName: to.Ptr("nomzbvwe"),
		// 					SingleSignOnURL: to.Ptr("espcbzjhtmgwfdkckhqk"),
		// 				},
		// 				LinkOrganization: &arminformaticadatamgmt.LinkOrganization{
		// 					Token: to.Ptr("jjfouhoqpumjvrdsfbimgcy"),
		// 				},
		// 				MarketplaceDetails: &arminformaticadatamgmt.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("ovenlecocg"),
		// 					OfferDetails: &arminformaticadatamgmt.OfferDetails{
		// 						OfferID: to.Ptr("cwswcfwmzhjcoksmueukegwaptvpcmbfyvixfhvgwnjyblqivqdkkwkunkgimiopwwkvgnwclmajhuty"),
		// 						PlanID: to.Ptr("jfnemevyivtlxhectiutdavdgfyidolivuojumdzckp"),
		// 						PlanName: to.Ptr("iaoxgaitteuoqgujkgxbdgryaobtkjjecuvchwutntrvmuorikrbqqegmelenbewhakiysprrnovjixyxrikscaptrbapbdspu"),
		// 						PublisherID: to.Ptr("zajxpfacudwongxjvnnuhhpygmnydchgowjccyuzsjonegmqxcqqpnzafanggowfqdixnnutyfvmvwrkx"),
		// 						TermID: to.Ptr("tcvvsxdjnjlfmjhmvwklptdmxetnzydxyuhfqchoubmtoeqbchnfxoxqzezlgpxdnzyvzgkynjxzzgetkqccxvpzahxattluqdipvbdktqmndfefitzuifqjpschzlbvixnvznkmmgjwvkplfhemnapsewgqxggdzdokryhv"),
		// 						TermUnit: to.Ptr("gjwmgevrblbosuogsvfspsgspetbnxaygkbelvadpgwiywl"),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
		// 				UserDetails: &arminformaticadatamgmt.UserDetails{
		// 					EmailAddress: to.Ptr("7_-46@13D--3.m-4x-.11.c-9-.DHLYFc"),
		// 					FirstName: to.Ptr("appvdclawzfjntdfdftjevlhvzropnxqtnypid"),
		// 					LastName: to.Ptr("nzirbvzmkxtbrlamyatlcszebxgcyncxoascojsmacwvjsjvn"),
		// 					PhoneNumber: to.Ptr("fvcjylxlmhdnshsgywnzlyvshu"),
		// 					Upn: to.Ptr("undljch"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
