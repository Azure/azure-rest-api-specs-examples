using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/ExternalNetworks_Update_MaximumSet_Gen.json
// this example is just showing the usage of "ExternalNetworks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricExternalNetworkResource created on azure
// for more information of creating NetworkFabricExternalNetworkResource, please refer to the document of NetworkFabricExternalNetworkResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string l3IsolationDomainName = "example-l3domain";
string externalNetworkName = "example-externalnetwork";
ResourceIdentifier networkFabricExternalNetworkResourceId = NetworkFabricExternalNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName, externalNetworkName);
NetworkFabricExternalNetworkResource networkFabricExternalNetwork = client.GetNetworkFabricExternalNetworkResource(networkFabricExternalNetworkResourceId);

// invoke the operation
NetworkFabricExternalNetworkPatch patch = new NetworkFabricExternalNetworkPatch
{
    Annotation = "annotation1",
    ImportRoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
    ExportRoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
    ImportRoutePolicy = new ImportRoutePolicy
    {
        ImportIPv4RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
        ImportIPv6RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
    },
    ExportRoutePolicy = new ExportRoutePolicy
    {
        ExportIPv4RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
        ExportIPv6RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName"),
    },
    PeeringOption = PeeringOption.OptionA,
    OptionBProperties = new L3OptionBProperties
    {
        ImportRouteTargets = { "65046:10039" },
        ExportRouteTargets = { "65046:10039" },
        RouteTargets = new RouteTargetInformation
        {
            ImportIPv4RouteTargets = { "65046:10039" },
            ImportIPv6RouteTargets = { "65046:10039" },
            ExportIPv4RouteTargets = { "65046:10039" },
            ExportIPv6RouteTargets = { "65046:10039" },
        },
    },
    OptionAProperties = new ExternalNetworkPatchOptionAProperties
    {
        Mtu = 1500,
        VlanId = 1001,
        PeerAsn = 65047L,
        BfdConfiguration = new BfdConfiguration
        {
            IntervalInMilliSeconds = 300,
            Multiplier = 15,
        },
        IngressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
        EgressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
        PrimaryIPv4Prefix = "10.1.1.0/30",
        PrimaryIPv6Prefix = "3FFE:FFFF:0:CD30::a0/126",
        SecondaryIPv4Prefix = "10.1.1.4/30",
        SecondaryIPv6Prefix = "3FFE:FFFF:0:CD30::a4/126",
    },
};
ArmOperation<NetworkFabricExternalNetworkResource> lro = await networkFabricExternalNetwork.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricExternalNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricExternalNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
