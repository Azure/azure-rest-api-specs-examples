using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VirtualNetworkGatewayNatRulePut.json
// this example is just showing the usage of "VirtualNetworkGatewayNatRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkGatewayResource created on azure
// for more information of creating VirtualNetworkGatewayResource, please refer to the document of VirtualNetworkGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualNetworkGatewayName = "gateway1";
ResourceIdentifier virtualNetworkGatewayResourceId = VirtualNetworkGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkGatewayName);
VirtualNetworkGatewayResource virtualNetworkGateway = client.GetVirtualNetworkGatewayResource(virtualNetworkGatewayResourceId);

// get the collection of this VirtualNetworkGatewayNatRuleResource
VirtualNetworkGatewayNatRuleCollection collection = virtualNetworkGateway.GetVirtualNetworkGatewayNatRules();

// invoke the operation
string natRuleName = "natRule1";
VirtualNetworkGatewayNatRuleData data = new VirtualNetworkGatewayNatRuleData
{
    VpnNatRuleType = VpnNatRuleType.Static,
    Mode = VpnNatRuleMode.EgressSnat,
    InternalMappings = {new VpnNatRuleMapping
    {
    AddressSpace = "10.4.0.0/24",
    PortRange = "200-300",
    }},
    ExternalMappings = {new VpnNatRuleMapping
    {
    AddressSpace = "192.168.21.0/24",
    PortRange = "300-400",
    }},
    IPConfigurationId = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/ipConfigurations/default",
};
ArmOperation<VirtualNetworkGatewayNatRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, natRuleName, data);
VirtualNetworkGatewayNatRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualNetworkGatewayNatRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
