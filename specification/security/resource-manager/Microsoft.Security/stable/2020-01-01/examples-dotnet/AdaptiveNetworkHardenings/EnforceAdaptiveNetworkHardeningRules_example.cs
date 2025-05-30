using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/EnforceAdaptiveNetworkHardeningRules_example.json
// this example is just showing the usage of "AdaptiveNetworkHardenings_Enforce" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AdaptiveNetworkHardeningResource created on azure
// for more information of creating AdaptiveNetworkHardeningResource, please refer to the document of AdaptiveNetworkHardeningResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "rg1";
string resourceNamespace = "Microsoft.Compute";
string resourceType = "virtualMachines";
string resourceName = "vm1";
string adaptiveNetworkHardeningResourceName = "default";
ResourceIdentifier adaptiveNetworkHardeningResourceId = AdaptiveNetworkHardeningResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceNamespace, resourceType, resourceName, adaptiveNetworkHardeningResourceName);
AdaptiveNetworkHardeningResource adaptiveNetworkHardening = client.GetAdaptiveNetworkHardeningResource(adaptiveNetworkHardeningResourceId);

// invoke the operation
AdaptiveNetworkHardeningEnforceContent content = new AdaptiveNetworkHardeningEnforceContent(new RecommendedSecurityRule[]
{
new RecommendedSecurityRule
{
Name = "rule1",
Direction = SecurityTrafficDirection.Inbound,
DestinationPort = 3389,
Protocols = {SecurityTransportProtocol.Tcp},
IPAddresses = {"100.10.1.1", "200.20.2.2", "81.199.3.0/24"},
},
new RecommendedSecurityRule
{
Name = "rule2",
Direction = SecurityTrafficDirection.Inbound,
DestinationPort = 22,
Protocols = {SecurityTransportProtocol.Tcp},
IPAddresses = {},
}
}, new string[] { "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1", "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/nsg2" });
await adaptiveNetworkHardening.EnforceAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
