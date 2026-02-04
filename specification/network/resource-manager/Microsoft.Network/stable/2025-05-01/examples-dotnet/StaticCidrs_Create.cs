using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/StaticCidrs_Create.json
// this example is just showing the usage of "StaticCidrs_Create" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this StaticCidrResource
StaticCidrCollection collection = ipamPool.GetStaticCidrs();

// invoke the operation
string staticCidrName = "TestStaticCidr";
StaticCidrData data = new StaticCidrData();
ArmOperation<StaticCidrResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, staticCidrName, data);
StaticCidrResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StaticCidrData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
