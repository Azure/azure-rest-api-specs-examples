using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkManagerDelete.json
// this example is just showing the usage of "NetworkManagers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkManagerResource created on azure
// for more information of creating NetworkManagerResource, please refer to the document of NetworkManagerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
ResourceIdentifier networkManagerResourceId = NetworkManagerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName);
NetworkManagerResource networkManager = client.GetNetworkManagerResource(networkManagerResourceId);

// invoke the operation
bool? force = false;
await networkManager.DeleteAsync(WaitUntil.Completed, force: force);

Console.WriteLine($"Succeeded");
