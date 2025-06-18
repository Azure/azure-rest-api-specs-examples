package armagricultureplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agricultureplatform/armagricultureplatform"
)

// Generated from example definition: 2024-06-01-preview/AgriService_ListAvailableSolutions_MaximumSet_Gen.json
func ExampleAgriServiceClient_ListAvailableSolutions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagricultureplatform.NewClientFactory("83D293F5-DEFD-4D48-B120-1DC713BE338A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgriServiceClient().ListAvailableSolutions(ctx, "rgopenapi", "abc123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armagricultureplatform.AgriServiceClientListAvailableSolutionsResponse{
	// 	AvailableAgriSolutionListResult: &armagricultureplatform.AvailableAgriSolutionListResult{
	// 		Solutions: []*armagricultureplatform.DataManagerForAgricultureSolution{
	// 			{
	// 				PartnerID: to.Ptr("dugq"),
	// 				SolutionID: to.Ptr("sgdbaiygsffxcokkxygtepomgyspz"),
	// 				PartnerTenantID: to.Ptr("nxvc"),
	// 				DataAccessScopes: []*string{
	// 					to.Ptr("ognbthj"),
	// 				},
	// 				MarketPlaceOfferDetails: &armagricultureplatform.MarketPlaceOfferDetails{
	// 					SaasOfferID: to.Ptr("xbzymkxqoggdcjrfyvpqaee"),
	// 					PublisherID: to.Ptr("ihvsmtzqbgwudeicsawqovi"),
	// 				},
	// 				SaasApplicationID: to.Ptr("ypzopzkbzukfxalmeu"),
	// 				AccessAzureDataManagerForAgricultureApplicationID: to.Ptr("khzwsikjlokrhdhotartjeofpiw"),
	// 				AccessAzureDataManagerForAgricultureApplicationName: to.Ptr("ztfnwoksuurzlizk"),
	// 				IsValidateInput: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
