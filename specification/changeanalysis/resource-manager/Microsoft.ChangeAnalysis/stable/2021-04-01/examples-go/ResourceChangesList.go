package armchangeanalysis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ResourceChangesList.json
func ExampleResourceChangesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchangeanalysis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceChangesClient().NewListPager("subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Web/sites/mysite", func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-25T12:09:03.141Z"); return t }(), func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T12:09:03.141Z"); return t }(), &armchangeanalysis.ResourceChangesClientListOptions{SkipToken: nil})
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
		// page.ChangeList = armchangeanalysis.ChangeList{
		// 	Value: []*armchangeanalysis.Change{
		// 		{
		// 			Name: to.Ptr("ARG_23fa00fd-dda0-4268-b482-2076825cf165_970d8c6d-6b78-4270-92ef-88d5aa2b5f0b_132316363294700000_132316498613900000"),
		// 			Type: to.Ptr("Microsoft.ChangeAnalysis/resourceChanges"),
		// 			ID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Web/sites/mysite/providers/Microsoft.ChangeAnalysis/resourceChanges/ARG_23fa00fd-dda0-4268-b482-2076825cf165_970d8c6d-6b78-4270-92ef-88d5aa2b5f0b_132316363294700000_132316498613900000"),
		// 			Properties: &armchangeanalysis.ChangeProperties{
		// 				ChangeType: to.Ptr(armchangeanalysis.ChangeTypeUpdate),
		// 				InitiatedByList: []*string{
		// 					to.Ptr("ellen@contoso.com")},
		// 					PropertyChanges: []*armchangeanalysis.PropertyChange{
		// 						{
		// 							Description: to.Ptr("The thumbprint of the certificate"),
		// 							ChangeCategory: to.Ptr(armchangeanalysis.ChangeCategoryUser),
		// 							ChangeType: to.Ptr(armchangeanalysis.ChangeTypeUpdate),
		// 							DisplayName: to.Ptr("publicCertificates[\"AppCert\"].properties.thumbprint"),
		// 							IsDataMasked: to.Ptr(false),
		// 							JSONPath: to.Ptr("value[1].properties.thumbprint"),
		// 							Level: to.Ptr(armchangeanalysis.LevelImportant),
		// 							NewValue: to.Ptr("3F2DF554-B063-4383-8BD3-4970BCF20A7E"),
		// 							OldValue: to.Ptr("21D0482F-E91E-4C14-8078-65BFDCDBCA64"),
		// 					}},
		// 					ResourceID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Web/sites/mysite"),
		// 					TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T02:17:41.390Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AST_b2ecd7ec-00dd-4d44-bd19-fefc88790c05_da36e22e-f541-44bd-8d89-8f10a27af4ad_132592897881782408_132592957873393845"),
		// 				Type: to.Ptr("Microsoft.ChangeAnalysis/changes"),
		// 				ID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Web/sites/mysite/extensions/AppStateTracker/providers/Microsoft.ChangeAnalysis/resourceChanges/AST_b2ecd7ec-00dd-4d44-bd19-fefc88790c05_da36e22e-f541-44bd-8d89-8f10a27af4ad_132592897881782408_132592957873393845"),
		// 				Properties: &armchangeanalysis.ChangeProperties{
		// 					ChangeType: to.Ptr(armchangeanalysis.ChangeTypeUpdate),
		// 					InitiatedByList: []*string{
		// 						to.Ptr("ellen@contoso.com")},
		// 						PropertyChanges: []*armchangeanalysis.PropertyChange{
		// 							{
		// 								Description: to.Ptr("Application setting"),
		// 								ChangeCategory: to.Ptr(armchangeanalysis.ChangeCategoryUser),
		// 								ChangeType: to.Ptr(armchangeanalysis.ChangeTypeUpdate),
		// 								DisplayName: to.Ptr("APPSETTING_DB_CONNSTR"),
		// 								IsDataMasked: to.Ptr(false),
		// 								JSONPath: to.Ptr("environment.environmentVariables.APPSETTING_DB_CONNSTR"),
		// 								Level: to.Ptr(armchangeanalysis.LevelImportant),
		// 								NewValue: to.Ptr("<new database connection string>"),
		// 								OldValue: to.Ptr("<old database connection string>"),
		// 						}},
		// 						ResourceID: to.Ptr("/subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Web/sites/mysite"),
		// 						TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T02:17:41.390Z"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}
