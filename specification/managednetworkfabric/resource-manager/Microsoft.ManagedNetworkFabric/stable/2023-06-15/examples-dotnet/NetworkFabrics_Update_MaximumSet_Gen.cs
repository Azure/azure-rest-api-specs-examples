using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_Update_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkFabrics_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
NetworkFabricPatch patch = new NetworkFabricPatch()
{
    Annotation = "annotation1",
    RackCount = 6,
    ServerCountPerRack = 10,
    IPv4Prefix = "10.18.0.0/17",
    IPv6Prefix = "3FFE:FFFF:0:CD40::/60",
    FabricAsn = 12345,
    TerminalServerConfiguration = new NetworkFabricPatchablePropertiesTerminalServerConfiguration()
    {
        PrimaryIPv4Prefix = "10.0.0.12/30",
        PrimaryIPv6Prefix = "4FFE:FFFF:0:CD30::a8/127",
        SecondaryIPv4Prefix = "40.0.0.14/30",
        SecondaryIPv6Prefix = "6FFE:FFFF:0:CD30::ac/127",
        Username = "username1",
        Password = "xxxxxxxx",
        SerialNumber = "1234567",
    },
    ManagementNetworkConfiguration = new ManagementNetworkConfigurationPatchableProperties()
    {
        InfrastructureVpnConfiguration = new VpnConfigurationPatchableProperties()
        {
            NetworkToNetworkInterconnectId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
            PeeringOption = PeeringOption.OptionB,
            OptionBProperties = new OptionBProperties()
            {
                ImportRouteTargets =
                {
                "65046:10050"
                },
                ExportRouteTargets =
                {
                "65046:10050"
                },
                RouteTargets = new RouteTargetInformation()
                {
                    ImportIPv4RouteTargets =
                    {
                    "65046:10050"
                    },
                    ImportIPv6RouteTargets =
                    {
                    "65046:10050"
                    },
                    ExportIPv4RouteTargets =
                    {
                    "65046:10050"
                    },
                    ExportIPv6RouteTargets =
                    {
                    "65046:10050"
                    },
                },
            },
            OptionAProperties = new VpnConfigurationPatchableOptionAProperties()
            {
                PrimaryIPv4Prefix = "10.0.0.12/30",
                PrimaryIPv6Prefix = "4FFE:FFFF:0:CD30::a8/127",
                SecondaryIPv4Prefix = "20.0.0.13/30",
                SecondaryIPv6Prefix = "6FFE:FFFF:0:CD30::ac/127",
                Mtu = 1501,
                VlanId = 3001,
                PeerAsn = 1235,
                BfdConfiguration = new BfdConfiguration()
                {
                    IntervalInMilliSeconds = 300,
                    Multiplier = 10,
                },
            },
        },
        WorkloadVpnConfiguration = new VpnConfigurationPatchableProperties()
        {
            NetworkToNetworkInterconnectId = new ResourceIdentifier("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourceGroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-fabric/networkToNetworkInterconnects/example-nni"),
            PeeringOption = PeeringOption.OptionA,
            OptionBProperties = new OptionBProperties()
            {
                ImportRouteTargets =
                {
                "65046:10050"
                },
                ExportRouteTargets =
                {
                "65046:10050"
                },
                RouteTargets = new RouteTargetInformation()
                {
                    ImportIPv4RouteTargets =
                    {
                    "65046:10050"
                    },
                    ImportIPv6RouteTargets =
                    {
                    "65046:10050"
                    },
                    ExportIPv4RouteTargets =
                    {
                    "65046:10050"
                    },
                    ExportIPv6RouteTargets =
                    {
                    "65046:10050"
                    },
                },
            },
            OptionAProperties = new VpnConfigurationPatchableOptionAProperties()
            {
                PrimaryIPv4Prefix = "10.0.0.14/30",
                PrimaryIPv6Prefix = "2FFE:FFFF:0:CD30::a7/126",
                SecondaryIPv4Prefix = "10.0.0.15/30",
                SecondaryIPv6Prefix = "2FFE:FFFF:0:CD30::ac/126",
                Mtu = 1500,
                VlanId = 3000,
                PeerAsn = 61234,
                BfdConfiguration = new BfdConfiguration()
                {
                    IntervalInMilliSeconds = 300,
                    Multiplier = 5,
                },
            },
        },
    },
    Tags =
    {
    ["keyID"] = "KeyValue",
    },
};
ArmOperation<NetworkFabricResource> lro = await networkFabric.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
