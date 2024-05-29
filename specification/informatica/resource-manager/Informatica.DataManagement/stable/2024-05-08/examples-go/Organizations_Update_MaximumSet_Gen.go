package arminformaticadatamgmt_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/informaticadatamgmt/arminformaticadatamgmt"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a07cb79078c828c5404a5154fea6c60d6e43256e/specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_Update_MaximumSet_Gen.json
func ExampleOrganizationsClient_Update_organizationsUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := arminformaticadatamgmt.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().Update(ctx, "rgopenapi", "_-", arminformaticadatamgmt.InformaticaOrganizationResourceUpdate{
		Properties: &arminformaticadatamgmt.OrganizationPropertiesCustomUpdate{
			CompanyDetails: &arminformaticadatamgmt.CompanyDetailsUpdate{
				Business:          to.Ptr("mwqblnruflwpolgbxpqbqneve"),
				CompanyName:       to.Ptr("xkrvbozrjcvappqeeyt"),
				Country:           to.Ptr("rvlzppgvopcw"),
				Domain:            to.Ptr("dponvwnrdrnzahcurqssesukbsokdd"),
				NumberOfEmployees: to.Ptr[int32](22),
				OfficeAddress:     to.Ptr("sfcx"),
			},
			ExistingResourceID: to.Ptr("uvwlcphdfkqnhrtddpsiacbowcxxo"),
			MarketplaceDetails: &arminformaticadatamgmt.MarketplaceDetailsUpdate{
				MarketplaceSubscriptionID: to.Ptr("szhyxzgjtssjmlguivepc"),
				OfferDetails: &arminformaticadatamgmt.OfferDetailsUpdate{
					OfferID:     to.Ptr("idaxbflabvjsippplyenvrpgeydsjxcmyubgukffkcdvlvrtwpdhnvdblxjsldiuswrchsibk"),
					PlanID:      to.Ptr("giihvvnwdwzkfqrhkpqzbgfotzyixnsvmxzauseebillhslauglzfxzvzvts"),
					PlanName:    to.Ptr("tfqjenotaewzdeerliteqxdawuqxhwdzbtiiimsaedrlsnbdoonnloakjtvnwhhrcyxxsgoachguthqvlahpjyofpoqpfacfmiaauawazkmxkjgvktbptojknzojtjrfzvbbjjkvstabqyaczxinijhoxrjukftsagpwgsvpmczopztmplipyufhuaumfx"),
					PublisherID: to.Ptr("ktzfghsyjqbsswhltoaemgotmnorhdogvkaxplutbjjqzuepxizliynyakersobagvpwvpzwjtjjxigsqgcyqaahaxdijghnexliofhfjlqzjmmbvrhcvjxdodnexxizbgfhjopbwzjojxsluasnwwsgcajefglbcvzpaeblanhmurcculndtfwnfjyxol"),
					TermID:      to.Ptr("eolmwogtgpdncqoigqcdomupwummaicwvdxgbskpdsmjizdfbdgbxbuekcpwmenqzbhqxpdnjtup"),
					TermUnit:    to.Ptr("nykqoplazujcwmfldntifjqrnx"),
				},
			},
			UserDetails: &arminformaticadatamgmt.UserDetailsUpdate{
				EmailAddress: to.Ptr("7_-46@13D--3.m-4x-.11.c-9-.DHLYFc"),
				FirstName:    to.Ptr("qguqrmanyupoi"),
				LastName:     to.Ptr("ugzg"),
				PhoneNumber:  to.Ptr("uxa"),
				Upn:          to.Ptr("viwjrkn"),
			},
		},
		Tags: map[string]*string{
			"key1918": to.Ptr("fbjvtuvzsghpl"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformaticaOrganizationResource = arminformaticadatamgmt.InformaticaOrganizationResource{
	// 	Name: to.Ptr("qmlpllxohjomejbeylyhlqwt"),
	// 	Type: to.Ptr("korjyotq"),
	// 	ID: to.Ptr("cadokiejnrth"),
	// 	SystemData: &arminformaticadatamgmt.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		CreatedBy: to.Ptr("kocqbxulqrggzbfrifpvy"),
	// 		CreatedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-05T15:45:15.582Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("lzpllqnildoamkmgf"),
	// 		LastModifiedByType: to.Ptr(arminformaticadatamgmt.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("pamjoudtssthlbhrnfjidr"),
	// 	Tags: map[string]*string{
	// 		"key8430": to.Ptr("cagshqtjlxtqqhdwtchokvxszybp"),
	// 	},
	// 	Properties: &arminformaticadatamgmt.OrganizationProperties{
	// 		CompanyDetails: &arminformaticadatamgmt.CompanyDetails{
	// 			Business: to.Ptr("mwqblnruflwpolgbxpqbqneve"),
	// 			CompanyName: to.Ptr("xkrvbozrjcvappqeeyt"),
	// 			Country: to.Ptr("rvlzppgvopcw"),
	// 			Domain: to.Ptr("dponvwnrdrnzahcurqssesukbsokdd"),
	// 			NumberOfEmployees: to.Ptr[int32](22),
	// 			OfficeAddress: to.Ptr("sfcx"),
	// 		},
	// 		InformaticaProperties: &arminformaticadatamgmt.InformaticaProperties{
	// 			InformaticaRegion: to.Ptr("zfqodqpbeflhedypiijdkc"),
	// 			OrganizationID: to.Ptr("wtdmhlwhkvgqdumaehgfgiqcxgnqpx"),
	// 			OrganizationName: to.Ptr("nomzbvwe"),
	// 			SingleSignOnURL: to.Ptr("espcbzjhtmgwfdkckhqk"),
	// 		},
	// 		LinkOrganization: &arminformaticadatamgmt.LinkOrganization{
	// 			Token: to.Ptr("jjfouhoqpumjvrdsfbimgcy"),
	// 		},
	// 		MarketplaceDetails: &arminformaticadatamgmt.MarketplaceDetails{
	// 			MarketplaceSubscriptionID: to.Ptr("szhyxzgjtssjmlguivepc"),
	// 			OfferDetails: &arminformaticadatamgmt.OfferDetails{
	// 				OfferID: to.Ptr("idaxbflabvjsippplyenvrpgeydsjxcmyubgukffkcdvlvrtwpdhnvdblxjsldiuswrchsibk"),
	// 				PlanID: to.Ptr("giihvvnwdwzkfqrhkpqzbgfotzyixnsvmxzauseebillhslauglzfxzvzvts"),
	// 				PlanName: to.Ptr("tfqjenotaewzdeerliteqxdawuqxhwdzbtiiimsaedrlsnbdoonnloakjtvnwhhrcyxxsgoachguthqvlahpjyofpoqpfacfmiaauawazkmxkjgvktbptojknzojtjrfzvbbjjkvstabqyaczxinijhoxrjukftsagpwgsvpmczopztmplipyufhuaumfx"),
	// 				PublisherID: to.Ptr("ktzfghsyjqbsswhltoaemgotmnorhdogvkaxplutbjjqzuepxizliynyakersobagvpwvpzwjtjjxigsqgcyqaahaxdijghnexliofhfjlqzjmmbvrhcvjxdodnexxizbgfhjopbwzjojxsluasnwwsgcajefglbcvzpaeblanhmurcculndtfwnfjyxol"),
	// 				TermID: to.Ptr("eolmwogtgpdncqoigqcdomupwummaicwvdxgbskpdsmjizdfbdgbxbuekcpwmenqzbhqxpdnjtup"),
	// 				TermUnit: to.Ptr("nykqoplazujcwmfldntifjqrnx"),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(arminformaticadatamgmt.ProvisioningStateSucceeded),
	// 		UserDetails: &arminformaticadatamgmt.UserDetails{
	// 			EmailAddress: to.Ptr("7_-46@13D--3.m-4x-.11.c-9-.DHLYFc"),
	// 			FirstName: to.Ptr("qguqrmanyupoi"),
	// 			LastName: to.Ptr("ugzg"),
	// 			PhoneNumber: to.Ptr("uxa"),
	// 			Upn: to.Ptr("viwjrkn"),
	// 		},
	// 	},
	// }
}
