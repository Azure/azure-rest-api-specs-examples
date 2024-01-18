using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/LoadBalancerBackendAddressPoolDelete.json
// this example is just showing the usage of "LoadBalancerBackendAddressPools_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackendAddressPoolResource created on azure
// for more information of creating BackendAddressPoolResource, please refer to the document of BackendAddressPoolResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string loadBalancerName = "lb";
string backendAddressPoolName = "backend";
ResourceIdentifier backendAddressPoolResourceId = BackendAddressPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName, backendAddressPoolName);
BackendAddressPoolResource backendAddressPool = client.GetBackendAddressPoolResource(backendAddressPoolResourceId);

// invoke the operation
await backendAddressPool.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
