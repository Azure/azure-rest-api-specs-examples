using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/DirectProviderRegistrations_CreateOrUpdate.json
// this example is just showing the usage of "ProviderRegistrations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ProviderRegistrationResource
ProviderRegistrationCollection collection = subscriptionResource.GetProviderRegistrations();

// invoke the operation
string providerNamespace = "Microsoft.Contoso";
ProviderRegistrationData data = new ProviderRegistrationData
{
    Properties = new ProviderRegistrationProperties
    {
        Services = {new ResourceProviderService
        {
        ServiceName = "tags",
        Status = ResourceProviderServiceStatus.Inactive,
        }},
        ServiceName = "root",
        ProviderVersion = "2.0",
        ProviderType = ResourceProviderType.Internal,
        Management = new ResourceProviderManagement
        {
            IncidentRoutingService = "Contoso Resource Provider",
            IncidentRoutingTeam = "Contoso Triage",
            IncidentContactEmail = "helpme@contoso.com",
            ServiceTreeInfos = {new ServiceTreeInfo
            {
            ServiceId = "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
            ComponentId = "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
            Readiness = ServiceTreeReadiness.InDevelopment,
            }},
        },
        Capabilities = { new ResourceProviderCapabilities("CSP_2015-05-01", ResourceProviderCapabilitiesEffect.Allow), new ResourceProviderCapabilities("CSP_MG_2017-12-01", ResourceProviderCapabilitiesEffect.Allow) },
        DstsConfiguration = new ProviderDstsConfiguration("prds-shim")
        {
            ServiceDnsName = "prds.sparta.azure.com",
        },
        NotificationOptions = ProviderNotificationOption.EmitSpendingLimit,
        ResourceHydrationAccounts = {new ResourceHydrationAccount
        {
        AccountName = "classichydrationprodsn01",
        SubscriptionId = "e4eae963-2d15-43e6-a097-98bd75b33edd",
        }, new ResourceHydrationAccount
        {
        AccountName = "classichydrationprodch01",
        SubscriptionId = "69e69ecb-e69c-41d4-99b8-87dd12781067",
        }},
        NotificationSubscriberSettings = {new SubscriberSetting
        {
        FilterRules = {new ProviderFilterRule
        {
        FilterQuery = "Resources | where event.eventType in ('Microsoft.Network/IpAddresses/write', 'Microsoft.KeyVault/vaults/move/action')",
        EndpointInformation = {new ProviderEndpointInformation
        {
        Endpoint = "https://userrp.azure.com/arnnotify",
        EndpointType = ProviderNotificationEndpointType.Webhook,
        SchemaVersion = "3.0",
        }, new ProviderEndpointInformation
        {
        Endpoint = "https://userrp.azure.com/arnnotify",
        EndpointType = ProviderNotificationEndpointType.Eventhub,
        SchemaVersion = "3.0",
        }},
        }},
        }},
        ManagementGroupGlobalNotificationEndpoints = {new ResourceProviderEndpoint
        {
        EndpointUri = new Uri("{your_management_group_notification_endpoint}"),
        }},
        OptionalFeatures = { "Microsoft.Resources/PlatformSubscription" },
        ResourceGroupLockOptionDuringMoveBlockActionVerb = BlockActionVerb.Action,
        ServiceClientOptionsType = ServiceClientOptionsType.DisableAutomaticDecompression,
        LegacyNamespace = "legacyNamespace",
        LegacyRegistrations = { "legacyRegistration" },
        CustomManifestVersion = "2.0",
    },
    Kind = ProviderRegistrationKind.Direct,
};
ArmOperation<ProviderRegistrationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, providerNamespace, data);
ProviderRegistrationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProviderRegistrationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
