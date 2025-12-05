using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerEffectiveSecurityAdminRulesList.json
// this example is just showing the usage of "ListNetworkManagerEffectiveSecurityAdminRules" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkResource created on azure
// for more information of creating VirtualNetworkResource, please refer to the document of VirtualNetworkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string virtualNetworkName = "testVirtualNetwork";
ResourceIdentifier virtualNetworkResourceId = VirtualNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName);
VirtualNetworkResource virtualNetwork = client.GetVirtualNetworkResource(virtualNetworkResourceId);

// invoke the operation and iterate over the result
NetworkManagementQueryContent content = new NetworkManagementQueryContent
{
    SkipToken = "FakeSkipTokenCode",
};
await foreach (EffectiveBaseSecurityAdminRule item in virtualNetwork.GetNetworkManagerEffectiveSecurityAdminRulesAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
