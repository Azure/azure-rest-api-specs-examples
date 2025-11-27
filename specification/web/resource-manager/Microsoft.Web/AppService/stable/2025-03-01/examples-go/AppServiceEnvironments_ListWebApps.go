package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_ListWebApps.json
func ExampleEnvironmentsClient_NewListWebAppsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewListWebAppsPager("test-rg", "test-ase", &armappservice.EnvironmentsClientListWebAppsOptions{PropertiesToInclude: nil})
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
		// page.WebAppCollection = armappservice.WebAppCollection{
		// 	Value: []*armappservice.Site{
		// 		{
		// 			Name: to.Ptr("test-site"),
		// 			Type: to.Ptr("Microsoft.Web/sites"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-site"),
		// 			Location: to.Ptr("Central US EUAP"),
		// 			Properties: &armappservice.SiteProperties{
		// 				AvailabilityState: to.Ptr(armappservice.SiteAvailabilityStateNormal),
		// 				ClientAffinityEnabled: to.Ptr(true),
		// 				ClientCertEnabled: to.Ptr(false),
		// 				ClientCertMode: to.Ptr(armappservice.ClientCertModeRequired),
		// 				ContainerSize: to.Ptr[int32](0),
		// 				CustomDomainVerificationID: to.Ptr("2982A67AD520FBCD070650FC77814FB03B62927C6EFCA2F5FF3BF5DC60088845"),
		// 				DailyMemoryTimeQuota: to.Ptr[int32](0),
		// 				DefaultHostName: to.Ptr("test-site.test-ase.p.azurewebsites.net"),
		// 				Enabled: to.Ptr(true),
		// 				EnabledHostNames: []*string{
		// 					to.Ptr("test-site.test-ase.p.azurewebsites.net"),
		// 					to.Ptr("test-site.scm.test-ase.p.azurewebsites.net")},
		// 					HostNameSSLStates: []*armappservice.HostNameSSLState{
		// 						{
		// 							Name: to.Ptr("test-site.test-ase.p.azurewebsites.net"),
		// 							HostType: to.Ptr(armappservice.HostTypeStandard),
		// 							SSLState: to.Ptr(armappservice.SSLStateDisabled),
		// 						},
		// 						{
		// 							Name: to.Ptr("test-site.scm.test-ase.p.azurewebsites.net"),
		// 							HostType: to.Ptr(armappservice.HostTypeRepository),
		// 							SSLState: to.Ptr(armappservice.SSLStateDisabled),
		// 					}},
		// 					HostNames: []*string{
		// 						to.Ptr("test-site.test-ase.p.azurewebsites.net")},
		// 						HostNamesDisabled: to.Ptr(false),
		// 						HostingEnvironmentProfile: &armappservice.HostingEnvironmentProfile{
		// 							Name: to.Ptr("test-ase"),
		// 							Type: to.Ptr("Microsoft.Web/hostingEnvironments"),
		// 							ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase"),
		// 						},
		// 						HTTPSOnly: to.Ptr(false),
		// 						HyperV: to.Ptr(false),
		// 						IsXenon: to.Ptr(false),
		// 						KeyVaultReferenceIdentity: to.Ptr("SystemAssigned"),
		// 						LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-08T12:41:04.123Z"); return t}()),
		// 						OutboundIPAddresses: to.Ptr("20.112.141.120"),
		// 						PossibleOutboundIPAddresses: to.Ptr("20.112.141.120"),
		// 						RedundancyMode: to.Ptr(armappservice.RedundancyModeNone),
		// 						RepositorySiteName: to.Ptr("test-site"),
		// 						Reserved: to.Ptr(false),
		// 						ResourceGroup: to.Ptr("test-rg"),
		// 						ScmSiteAlsoStopped: to.Ptr(false),
		// 						ServerFarmID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/serverfarms/test-serverfarm"),
		// 						SiteConfig: &armappservice.SiteConfig{
		// 							FunctionAppScaleLimit: to.Ptr[int32](0),
		// 							LinuxFxVersion: to.Ptr(""),
		// 							MinimumElasticInstanceCount: to.Ptr[int32](0),
		// 							NumberOfWorkers: to.Ptr[int32](1),
		// 						},
		// 						State: to.Ptr("Running"),
		// 						StorageAccountRequired: to.Ptr(false),
		// 						UsageState: to.Ptr(armappservice.UsageStateNormal),
		// 					},
		// 			}},
		// 		}
	}
}
