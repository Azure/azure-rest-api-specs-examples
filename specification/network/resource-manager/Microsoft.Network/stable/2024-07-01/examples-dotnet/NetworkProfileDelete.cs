using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkProfileDelete.json
// this example is just showing the usage of "NetworkProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkProfileResource created on azure
// for more information of creating NetworkProfileResource, please refer to the document of NetworkProfileResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkProfileName = "networkProfile1";
ResourceIdentifier networkProfileResourceId = NetworkProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkProfileName);
NetworkProfileResource networkProfile = client.GetNetworkProfileResource(networkProfileResourceId);

// invoke the operation
await networkProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
