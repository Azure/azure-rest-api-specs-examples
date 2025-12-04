using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/QueryInboundNatRulePortMapping.json
// this example is just showing the usage of "LoadBalancers_ListInboundNatRulePortMappings" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackendAddressPoolResource created on azure
// for more information of creating BackendAddressPoolResource, please refer to the document of BackendAddressPoolResource
string subscriptionId = "subid";
string groupName = "rg1";
string loadBalancerName = "lb1";
string backendPoolName = "bp1";
ResourceIdentifier backendAddressPoolResourceId = BackendAddressPoolResource.CreateResourceIdentifier(subscriptionId, groupName, loadBalancerName, backendPoolName);
BackendAddressPoolResource backendAddressPool = client.GetBackendAddressPoolResource(backendAddressPoolResourceId);

// invoke the operation
QueryInboundNatRulePortMappingContent content = new QueryInboundNatRulePortMappingContent
{
    IPAddress = "10.0.0.4",
};
ArmOperation<BackendAddressInboundNatRulePortMappings> lro = await backendAddressPool.GetInboundNatRulePortMappingsLoadBalancerAsync(WaitUntil.Completed, content);
BackendAddressInboundNatRulePortMappings result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
