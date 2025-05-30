using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_Create_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkFabrics_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkFabricResource
NetworkFabricCollection collection = resourceGroupResource.GetNetworkFabrics();

// invoke the operation
string networkFabricName = "example-fabric";
NetworkFabricData data = new NetworkFabricData(
    new AzureLocation("eastuseuap"),
    "M4-A400-A100-C16-aa",
    new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-fabricController"),
    8,
    "10.18.0.0/19",
    29249L,
    new TerminalServerConfiguration
    {
        PrimaryIPv4Prefix = "10.0.0.12/30",
        PrimaryIPv6Prefix = "4FFE:FFFF:0:CD30::a8/127",
        SecondaryIPv4Prefix = "20.0.0.13/30",
        SecondaryIPv6Prefix = "6FFE:FFFF:0:CD30::ac/127",
        Username = "username",
        Password = "xxxx",
        SerialNumber = "123456",
    },
    new ManagementNetworkConfigurationProperties(new VpnConfigurationProperties(PeeringOption.OptionA)
    {
        NetworkToNetworkInterconnectId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
        OptionBProperties = new OptionBProperties
        {
            ImportRouteTargets = { "65046:10050" },
            ExportRouteTargets = { "65046:10050" },
            RouteTargets = new RouteTargetInformation
            {
                ImportIPv4RouteTargets = { "65046:10039" },
                ImportIPv6RouteTargets = { "65046:10039" },
                ExportIPv4RouteTargets = { "65046:10039" },
                ExportIPv6RouteTargets = { "65046:10039" },
            },
        },
        OptionAProperties = new VpnConfigurationOptionAProperties
        {
            PrimaryIPv4Prefix = "10.0.0.12/30",
            PrimaryIPv6Prefix = "4FFE:FFFF:0:CD30::a8/127",
            SecondaryIPv4Prefix = "20.0.0.13/30",
            SecondaryIPv6Prefix = "6FFE:FFFF:0:CD30::ac/127",
            Mtu = 1501,
            VlanId = 3001,
            PeerAsn = 1235L,
            BfdConfiguration = new BfdConfiguration
            {
                IntervalInMilliSeconds = 300,
                Multiplier = 10,
            },
        },
    }, new VpnConfigurationProperties(PeeringOption.OptionA)
    {
        NetworkToNetworkInterconnectId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
        OptionBProperties = new OptionBProperties
        {
            ImportRouteTargets = { "65046:10050" },
            ExportRouteTargets = { "65046:10050" },
            RouteTargets = new RouteTargetInformation
            {
                ImportIPv4RouteTargets = { "65046:10039" },
                ImportIPv6RouteTargets = { "65046:10039" },
                ExportIPv4RouteTargets = { "65046:10039" },
                ExportIPv6RouteTargets = { "65046:10039" },
            },
        },
        OptionAProperties = new VpnConfigurationOptionAProperties
        {
            PrimaryIPv4Prefix = "10.0.0.14/30",
            PrimaryIPv6Prefix = "2FFE:FFFF:0:CD30::a7/126",
            SecondaryIPv4Prefix = "10.0.0.15/30",
            SecondaryIPv6Prefix = "2FFE:FFFF:0:CD30::ac/126",
            Mtu = 1500,
            VlanId = 3000,
            PeerAsn = 61234L,
            BfdConfiguration = new BfdConfiguration
            {
                IntervalInMilliSeconds = 300,
                Multiplier = 5,
            },
        },
    }))
{
    Annotation = "annotation",
    RackCount = 4,
    IPv6Prefix = "3FFE:FFFF:0:CD40::/59",
    Tags =
    {
    ["keyID"] = "keyValue"
    },
};
ArmOperation<NetworkFabricResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkFabricName, data);
NetworkFabricResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
