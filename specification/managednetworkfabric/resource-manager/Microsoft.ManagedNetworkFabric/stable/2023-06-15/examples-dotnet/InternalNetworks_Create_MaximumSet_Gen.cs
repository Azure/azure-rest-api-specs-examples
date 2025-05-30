using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/InternalNetworks_Create_MaximumSet_Gen.json
// this example is just showing the usage of "InternalNetworks_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricL3IsolationDomainResource created on azure
// for more information of creating NetworkFabricL3IsolationDomainResource, please refer to the document of NetworkFabricL3IsolationDomainResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string l3IsolationDomainName = "example-l3domain";
ResourceIdentifier networkFabricL3IsolationDomainResourceId = NetworkFabricL3IsolationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName);
NetworkFabricL3IsolationDomainResource networkFabricL3IsolationDomain = client.GetNetworkFabricL3IsolationDomainResource(networkFabricL3IsolationDomainResourceId);

// get the collection of this NetworkFabricInternalNetworkResource
NetworkFabricInternalNetworkCollection collection = networkFabricL3IsolationDomain.GetNetworkFabricInternalNetworks();

// invoke the operation
string internalNetworkName = "example-internalnetwork";
NetworkFabricInternalNetworkData data = new NetworkFabricInternalNetworkData(755)
{
    Annotation = "annotation",
    Mtu = 1500,
    ConnectedIPv4Subnets = {new ConnectedSubnet("10.0.0.0/24")
    {
    Annotation = "annotation",
    }},
    ConnectedIPv6Subnets = {new ConnectedSubnet("3FFE:FFFF:0:CD30::a0/29")
    {
    Annotation = "annotation",
    }},
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
    IngressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
    EgressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
    IsMonitoringEnabled = IsMonitoringEnabled.True,
    Extension = StaticRouteConfigurationExtension.NoExtension,
    BgpConfiguration = new InternalNetworkBgpConfiguration
    {
        BfdConfiguration = new BfdConfiguration
        {
            IntervalInMilliSeconds = 300,
            Multiplier = 5,
        },
        DefaultRouteOriginate = NetworkFabricBooleanValue.True,
        AllowAS = 10,
        AllowASOverride = AllowASOverride.Enable,
        PeerAsn = 61234L,
        IPv4ListenRangePrefixes = { "10.1.0.0/25" },
        IPv6ListenRangePrefixes = { "2fff::/66" },
        IPv4NeighborAddress = {new NeighborAddress
        {
        Address = "10.1.0.0",
        }},
        IPv6NeighborAddress = {new NeighborAddress
        {
        Address = "2fff::",
        }},
        Annotation = "annotation",
    },
    StaticRouteConfiguration = new InternalNetworkStaticRouteConfiguration
    {
        Extension = StaticRouteConfigurationExtension.NoExtension,
        BfdConfiguration = new BfdConfiguration
        {
            IntervalInMilliSeconds = 300,
            Multiplier = 15,
        },
        IPv4Routes = { new StaticRouteProperties("jffgck", new string[] { "10.0.0.1" }) },
        IPv6Routes = { new StaticRouteProperties("2fff::/64", new string[] { "3ffe::1" }) },
    },
};
ArmOperation<NetworkFabricInternalNetworkResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, internalNetworkName, data);
NetworkFabricInternalNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricInternalNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
