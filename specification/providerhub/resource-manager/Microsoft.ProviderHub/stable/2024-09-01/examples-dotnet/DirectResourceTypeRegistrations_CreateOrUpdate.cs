using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/DirectResourceTypeRegistrations_CreateOrUpdate.json
// this example is just showing the usage of "ResourceTypeRegistrations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProviderRegistrationResource created on azure
// for more information of creating ProviderRegistrationResource, please refer to the document of ProviderRegistrationResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
ResourceIdentifier providerRegistrationResourceId = ProviderRegistrationResource.CreateResourceIdentifier(subscriptionId, providerNamespace);
ProviderRegistrationResource providerRegistration = client.GetProviderRegistrationResource(providerRegistrationResourceId);

// get the collection of this ResourceTypeRegistrationResource
ResourceTypeRegistrationCollection collection = providerRegistration.GetResourceTypeRegistrations();

// invoke the operation
string resourceType = "employees";
ResourceTypeRegistrationData data = new ResourceTypeRegistrationData
{
    Properties = new ResourceTypeRegistrationProperties
    {
        RoutingType = ResourceRoutingType.Default,
        AdditionalOptions = AdditionalOptionResourceTypeRegistration.ProtectedAsyncOperationPolling,
        Regionality = ResourceTypeRegistrationRegionality.Regional,
        Endpoints = {new ResourceTypeEndpoint
        {
        ApiVersions = {"2020-06-01-preview"},
        Locations = {new AzureLocation("West US"), new AzureLocation("East US"), new AzureLocation("North Europe")},
        RequiredFeatures = {"<feature flag>"},
        }},
        SwaggerSpecifications = {new SwaggerSpecification
        {
        ApiVersions = {"2020-06-01-preview"},
        SwaggerSpecFolderUri = new Uri("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
        }},
        RequestHeaderOptions = new ProviderRequestHeaderOptions
        {
            OptOutHeaders = OptOutHeaderType.SystemDataCreatedByLastModifiedBy,
        },
        ResourceConcurrencyControlOptions =
        {
        ["patch"] = new ResourceConcurrencyControlOption
        {
        Policy = ResourceConcurrencyPolicy.SynchronizeBeginExtension,
        },
        ["post"] = new ResourceConcurrencyControlOption
        {
        Policy = ResourceConcurrencyPolicy.SynchronizeBeginExtension,
        },
        ["put"] = new ResourceConcurrencyControlOption
        {
        Policy = ResourceConcurrencyPolicy.SynchronizeBeginExtension,
        }
        },
        ResourceGraphConfiguration = new ResourceGraphConfiguration
        {
            IsEnabled = true,
            ApiVersion = "2019-01-01",
        },
        Management = new ResourceProviderManagement
        {
            ManifestOwners = { "SPARTA-PlatformServiceAdministrator" },
            AuthorizationOwners = { "RPAAS-PlatformServiceAdministrator" },
            IncidentRoutingService = "",
            IncidentRoutingTeam = "",
            IncidentContactEmail = "helpme@contoso.com",
            ServiceTreeInfos = {new ServiceTreeInfo
            {
            ServiceId = "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
            ComponentId = "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
            Readiness = ServiceTreeReadiness.InDevelopment,
            }},
            ResourceAccessPolicy = ResourceAccessPolicy.NotSpecified,
        },
        IsNoncompliantCollectionResponseAllowed = true,
        Metadata = { },
        Notifications = {new ProviderNotification
        {
        NotificationType = ProviderNotificationType.SubscriptionNotification,
        SkipNotifications = SkipNotification.Disabled,
        }},
        TemplateDeploymentPolicy = new TemplateDeploymentPolicy(TemplateDeploymentCapability.Preflight, new TemplateDeploymentPreflightOption("ValidationRequests, DeploymentRequests"))
        {
            PreflightNotifications = TemplateDeploymentPreflightNotification.None,
        },
        IsEmptyRoleAssignmentsAllowed = false,
        PolicyExecutionType = PolicyExecutionType.BypassPolicies,
        AvailabilityZonePolicy = AvailabilityZonePolicy.MultiZoned,
        DstsConfiguration = new ProviderDstsConfiguration("prds-shim")
        {
            ServiceDnsName = "prds.sparta.azure.com",
        },
        AsyncTimeoutRules = {new AsyncTimeoutRule
        {
        ActionName = "Microsoft.ClassicCompute/domainNames/write",
        Timeout = "PT12H",
        }},
        CommonApiVersions = { "2021-01-01" },
        ApiProfiles = {new ResourceTypeRegistrationApiProfile
        {
        ProfileVersion = "2018-03-01-hybrid",
        ApiVersion = "2018-02-01",
        }, new ResourceTypeRegistrationApiProfile
        {
        ProfileVersion = "2019-03-01-hybrid",
        ApiVersion = "2016-06-01",
        }},
        LinkedOperationRules = { new LinkedOperationRule(LinkedOperation.CrossSubscriptionResourceMove, LinkedAction.Blocked), new LinkedOperationRule(LinkedOperation.CrossResourceGroupResourceMove, LinkedAction.Validate) },
        LegacyName = "legacyName",
        LegacyNames = { "legacyName" },
        AllowedTemplateDeploymentReferenceActions = { "ListKeys", "ListSAS" },
        LegacyPolicy = new ResourceTypeRegistrationLegacyPolicy
        {
            DisallowedLegacyOperations = { ProviderLegacyOperation.Create },
            DisallowedConditions = {new LegacyDisallowedCondition
            {
            DisallowedLegacyOperations = {ProviderLegacyOperation.Create, ProviderLegacyOperation.Delete},
            Feature = "Microsoft.RP/ArmOnlyJobCollections",
            }},
        },
        ManifestLink = "https://azure.com",
        CapacityRule = new ResourceTypeRegistrationCapacityRule
        {
            CapacityPolicy = ResourceTypeRegistrationCapacityPolicy.Restricted,
            SkuAlias = "incorrectAlias",
        },
        IsAddOnPlanConversionAllowed = true,
        AllowedResourceNames = {new AllowedResourceName
        {
        Name = "name1",
        GetActionVerb = "list",
        }, new AllowedResourceName
        {
        Name = "name2",
        }},
        ResourceCache = new ResourceTypeRegistrationResourceCache
        {
            IsResourceCacheEnabled = true,
            ResourceCacheExpirationTimespan = "PT2M",
        },
        ResourceQueryManagementFilterOption = ProviderResourceQueryFilterOption.EnableSubscriptionFilterOnTenant,
        AreTagsSupported = true,
        ResourceManagementOptions = new ResourceTypeRegistrationResourceManagementOptions
        {
            BatchProvisioningSupportSupportedOperations = new ResourceManagementSupportedOperation("Get, Delete"),
            DeleteDependencies = {new ResourceTypeRegistrationDeleteDependency
            {
            LinkedProperty = "properties.edgeProfile.subscription.id",
            }},
        },
        GroupingTag = "groupingTag",
        IsAddResourceListTargetLocationsAllowed = true,
        CommonApiVersionsMergeMode = CommonApiVersionsMergeMode.Merge,
        RoutingRuleHostResourceType = "servers/databases",
    },
};
ArmOperation<ResourceTypeRegistrationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceType, data);
ResourceTypeRegistrationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceTypeRegistrationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
