package armneonpostgres_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
)

// Generated from example definition: 2025-03-01/Organizations_Get_MaximumSet_Gen.json
func ExampleOrganizationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("9B8E3300-C5FA-442B-A259-3F6F614D5BD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().Get(ctx, "rgneon", "test-org", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.OrganizationsClientGetResponse{
	// 	OrganizationResource: &armneonpostgres.OrganizationResource{
	// 		Properties: &armneonpostgres.OrganizationProperties{
	// 			MarketplaceDetails: &armneonpostgres.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("xfahbbbzwlcwhhjbxarnwfcy"),
	// 				SubscriptionStatus: to.Ptr(armneonpostgres.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armneonpostgres.OfferDetails{
	// 					PublisherID: to.Ptr("eibghzuyqsksouwlgqphhmuxeqeigf"),
	// 					OfferID: to.Ptr("qscggwfdnippiwrrnmuscg"),
	// 					PlanID: to.Ptr("sveqoxtdwxutxmtniuufyrdu"),
	// 					PlanName: to.Ptr("t"),
	// 					TermUnit: to.Ptr("jnxhyql"),
	// 					TermID: to.Ptr("uptombvymytfonj"),
	// 				},
	// 			},
	// 			UserDetails: &armneonpostgres.UserDetails{
	// 				FirstName: to.Ptr("zhelh"),
	// 				LastName: to.Ptr("zbdhouyeozylnerrc"),
	// 				EmailAddress: to.Ptr("test@contoso.com"),
	// 				Upn: to.Ptr("mixcikvxlnhkfugetqlngz"),
	// 				PhoneNumber: to.Ptr("zmejenytglrmjnt"),
	// 			},
	// 			CompanyDetails: &armneonpostgres.CompanyDetails{
	// 				CompanyName: to.Ptr("xtul"),
	// 				Country: to.Ptr("ycmyjdcpyjieemfrthfyxdlvn"),
	// 				OfficeAddress: to.Ptr("icirtoqmmozijk"),
	// 				BusinessPhone: to.Ptr("hucxvzcvpaupqjkgb"),
	// 				Domain: to.Ptr("snoshqumfsthyofpnrsgyjhszvgtj"),
	// 				NumberOfEmployees: to.Ptr[int64](12),
	// 			},
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			PartnerOrganizationProperties: &armneonpostgres.PartnerOrganizationProperties{
	// 				OrganizationID: to.Ptr("hzejhmftwsruhwspvtwoy"),
	// 				OrganizationName: to.Ptr("entity-name"),
	// 				SingleSignOnProperties: &armneonpostgres.SingleSignOnProperties{
	// 					SingleSignOnState: to.Ptr(armneonpostgres.SingleSignOnStatesInitial),
	// 					EnterpriseAppID: to.Ptr("urtjzjfr"),
	// 					SingleSignOnURL: to.Ptr("gcmlwvtxcsjozitm"),
	// 					AADDomains: []*string{
	// 						to.Ptr("mdzbelaiphukhe"),
	// 					},
	// 				},
	// 			},
	// 			ProjectProperties: &armneonpostgres.ProjectProperties{
	// 				EntityID: to.Ptr("entity-id"),
	// 				EntityName: to.Ptr("entity-name"),
	// 				CreatedAt: to.Ptr("eazudrgcnzbydedhwcmgwoauc"),
	// 				ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 				Attributes: []*armneonpostgres.Attributes{
	// 					{
	// 						Name: to.Ptr("trhvzyvaqy"),
	// 						Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 					},
	// 				},
	// 				RegionID: to.Ptr("tlcltldfrnxh"),
	// 				Storage: to.Ptr[int64](7),
	// 				PgVersion: to.Ptr[int32](10),
	// 				HistoryRetention: to.Ptr[int32](7),
	// 				DefaultEndpointSettings: &armneonpostgres.DefaultEndpointSettings{
	// 					AutoscalingLimitMinCu: to.Ptr[float32](26),
	// 					AutoscalingLimitMaxCu: to.Ptr[float32](20),
	// 				},
	// 				Branch: &armneonpostgres.BranchProperties{
	// 					EntityID: to.Ptr("entity-id"),
	// 					EntityName: to.Ptr("entity-name"),
	// 					CreatedAt: to.Ptr("dzbqaiixq"),
	// 					ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 					Attributes: []*armneonpostgres.Attributes{
	// 						{
	// 							Name: to.Ptr("trhvzyvaqy"),
	// 							Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 						},
	// 					},
	// 					ProjectID: to.Ptr("oik"),
	// 					ParentID: to.Ptr("entity-id"),
	// 					RoleName: to.Ptr("qrrairsupyosxnqotdwhbpc"),
	// 					DatabaseName: to.Ptr("duhxebzhd"),
	// 					Roles: []*armneonpostgres.NeonRoleProperties{
	// 						{
	// 							EntityID: to.Ptr("entity-id"),
	// 							EntityName: to.Ptr("entity-name"),
	// 							CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 							ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 							Attributes: []*armneonpostgres.Attributes{
	// 								{
	// 									Name: to.Ptr("trhvzyvaqy"),
	// 									Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 								},
	// 							},
	// 							BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 							Permissions: []*string{
	// 								to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 							},
	// 							IsSuperUser: to.Ptr(true),
	// 						},
	// 					},
	// 					Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 						{
	// 							EntityID: to.Ptr("entity-id"),
	// 							EntityName: to.Ptr("entity-name"),
	// 							CreatedAt: to.Ptr("wgdmylla"),
	// 							ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 							Attributes: []*armneonpostgres.Attributes{
	// 								{
	// 									Name: to.Ptr("trhvzyvaqy"),
	// 									Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 								},
	// 							},
	// 							BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 							OwnerName: to.Ptr("odmbeg"),
	// 						},
	// 					},
	// 					Endpoints: []*armneonpostgres.EndpointProperties{
	// 						{
	// 							EntityID: to.Ptr("entity-id"),
	// 							EntityName: to.Ptr("entity-name"),
	// 							CreatedAt: to.Ptr("vhcilurdd"),
	// 							ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 							Attributes: []*armneonpostgres.Attributes{
	// 								{
	// 									Name: to.Ptr("trhvzyvaqy"),
	// 									Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 								},
	// 							},
	// 							ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 							BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 							EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 						},
	// 					},
	// 				},
	// 				Roles: []*armneonpostgres.NeonRoleProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("sqpvswctybrhimiwidhnnlxclfry"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("wxbojkmdgaggkfiwqfakdkbyztm"),
	// 						Permissions: []*string{
	// 							to.Ptr("myucqecpjriewzohxvadgkhiudnyx"),
	// 						},
	// 						IsSuperUser: to.Ptr(true),
	// 					},
	// 				},
	// 				Databases: []*armneonpostgres.NeonDatabaseProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("wgdmylla"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						BranchID: to.Ptr("orfdwdmzvfvlnrgussvcvoek"),
	// 						OwnerName: to.Ptr("odmbeg"),
	// 					},
	// 				},
	// 				Endpoints: []*armneonpostgres.EndpointProperties{
	// 					{
	// 						EntityID: to.Ptr("entity-id"),
	// 						EntityName: to.Ptr("entity-name"),
	// 						CreatedAt: to.Ptr("vhcilurdd"),
	// 						ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 						Attributes: []*armneonpostgres.Attributes{
	// 							{
	// 								Name: to.Ptr("trhvzyvaqy"),
	// 								Value: to.Ptr("evpkgsskyavybxwwssm"),
	// 							},
	// 						},
	// 						ProjectID: to.Ptr("rtvdeeflqzlrpfzhjqhcsfbldw"),
	// 						BranchID: to.Ptr("rzsyrhpfbydxtfkpaa"),
	// 						EndpointType: to.Ptr(armneonpostgres.EndpointTypeReadOnly),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key8832": to.Ptr("rvukepuxkykdtqjtwk"),
	// 		},
	// 		Location: to.Ptr("kcdph"),
	// 		ID: to.Ptr("/subscriptions/9B8E3300-C5FA-442B-A259-3F6F614D5BD4/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/test-org"),
	// 		Name: to.Ptr("aocvhndykwhgolfixbqhwtmhiriu"),
	// 		Type: to.Ptr("bvtwhithilvtizpde"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("hnyidmqyvvtsddrwkmrqlwtlew"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("szuncyyauzxhpzlbcvjkeamp"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-24T04:03:54.769Z"); return t}()),
	// 		},
	// 	},
	// }
}
