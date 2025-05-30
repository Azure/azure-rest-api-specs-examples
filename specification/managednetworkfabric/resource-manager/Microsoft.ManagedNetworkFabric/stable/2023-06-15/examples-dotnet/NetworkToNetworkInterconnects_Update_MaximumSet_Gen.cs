using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_Update_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkToNetworkInterconnects_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkToNetworkInterconnectResource created on azure
// for more information of creating NetworkToNetworkInterconnectResource, please refer to the document of NetworkToNetworkInterconnectResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkFabricName = "example-fabric";
string networkToNetworkInterconnectName = "example-nni";
ResourceIdentifier networkToNetworkInterconnectResourceId = NetworkToNetworkInterconnectResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFabricName, networkToNetworkInterconnectName);
NetworkToNetworkInterconnectResource networkToNetworkInterconnect = client.GetNetworkToNetworkInterconnectResource(networkToNetworkInterconnectResourceId);

// invoke the operation
NetworkToNetworkInterconnectPatch patch = new NetworkToNetworkInterconnectPatch
{
    Layer2Configuration = new Layer2Configuration
    {
        Mtu = 1500,
        Interfaces = { new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice/networkInterfaces/example-networkInterface") },
    },
    OptionBLayer3Configuration = new OptionBLayer3Configuration
    {
        PeerAsn = 2345L,
        VlanId = 1235,
        PrimaryIPv4Prefix = "20.0.0.12/29",
        PrimaryIPv6Prefix = "4FFE:FFFF:0:CD30::a8/127",
        SecondaryIPv4Prefix = "20.0.0.14/29",
        SecondaryIPv6Prefix = "6FFE:FFFF:0:CD30::ac/127",
    },
    NpbStaticRouteConfiguration = new NpbStaticRouteConfiguration
    {
        BfdConfiguration = new BfdConfiguration
        {
            IntervalInMilliSeconds = 310,
            Multiplier = 15,
        },
        IPv4Routes = { new StaticRouteProperties("20.0.0.11/30", new string[] { "21.20.20.10" }) },
        IPv6Routes = { new StaticRouteProperties("4FFE:FFFF:0:CD30::ac/127", new string[] { "5FFE:FFFF:0:CD30::ac" }) },
    },
    ImportRoutePolicy = new ImportRoutePolicyInformation
    {
        ImportIPv4RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1"),
        ImportIPv6RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1"),
    },
    ExportRoutePolicy = new ExportRoutePolicyInformation
    {
        ExportIPv4RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1"),
        ExportIPv6RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy1"),
    },
    EgressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
    IngressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
};
ArmOperation<NetworkToNetworkInterconnectResource> lro = await networkToNetworkInterconnect.UpdateAsync(WaitUntil.Completed, patch);
NetworkToNetworkInterconnectResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkToNetworkInterconnectData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
