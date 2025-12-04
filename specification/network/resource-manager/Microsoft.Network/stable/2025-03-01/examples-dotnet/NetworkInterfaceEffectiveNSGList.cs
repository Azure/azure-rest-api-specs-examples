using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkInterfaceEffectiveNSGList.json
// this example is just showing the usage of "NetworkInterfaces_ListEffectiveNetworkSecurityGroups" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkInterfaceResource created on azure
// for more information of creating NetworkInterfaceResource, please refer to the document of NetworkInterfaceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkInterfaceName = "nic1";
ResourceIdentifier networkInterfaceResourceId = NetworkInterfaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkInterfaceName);
NetworkInterfaceResource networkInterface = client.GetNetworkInterfaceResource(networkInterfaceResourceId);

// invoke the operation
ArmOperation<EffectiveNetworkSecurityGroupListResult> lro = await networkInterface.GetEffectiveNetworkSecurityGroupsAsync(WaitUntil.Completed);
EffectiveNetworkSecurityGroupListResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
