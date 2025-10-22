using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/NewRegionFrontloadRelease_GenerateManifest.json
// this example is just showing the usage of "NewRegionFrontloadRelease_GenerateManifest" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
ProviderFrontloadPayload properties = new ProviderFrontloadPayload(new ProviderFrontloadPayloadProperties(
    "Rollout",
    "Microsoft.Contoso",
    "Israel Central",
    "eastus",
    AvailableCheckInManifestEnvironment.Prod,
    ServiceFeatureFlagAction.DoNotCreate,
    new string[] { "servers" },
    new string[] { "monitors" },
    new ManifestLevelPropertyBag
    {
        ResourceHydrationAccounts = {new ResourceHydrationAccount
        {
        AccountName = "classichydrationprodsn01",
        SubscriptionId = "e4eae963-2d15-43e6-a097-98bd75b33edd",
        }},
    },
    new ResourceTypeEndpointBase(
    true,
    new string[] { "2024-04-01-preview" },
    new Uri("https://resource-endpoint.com/"),
    new string[] { "East US" },
    new string[] { "<feature flag>" },
    default,
    XmlConvert.ToTimeSpan("PT20S"),
    ProviderEndpointType.Production,
    new ProviderDstsConfiguration("resourceprovider")
    {
        ServiceDnsName = "messaging.azure-ppe.net",
    },
    "http://endpointuri/westus/skus",
    "2024-04-01-preview",
    new string[] { "zone1" }),
    new string[] { "apiversion" }));
ResourceProviderManifest result = await providerRegistration.GenerateManifestNewRegionFrontloadReleaseAsync(properties);

Console.WriteLine($"Succeeded: {result}");
