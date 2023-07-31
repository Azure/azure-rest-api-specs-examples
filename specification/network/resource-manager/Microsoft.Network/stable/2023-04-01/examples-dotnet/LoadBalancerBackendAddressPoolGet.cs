using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/LoadBalancerBackendAddressPoolGet.json
// this example is just showing the usage of "LoadBalancerBackendAddressPools_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LoadBalancerResource created on azure
// for more information of creating LoadBalancerResource, please refer to the document of LoadBalancerResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string loadBalancerName = "lb";
ResourceIdentifier loadBalancerResourceId = LoadBalancerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName);
LoadBalancerResource loadBalancer = client.GetLoadBalancerResource(loadBalancerResourceId);

// get the collection of this BackendAddressPoolResource
BackendAddressPoolCollection collection = loadBalancer.GetBackendAddressPools();

// invoke the operation
string backendAddressPoolName = "backend";
bool result = await collection.ExistsAsync(backendAddressPoolName);

Console.WriteLine($"Succeeded: {result}");
