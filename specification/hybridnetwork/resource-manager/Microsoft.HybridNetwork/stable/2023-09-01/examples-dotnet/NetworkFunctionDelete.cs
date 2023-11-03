using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDelete.json
// this example is just showing the usage of "NetworkFunctions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFunctionResource created on azure
// for more information of creating NetworkFunctionResource, please refer to the document of NetworkFunctionResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string networkFunctionName = "testNf";
ResourceIdentifier networkFunctionResourceId = NetworkFunctionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFunctionName);
NetworkFunctionResource networkFunction = client.GetNetworkFunctionResource(networkFunctionResourceId);

// invoke the operation
await networkFunction.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
