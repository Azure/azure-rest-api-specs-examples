using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/IpamPools_Delete.json
// this example is just showing the usage of "IpamPools_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IpamPoolResource created on azure
// for more information of creating IpamPoolResource, please refer to the document of IpamPoolResource
string subscriptionId = "11111111-1111-1111-1111-111111111111";
string resourceGroupName = "rg1";
string networkManagerName = "TestNetworkManager";
string poolName = "TestPool";
ResourceIdentifier ipamPoolResourceId = IpamPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, poolName);
IpamPoolResource ipamPool = client.GetIpamPoolResource(ipamPoolResourceId);

// invoke the operation
await ipamPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
