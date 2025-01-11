using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/LoadBalancerDelete.json
// this example is just showing the usage of "LoadBalancers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LoadBalancerResource created on azure
// for more information of creating LoadBalancerResource, please refer to the document of LoadBalancerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string loadBalancerName = "lb";
ResourceIdentifier loadBalancerResourceId = LoadBalancerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName);
LoadBalancerResource loadBalancer = client.GetLoadBalancerResource(loadBalancerResourceId);

// invoke the operation
await loadBalancer.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
