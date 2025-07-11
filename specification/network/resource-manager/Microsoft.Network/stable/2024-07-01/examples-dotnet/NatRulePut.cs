using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NatRulePut.json
// this example is just showing the usage of "NatRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnGatewayResource created on azure
// for more information of creating VpnGatewayResource, please refer to the document of VpnGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string gatewayName = "gateway1";
ResourceIdentifier vpnGatewayResourceId = VpnGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName);
VpnGatewayResource vpnGateway = client.GetVpnGatewayResource(vpnGatewayResourceId);

// get the collection of this VpnGatewayNatRuleResource
VpnGatewayNatRuleCollection collection = vpnGateway.GetVpnGatewayNatRules();

// invoke the operation
string natRuleName = "natRule1";
VpnGatewayNatRuleData data = new VpnGatewayNatRuleData
{
    VpnNatRuleType = VpnNatRuleType.Static,
    Mode = VpnNatRuleMode.EgressSnat,
    InternalMappings = {new VpnNatRuleMapping
    {
    AddressSpace = "10.4.0.0/24",
    }},
    ExternalMappings = {new VpnNatRuleMapping
    {
    AddressSpace = "192.168.21.0/24",
    }},
    IPConfigurationId = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG/ipConfigurations/default",
};
ArmOperation<VpnGatewayNatRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, natRuleName, data);
VpnGatewayNatRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VpnGatewayNatRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
