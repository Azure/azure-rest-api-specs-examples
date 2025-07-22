package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/ResourceTypeRegistrations_ListByProviderRegistration.json
func ExampleResourceTypeRegistrationsClient_NewListByProviderRegistrationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceTypeRegistrationsClient().NewListByProviderRegistrationPager("Microsoft.Contoso", nil)
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
		// page.ResourceTypeRegistrationArrayResponseWithContinuation = armproviderhub.ResourceTypeRegistrationArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.ResourceTypeRegistration{
		// 		{
		// 			Name: to.Ptr("Microsoft.Contoso/employees"),
		// 			Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/resourceTypeRegistrations"),
		// 			ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/resourceTypeRegistrations/employees"),
		// 			SystemData: &armproviderhub.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 			},
		// 			Properties: &armproviderhub.ResourceTypeRegistrationProperties{
		// 				CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
		// 				EnableAsyncOperation: to.Ptr(false),
		// 				EnableThirdPartyS2S: to.Ptr(false),
		// 				Endpoints: []*armproviderhub.ResourceTypeEndpoint{
		// 					{
		// 						APIVersions: []*string{
		// 							to.Ptr("2018-11-01-preview"),
		// 							to.Ptr("2020-01-01-preview"),
		// 							to.Ptr("2019-01-01")},
		// 							Locations: []*string{
		// 								to.Ptr("East Asia"),
		// 								to.Ptr("East US"),
		// 								to.Ptr("North Europe"),
		// 								to.Ptr("Southeast Asia"),
		// 								to.Ptr("East US 2 EUAP"),
		// 								to.Ptr("Central US EUAP"),
		// 								to.Ptr("West Europe"),
		// 								to.Ptr("West US"),
		// 								to.Ptr("West Central US"),
		// 								to.Ptr("West US 2")},
		// 								RequiredFeatures: []*string{
		// 									to.Ptr("Microsoft.Contoso/RPaaSSampleApp")},
		// 							}},
		// 							Management: &armproviderhub.ResourceTypeRegistrationPropertiesManagement{
		// 								AuthorizationOwners: []*string{
		// 									to.Ptr("RPAAS-PlatformServiceAdministrator")},
		// 									IncidentContactEmail: to.Ptr("helpme@contoso.com"),
		// 									IncidentRoutingService: to.Ptr(""),
		// 									IncidentRoutingTeam: to.Ptr(""),
		// 									ManifestOwners: []*string{
		// 										to.Ptr("SPARTA-PlatformServiceAdministrator")},
		// 										ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
		// 										ServiceTreeInfos: []*armproviderhub.ServiceTreeInfo{
		// 											{
		// 												ComponentID: to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
		// 												Readiness: to.Ptr(armproviderhub.ReadinessInDevelopment),
		// 												ServiceID: to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
		// 										}},
		// 									},
		// 									Metadata: map[string]any{
		// 									},
		// 									Notifications: []*armproviderhub.Notification{
		// 										{
		// 											NotificationType: to.Ptr(armproviderhub.NotificationTypeSubscriptionNotification),
		// 											SkipNotifications: to.Ptr(armproviderhub.SkipNotificationsDisabled),
		// 									}},
		// 									OpenAPIConfiguration: &armproviderhub.OpenAPIConfiguration{
		// 										Validation: &armproviderhub.OpenAPIValidation{
		// 											AllowNoncompliantCollectionResponse: to.Ptr(false),
		// 										},
		// 									},
		// 									ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 									Regionality: to.Ptr(armproviderhub.RegionalityRegional),
		// 									RequestHeaderOptions: &armproviderhub.ResourceTypeRegistrationPropertiesRequestHeaderOptions{
		// 										OptOutHeaders: to.Ptr(armproviderhub.OptOutHeaderTypeSystemDataCreatedByLastModifiedBy),
		// 									},
		// 									ResourceConcurrencyControlOptions: map[string]*armproviderhub.ResourceConcurrencyControlOption{
		// 										"patch": &armproviderhub.ResourceConcurrencyControlOption{
		// 											Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
		// 										},
		// 										"post": &armproviderhub.ResourceConcurrencyControlOption{
		// 											Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
		// 										},
		// 										"put": &armproviderhub.ResourceConcurrencyControlOption{
		// 											Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
		// 										},
		// 									},
		// 									ResourceGraphConfiguration: &armproviderhub.ResourceTypeRegistrationPropertiesResourceGraphConfiguration{
		// 										APIVersion: to.Ptr("2019-01-01"),
		// 										Enabled: to.Ptr(true),
		// 									},
		// 									RoutingType: to.Ptr(armproviderhub.RoutingTypeDefault),
		// 									SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
		// 										{
		// 											APIVersions: []*string{
		// 												to.Ptr("2018-11-01-preview"),
		// 												to.Ptr("2020-01-01-preview"),
		// 												to.Ptr("2019-01-01")},
		// 												SwaggerSpecFolderURI: to.Ptr("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
		// 										}},
		// 									},
		// 							}},
		// 						}
	}
}
