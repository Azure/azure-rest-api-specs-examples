package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/ResourceTypeRegistrations_CreateOrUpdate.json
func ExampleResourceTypeRegistrationsClient_BeginCreateOrUpdate_resourceTypeRegistrationsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourceTypeRegistrationsClient().BeginCreateOrUpdate(ctx, "Microsoft.Contoso", "employees", armproviderhub.ResourceTypeRegistration{
		Properties: &armproviderhub.ResourceTypeRegistrationProperties{
			RoutingType:                to.Ptr(armproviderhub.RoutingTypeDefault),
			Regionality:                to.Ptr(armproviderhub.RegionalityRegional),
			CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
			Endpoints: []*armproviderhub.ResourceTypeEndpoint{
				{
					APIVersions: []*string{
						to.Ptr("2020-06-01-preview"),
					},
					Locations: []*string{
						to.Ptr("West US"),
						to.Ptr("East US"),
						to.Ptr("North Europe"),
					},
					RequiredFeatures: []*string{
						to.Ptr("<feature flag>"),
					},
				},
			},
			ResourceConcurrencyControlOptions: map[string]*armproviderhub.ResourceConcurrencyControlOption{
				"put": {
					Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
				},
				"patch": {
					Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
				},
				"post": {
					Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
				},
			},
			SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
				{
					APIVersions: []*string{
						to.Ptr("2020-06-01-preview"),
					},
					SwaggerSpecFolderURI: to.Ptr("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
				},
			},
			ResourceGraphConfiguration: &armproviderhub.ResourceTypeRegistrationPropertiesResourceGraphConfiguration{
				Enabled:    to.Ptr(true),
				APIVersion: to.Ptr("2019-01-01"),
			},
			Management: &armproviderhub.ResourceTypeRegistrationPropertiesManagement{
				ManifestOwners: []*string{
					to.Ptr("Contoso-PlatformServiceAdministrator"),
				},
				AuthorizationOwners: []*string{
					to.Ptr("RPAAS-PlatformServiceAdministrator"),
				},
				IncidentRoutingService: to.Ptr(""),
				IncidentRoutingTeam:    to.Ptr(""),
				IncidentContactEmail:   to.Ptr("helpme@contoso.com"),
				ResourceAccessPolicy:   to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
			},
			Metadata: map[string]any{},
			Notifications: []*armproviderhub.Notification{
				{
					NotificationType:  to.Ptr(armproviderhub.NotificationTypeSubscriptionNotification),
					SkipNotifications: to.Ptr(armproviderhub.SkipNotificationsDisabled),
				},
			},
			OpenAPIConfiguration: &armproviderhub.OpenAPIConfiguration{
				Validation: &armproviderhub.OpenAPIValidation{
					AllowNoncompliantCollectionResponse: to.Ptr(true),
				},
			},
			RequestHeaderOptions: &armproviderhub.ResourceTypeRegistrationPropertiesRequestHeaderOptions{
				OptOutHeaders: to.Ptr(armproviderhub.OptOutHeaderTypeSystemDataCreatedByLastModifiedBy),
			},
			ThrottlingRules: []*armproviderhub.ThrottlingRule{
				{
					Action: to.Ptr("Microsoft.Foo/checkNameAvailability/write"),
					Metrics: []*armproviderhub.ThrottlingMetric{
						{
							Type:       to.Ptr(armproviderhub.ThrottlingMetricTypeNumberOfRequests),
							BucketSize: to.Ptr("XLarge"),
							Limit:      to.Ptr[int64](1),
						},
					},
				},
			},
			PrivateEndpointConfiguration: &armproviderhub.PrivateEndpointConfiguration{
				MinAPIVersion: to.Ptr("2022-10-01"),
				GroupConnectivityInformation: []*armproviderhub.GroupConnectivityInformation{
					{
						GroupID: to.Ptr("Sql"),
						RequiredMembers: []*string{
							to.Ptr("Sql_Member"),
						},
						RequiredZoneNames: []*string{
							to.Ptr("Zone"),
						},
						RedirectMapID: to.Ptr("test"),
					},
				},
			},
			WriteLock: &armproviderhub.WriteLockConfiguration{
				State: to.Ptr(armproviderhub.WriteLockStateEnabled),
			},
			MarketplaceType: to.Ptr(armproviderhub.MarketplaceTypeProviderHub),
			ResourceManagementOptions: &armproviderhub.ResourceTypeRegistrationPropertiesResourceManagementOptions{
				BatchProvisioningSupport: &armproviderhub.ResourceTypeRegistrationPropertiesResourceManagementOptionsBatchProvisioningSupport{
					MaxBatchSize: to.Ptr[int64](10),
					ActionConfigurations: []*armproviderhub.ActionConfiguration{
						{
							AuthorizationAction: to.Ptr("Microsoft.Contoso/authorize"),
							MaxBatchSize:        to.Ptr[int64](5),
						},
					},
					BatchContractVersion: to.Ptr("2020-06-01-preview"),
					MaxNestedBatchSize:   to.Ptr[int64](5),
					RequiredFeatures: []*string{
						to.Ptr("Microsoft.Contoso/feature1"),
					},
					SupportedOperations: to.Ptr(armproviderhub.SupportedOperationsGet),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.ResourceTypeRegistrationsClientCreateOrUpdateResponse{
	// 	ResourceTypeRegistration: armproviderhub.ResourceTypeRegistration{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/resourceTypeRegistrations/employees"),
	// 		Name: to.Ptr("Microsoft.Contoso/employees"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/resourceTypeRegistrations"),
	// 		Properties: &armproviderhub.ResourceTypeRegistrationProperties{
	// 			RoutingType: to.Ptr(armproviderhub.RoutingTypeDefault),
	// 			Regionality: to.Ptr(armproviderhub.RegionalityRegional),
	// 			CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
	// 			Endpoints: []*armproviderhub.ResourceTypeEndpoint{
	// 				{
	// 					APIVersions: []*string{
	// 						to.Ptr("2018-11-01-preview"),
	// 						to.Ptr("2020-01-01-preview"),
	// 						to.Ptr("2019-01-01"),
	// 					},
	// 					Locations: []*string{
	// 						to.Ptr("East Asia"),
	// 						to.Ptr("East US"),
	// 						to.Ptr("North Europe"),
	// 						to.Ptr("Southeast Asia"),
	// 						to.Ptr("East US 2 EUAP"),
	// 						to.Ptr("Central US EUAP"),
	// 						to.Ptr("West Europe"),
	// 						to.Ptr("West US"),
	// 						to.Ptr("West Central US"),
	// 						to.Ptr("West US 2"),
	// 					},
	// 					RequiredFeatures: []*string{
	// 						to.Ptr("Microsoft.Contoso/RPaaSSampleApp"),
	// 					},
	// 				},
	// 			},
	// 			SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
	// 				{
	// 					APIVersions: []*string{
	// 						to.Ptr("2018-11-01-preview"),
	// 						to.Ptr("2020-01-01-preview"),
	// 						to.Ptr("2019-01-01"),
	// 					},
	// 					SwaggerSpecFolderURI: to.Ptr("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
	// 				},
	// 			},
	// 			EnableAsyncOperation: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 			EnableThirdPartyS2S: to.Ptr(false),
	// 			ResourceDeletionPolicy: to.Ptr(armproviderhub.RPaaSResourceDeletionPolicyCascadeDeleteProxyOnlyChildren),
	// 			ResourceDeletionPolicies: []*armproviderhub.ResourceDeletionPolicyAndProperties{
	// 				{
	// 					PolicyName: to.Ptr(armproviderhub.ResourceDeletionPolicySoftDelete),
	// 					Properties: &armproviderhub.ResourceDeletionPolicyProperties{
	// 						MinimumRetentionTime: to.Ptr("P0DT1H0M"),
	// 						MaximumRetentionTime: to.Ptr("P1DT0H0M"),
	// 					},
	// 				},
	// 			},
	// 			ResourceConcurrencyControlOptions: map[string]*armproviderhub.ResourceConcurrencyControlOption{
	// 				"put": &armproviderhub.ResourceConcurrencyControlOption{
	// 					Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
	// 				},
	// 				"patch": &armproviderhub.ResourceConcurrencyControlOption{
	// 					Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
	// 				},
	// 				"post": &armproviderhub.ResourceConcurrencyControlOption{
	// 					Policy: to.Ptr(armproviderhub.PolicySynchronizeBeginExtension),
	// 				},
	// 			},
	// 			ResourceGraphConfiguration: &armproviderhub.ResourceTypeRegistrationPropertiesResourceGraphConfiguration{
	// 				Enabled: to.Ptr(true),
	// 				APIVersion: to.Ptr("2019-01-01"),
	// 			},
	// 			Management: &armproviderhub.ResourceTypeRegistrationPropertiesManagement{
	// 				ManifestOwners: []*string{
	// 					to.Ptr("Contoso-PlatformServiceAdministrator"),
	// 				},
	// 				IncidentRoutingService: to.Ptr(""),
	// 				IncidentRoutingTeam: to.Ptr(""),
	// 				IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 				ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
	// 			},
	// 			Metadata: map[string]any{
	// 			},
	// 			Notifications: []*armproviderhub.Notification{
	// 				{
	// 					NotificationType: to.Ptr(armproviderhub.NotificationTypeSubscriptionNotification),
	// 					SkipNotifications: to.Ptr(armproviderhub.SkipNotificationsDisabled),
	// 				},
	// 			},
	// 			OpenAPIConfiguration: &armproviderhub.OpenAPIConfiguration{
	// 				Validation: &armproviderhub.OpenAPIValidation{
	// 					AllowNoncompliantCollectionResponse: to.Ptr(true),
	// 				},
	// 			},
	// 			RequestHeaderOptions: &armproviderhub.ResourceTypeRegistrationPropertiesRequestHeaderOptions{
	// 				OptOutHeaders: to.Ptr(armproviderhub.OptOutHeaderTypeSystemDataCreatedByLastModifiedBy),
	// 			},
	// 			ThrottlingRules: []*armproviderhub.ThrottlingRule{
	// 				{
	// 					Action: to.Ptr("Microsoft.Foo/checkNameAvailability/write"),
	// 					Metrics: []*armproviderhub.ThrottlingMetric{
	// 						{
	// 							Type: to.Ptr(armproviderhub.ThrottlingMetricTypeNumberOfRequests),
	// 							BucketSize: to.Ptr("XLarge"),
	// 							Limit: to.Ptr[int64](1),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			WriteLock: &armproviderhub.WriteLockConfiguration{
	// 				State: to.Ptr(armproviderhub.WriteLockStateEnabled),
	// 			},
	// 			ResourceManagementOptions: &armproviderhub.ResourceTypeRegistrationPropertiesResourceManagementOptions{
	// 				BatchProvisioningSupport: &armproviderhub.ResourceTypeRegistrationPropertiesResourceManagementOptionsBatchProvisioningSupport{
	// 					MaxBatchSize: to.Ptr[int64](10),
	// 					ActionConfigurations: []*armproviderhub.ActionConfiguration{
	// 						{
	// 							AuthorizationAction: to.Ptr("Microsoft.Contoso/authorize"),
	// 							MaxBatchSize: to.Ptr[int64](5),
	// 						},
	// 					},
	// 					BatchContractVersion: to.Ptr("2020-06-01-preview"),
	// 					MaxNestedBatchSize: to.Ptr[int64](5),
	// 					RequiredFeatures: []*string{
	// 						to.Ptr("Microsoft.Contoso/feature1"),
	// 					},
	// 					SupportedOperations: to.Ptr(armproviderhub.SupportedOperationsGet),
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armproviderhub.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 		},
	// 	},
	// }
}
