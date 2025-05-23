using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_Create_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkToNetworkInterconnects_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricResource created on azure
// for more information of creating NetworkFabricResource, please refer to the document of NetworkFabricResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkFabricName = "example-fabric";
ResourceIdentifier networkFabricResourceId = NetworkFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFabricName);
NetworkFabricResource networkFabric = client.GetNetworkFabricResource(networkFabricResourceId);

// get the collection of this NetworkToNetworkInterconnectResource
NetworkToNetworkInterconnectCollection collection = networkFabric.GetNetworkToNetworkInterconnects();

// invoke the operation
string networkToNetworkInterconnectName = "example-nni";
NetworkToNetworkInterconnectData data = new NetworkToNetworkInterconnectData(NetworkFabricBooleanValue.True)
{
    NniType = NniType.CE,
    IsManagementType = IsManagementType.True,
    Layer2Configuration = new Layer2Configuration
    {
        Mtu = 1500,
        Interfaces = { new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice/networkInterfaces/example-networkInterface") },
    },
    OptionBLayer3Configuration = new NetworkToNetworkInterconnectOptionBLayer3Configuration
    {
        PeerAsn = 61234L,
        VlanId = 1234,
        PrimaryIPv4Prefix = "10.0.0.12/30",
        PrimaryIPv6Prefix = "4FFE:FFFF:0:CD30::a8/127",
        SecondaryIPv4Prefix = "40.0.0.14/30",
        SecondaryIPv6Prefix = "6FFE:FFFF:0:CD30::ac/127",
    },
    NpbStaticRouteConfiguration = new NpbStaticRouteConfiguration
    {
        BfdConfiguration = new BfdConfiguration
        {
            IntervalInMilliSeconds = 300,
            Multiplier = 25,
        },
        IPv4Routes = { new StaticRouteProperties("20.0.0.12/30", new string[] { "21.20.20.20" }) },
        IPv6Routes = { new StaticRouteProperties("3FFE:FFFF:0:CD30::ac/127", new string[] { "4FFE:FFFF:0:CD30::ac" }) },
    },
    ImportRoutePolicy = new ImportRoutePolicyInformation
    {
        ImportIPv4RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
        ImportIPv6RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
    },
    ExportRoutePolicy = new ExportRoutePolicyInformation
    {
        ExportIPv4RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
        ExportIPv6RoutePolicyId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/routePolicies/example-routePolicy"),
    },
    EgressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
    IngressAclId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/accessControlLists/example-acl"),
};
ArmOperation<NetworkToNetworkInterconnectResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkToNetworkInterconnectName, data);
NetworkToNetworkInterconnectResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkToNetworkInterconnectData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
