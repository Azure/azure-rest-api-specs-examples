using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorCreate.json
// this example is just showing the usage of "FrontDoors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this FrontDoorResource
FrontDoorCollection collection = resourceGroupResource.GetFrontDoors();

// invoke the operation
string frontDoorName = "frontDoor1";
FrontDoorData data = new FrontDoorData(new AzureLocation("westus"))
{
    RoutingRules = {new RoutingRuleData
    {
    FrontendEndpoints = {new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/frontendEndpoint1"),
    }, new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/default"),
    }},
    AcceptedProtocols = {FrontDoorProtocol.Http},
    PatternsToMatch = {"/*"},
    EnabledState = RoutingRuleEnabledState.Enabled,
    RouteConfiguration = new ForwardingConfiguration
    {
    BackendPoolId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/backendPools/backendPool1"),
    },
    RulesEngineId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/rulesEngines/rulesEngine1"),
    WebApplicationFirewallPolicyLinkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
    Name = "routingRule1",
    }},
    LoadBalancingSettings = {new FrontDoorLoadBalancingSettingsData
    {
    SampleSize = 4,
    SuccessfulSamplesRequired = 2,
    Name = "loadBalancingSettings1",
    }},
    HealthProbeSettings = {new FrontDoorHealthProbeSettingsData
    {
    Path = "/",
    Protocol = FrontDoorProtocol.Http,
    IntervalInSeconds = 120,
    HealthProbeMethod = FrontDoorHealthProbeMethod.Head,
    EnabledState = HealthProbeEnabled.Enabled,
    Name = "healthProbeSettings1",
    }},
    BackendPools = {new FrontDoorBackendPool
    {
    Backends = {new FrontDoorBackend
    {
    Address = "w3.contoso.com",
    HttpPort = 80,
    HttpsPort = 443,
    Priority = 2,
    Weight = 1,
    }, new FrontDoorBackend
    {
    Address = "contoso.com.website-us-west-2.othercloud.net",
    PrivateLinkResourceId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1"),
    PrivateLinkLocation = new AzureLocation("eastus"),
    PrivateLinkApprovalMessage = "Please approve the connection request for this Private Link",
    HttpPort = 80,
    HttpsPort = 443,
    Priority = 1,
    Weight = 2,
    }, new FrontDoorBackend
    {
    Address = "10.0.1.5",
    PrivateLinkAlias = "APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice",
    PrivateLinkApprovalMessage = "Please approve this request to connect to the Private Link",
    HttpPort = 80,
    HttpsPort = 443,
    Priority = 1,
    Weight = 1,
    }},
    LoadBalancingSettingsId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/loadBalancingSettings/loadBalancingSettings1"),
    HealthProbeSettingsId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/healthProbeSettings/healthProbeSettings1"),
    Name = "backendPool1",
    }},
    FrontendEndpoints = {new FrontendEndpointData
    {
    HostName = "www.contoso.com",
    SessionAffinityEnabledState = SessionAffinityEnabledState.Enabled,
    SessionAffinityTtlInSeconds = 60,
    WebApplicationFirewallPolicyLinkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
    Name = "frontendEndpoint1",
    }, new FrontendEndpointData
    {
    HostName = "frontDoor1.azurefd.net",
    Name = "default",
    }},
    BackendPoolsSettings = new BackendPoolsSettings
    {
        EnforceCertificateNameCheck = EnforceCertificateNameCheckEnabledState.Enabled,
        SendRecvTimeoutInSeconds = 60,
    },
    EnabledState = FrontDoorEnabledState.Enabled,
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ArmOperation<FrontDoorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, frontDoorName, data);
FrontDoorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
