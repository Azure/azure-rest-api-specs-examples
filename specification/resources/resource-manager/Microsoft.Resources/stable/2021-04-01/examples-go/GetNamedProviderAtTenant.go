package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetNamedProviderAtTenant.json
func ExampleProvidersClient_GetAtTenantScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().GetAtTenantScope(ctx, "Microsoft.Storage", &armresources.ProvidersClientGetAtTenantScopeOptions{Expand: to.Ptr("resourceTypes/aliases")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Provider = armresources.Provider{
	// 	Namespace: to.Ptr("Microsoft.Storage"),
	// 	ResourceTypes: []*armresources.ProviderResourceType{
	// 		{
	// 			Aliases: []*armresources.Alias{
	// 				{
	// 					Name: to.Ptr("Microsoft.Storage/storageAccounts/accountType"),
	// 					DefaultPath: to.Ptr("sku.name"),
	// 					Paths: []*armresources.AliasPath{
	// 						{
	// 							Path: to.Ptr("properties.accountType"),
	// 							APIVersions: []*string{
	// 								to.Ptr("2015-06-15"),
	// 								to.Ptr("2015-05-01-preview")},
	// 							},
	// 							{
	// 								Path: to.Ptr("sku.name"),
	// 								APIVersions: []*string{
	// 									to.Ptr("2018-11-01"),
	// 									to.Ptr("2018-11-09"),
	// 									to.Ptr("2018-07-01"),
	// 									to.Ptr("2018-03-01-Preview"),
	// 									to.Ptr("2018-02-01"),
	// 									to.Ptr("2017-10-01"),
	// 									to.Ptr("2017-06-01"),
	// 									to.Ptr("2016-12-01"),
	// 									to.Ptr("2016-05-01"),
	// 									to.Ptr("2016-01-01")},
	// 							}},
	// 						},
	// 						{
	// 							Name: to.Ptr("Microsoft.Storage/storageAccounts/sku.name"),
	// 							DefaultPath: to.Ptr("sku.name"),
	// 							Paths: []*armresources.AliasPath{
	// 								{
	// 									Path: to.Ptr("properties.accountType"),
	// 									APIVersions: []*string{
	// 										to.Ptr("2015-06-15"),
	// 										to.Ptr("2015-05-01-preview")},
	// 									},
	// 									{
	// 										Path: to.Ptr("sku.name"),
	// 										APIVersions: []*string{
	// 											to.Ptr("2018-11-01"),
	// 											to.Ptr("2018-11-09"),
	// 											to.Ptr("2018-07-01"),
	// 											to.Ptr("2018-03-01-Preview"),
	// 											to.Ptr("2018-02-01"),
	// 											to.Ptr("2017-10-01"),
	// 											to.Ptr("2017-06-01"),
	// 											to.Ptr("2016-12-01"),
	// 											to.Ptr("2016-05-01"),
	// 											to.Ptr("2016-01-01")},
	// 									}},
	// 								},
	// 								{
	// 									Name: to.Ptr("Microsoft.Storage/storageAccounts/accessTier"),
	// 									DefaultPath: to.Ptr("properties.accessTier"),
	// 									Paths: []*armresources.AliasPath{
	// 										{
	// 											Path: to.Ptr("properties.accessTier"),
	// 											APIVersions: []*string{
	// 												to.Ptr("2018-11-01"),
	// 												to.Ptr("2018-11-09"),
	// 												to.Ptr("2018-07-01"),
	// 												to.Ptr("2018-03-01-Preview"),
	// 												to.Ptr("2018-02-01"),
	// 												to.Ptr("2017-10-01"),
	// 												to.Ptr("2017-06-01"),
	// 												to.Ptr("2016-12-01"),
	// 												to.Ptr("2016-05-01"),
	// 												to.Ptr("2016-01-01"),
	// 												to.Ptr("2015-06-15"),
	// 												to.Ptr("2015-05-01-preview")},
	// 										}},
	// 									},
	// 									{
	// 										Name: to.Ptr("Microsoft.Storage/storageAccounts/enableBlobEncryption"),
	// 										DefaultPath: to.Ptr("properties.encryption.services.blob.enabled"),
	// 										Paths: []*armresources.AliasPath{
	// 											{
	// 												Path: to.Ptr("properties.encryption.services.blob.enabled"),
	// 												APIVersions: []*string{
	// 													to.Ptr("2018-11-01"),
	// 													to.Ptr("2018-11-09"),
	// 													to.Ptr("2018-07-01"),
	// 													to.Ptr("2018-03-01-Preview"),
	// 													to.Ptr("2018-02-01"),
	// 													to.Ptr("2017-10-01"),
	// 													to.Ptr("2017-06-01"),
	// 													to.Ptr("2016-12-01"),
	// 													to.Ptr("2016-05-01"),
	// 													to.Ptr("2016-01-01"),
	// 													to.Ptr("2015-06-15"),
	// 													to.Ptr("2015-05-01-preview")},
	// 											}},
	// 										},
	// 										{
	// 											Name: to.Ptr("Microsoft.Storage/storageAccounts/enableFileEncryption"),
	// 											DefaultPath: to.Ptr("properties.encryption.services.file.enabled"),
	// 											Paths: []*armresources.AliasPath{
	// 												{
	// 													Path: to.Ptr("properties.encryption.services.file.enabled"),
	// 													APIVersions: []*string{
	// 														to.Ptr("2018-11-01"),
	// 														to.Ptr("2018-11-09"),
	// 														to.Ptr("2018-07-01"),
	// 														to.Ptr("2018-03-01-Preview"),
	// 														to.Ptr("2018-02-01"),
	// 														to.Ptr("2017-10-01"),
	// 														to.Ptr("2017-06-01"),
	// 														to.Ptr("2016-12-01"),
	// 														to.Ptr("2016-05-01"),
	// 														to.Ptr("2016-01-01"),
	// 														to.Ptr("2015-06-15"),
	// 														to.Ptr("2015-05-01-preview")},
	// 												}},
	// 											},
	// 											{
	// 												Name: to.Ptr("Microsoft.Storage/storageAccounts/supportsHttpsTrafficOnly"),
	// 												DefaultMetadata: &armresources.AliasPathMetadata{
	// 													Type: to.Ptr(armresources.AliasPathTokenTypeBoolean),
	// 													Attributes: to.Ptr(armresources.AliasPathAttributesModifiable),
	// 												},
	// 												DefaultPath: to.Ptr("properties.supportsHttpsTrafficOnly"),
	// 												Paths: []*armresources.AliasPath{
	// 													{
	// 														Path: to.Ptr("properties.supportsHttpsTrafficOnly"),
	// 														APIVersions: []*string{
	// 															to.Ptr("2018-11-09"),
	// 															to.Ptr("2018-03-01-Preview"),
	// 															to.Ptr("2016-05-01"),
	// 															to.Ptr("2016-01-01"),
	// 															to.Ptr("2015-06-15"),
	// 															to.Ptr("2015-05-01-preview")},
	// 															Metadata: &armresources.AliasPathMetadata{
	// 																Type: to.Ptr(armresources.AliasPathTokenTypeNotSpecified),
	// 																Attributes: to.Ptr(armresources.AliasPathAttributesNone),
	// 															},
	// 													}},
	// 											}},
	// 											APIProfiles: []*armresources.APIProfile{
	// 												{
	// 													APIVersion: to.Ptr("2017-10-01"),
	// 													ProfileVersion: to.Ptr("2019-03-01-hybrid"),
	// 												},
	// 												{
	// 													APIVersion: to.Ptr("2016-01-01"),
	// 													ProfileVersion: to.Ptr("2017-03-09-profile"),
	// 												},
	// 												{
	// 													APIVersion: to.Ptr("2016-01-01"),
	// 													ProfileVersion: to.Ptr("2018-03-01-hybrid"),
	// 												},
	// 												{
	// 													APIVersion: to.Ptr("2017-10-01"),
	// 													ProfileVersion: to.Ptr("2018-06-01-profile"),
	// 											}},
	// 											APIVersions: []*string{
	// 												to.Ptr("2019-06-01"),
	// 												to.Ptr("2019-04-01"),
	// 												to.Ptr("2018-11-01"),
	// 												to.Ptr("2018-07-01"),
	// 												to.Ptr("2018-03-01-preview"),
	// 												to.Ptr("2018-02-01"),
	// 												to.Ptr("2017-10-01"),
	// 												to.Ptr("2017-06-01"),
	// 												to.Ptr("2016-12-01"),
	// 												to.Ptr("2016-05-01"),
	// 												to.Ptr("2016-01-01"),
	// 												to.Ptr("2015-06-15"),
	// 												to.Ptr("2015-05-01-preview")},
	// 												Capabilities: to.Ptr("SupportsTags, SupportsLocation"),
	// 												DefaultAPIVersion: to.Ptr("2019-06-01"),
	// 												Locations: []*string{
	// 													to.Ptr("East US"),
	// 													to.Ptr("East US 2"),
	// 													to.Ptr("East US 2 (Stage)"),
	// 													to.Ptr("West US"),
	// 													to.Ptr("West Europe"),
	// 													to.Ptr("East Asia"),
	// 													to.Ptr("Southeast Asia"),
	// 													to.Ptr("Japan East"),
	// 													to.Ptr("Japan West"),
	// 													to.Ptr("North Central US"),
	// 													to.Ptr("South Central US"),
	// 													to.Ptr("Central US"),
	// 													to.Ptr("North Europe"),
	// 													to.Ptr("Brazil South"),
	// 													to.Ptr("Australia East"),
	// 													to.Ptr("Australia Southeast"),
	// 													to.Ptr("South India"),
	// 													to.Ptr("Central India"),
	// 													to.Ptr("West India"),
	// 													to.Ptr("Canada East"),
	// 													to.Ptr("Canada Central"),
	// 													to.Ptr("West US 2"),
	// 													to.Ptr("West Central US"),
	// 													to.Ptr("UK South"),
	// 													to.Ptr("UK West"),
	// 													to.Ptr("Korea Central"),
	// 													to.Ptr("Korea South"),
	// 													to.Ptr("East US 2 EUAP"),
	// 													to.Ptr("Central US EUAP"),
	// 													to.Ptr("France Central"),
	// 													to.Ptr("France South"),
	// 													to.Ptr("Australia Central"),
	// 													to.Ptr("Australia Central 2"),
	// 													to.Ptr("South Africa West"),
	// 													to.Ptr("South Africa North"),
	// 													to.Ptr("UAE Central"),
	// 													to.Ptr("UAE North"),
	// 													to.Ptr("Switzerland North"),
	// 													to.Ptr("Switzerland West"),
	// 													to.Ptr("Germany West Central"),
	// 													to.Ptr("Germany North"),
	// 													to.Ptr("Norway East"),
	// 													to.Ptr("Norway West"),
	// 													to.Ptr("South Central US STG"),
	// 													to.Ptr("Brazil Southeast")},
	// 													ResourceType: to.Ptr("storageAccounts"),
	// 											}},
	// 										}
}
