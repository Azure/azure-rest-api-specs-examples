using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/InternalNetworks_Update_MaximumSet_Gen.json
// this example is just showing the usage of "InternalNetworks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InternalNetworkResource created on azure
// for more information of creating InternalNetworkResource, please refer to the document of InternalNetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l3IsolationDomainName = "example-l3domain";
string internalNetworkName = "example-internalnetwork";
ResourceIdentifier internalNetworkResourceId = InternalNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName, internalNetworkName);
InternalNetworkResource internalNetwork = client.GetInternalNetworkResource(internalNetworkResourceId);

// invoke the operation
InternalNetworkPatch patch = new InternalNetworkPatch()
{
    Mtu = 1500,
    ConnectedIPv4Subnets =
    {
    new ConnectedSubnet()
    {
    Prefix = "10.0.0.0/24",
    }
    },
    ConnectedIPv6Subnets =
    {
    new ConnectedSubnet()
    {
    Prefix = "3FFE:FFFF:0:CD30::a0/29",
    }
    },
    StaticRouteConfiguration = new StaticRouteConfiguration()
    {
        BfdConfiguration = new BfdConfiguration(),
        IPv4Routes =
        {
        new StaticRouteProperties("10.1.0.0/24",new string[]
        {
        "10.0.0.1"
        })
        },
        IPv6Routes =
        {
        new StaticRouteProperties("2fff::/64",new string[]
        {
        "2ffe::1"
        })
        },
    },
    BgpConfiguration = new BgpConfiguration(6)
    {
        BfdConfiguration = new BfdConfiguration(),
        DefaultRouteOriginate = BooleanEnumProperty.True,
        AllowAS = 1,
        AllowASOverride = AllowASOverride.Enable,
        IPv4ListenRangePrefixes =
        {
        "10.1.0.0/25"
        },
        IPv6ListenRangePrefixes =
        {
        "2fff::/66"
        },
        IPv4NeighborAddress =
        {
        new NeighborAddress()
        {
        Address = "10.1.0.0",
        }
        },
        IPv6NeighborAddress =
        {
        new NeighborAddress()
        {
        Address = "2fff::",
        }
        },
    },
    ImportRoutePolicyId = "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName1",
    ExportRoutePolicyId = "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName2",
};
ArmOperation<InternalNetworkResource> lro = await internalNetwork.UpdateAsync(WaitUntil.Completed, patch);
InternalNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InternalNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
