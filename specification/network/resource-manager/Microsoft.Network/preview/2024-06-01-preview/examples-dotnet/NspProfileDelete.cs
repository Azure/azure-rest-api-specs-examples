using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/preview/2024-06-01-preview/examples/NspProfileDelete.json
// this example is just showing the usage of "NetworkSecurityPerimeterProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterProfileResource created on azure
// for more information of creating NetworkSecurityPerimeterProfileResource, please refer to the document of NetworkSecurityPerimeterProfileResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp1";
string profileName = "profile1";
ResourceIdentifier networkSecurityPerimeterProfileResourceId = NetworkSecurityPerimeterProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName, profileName);
NetworkSecurityPerimeterProfileResource networkSecurityPerimeterProfile = client.GetNetworkSecurityPerimeterProfileResource(networkSecurityPerimeterProfileResourceId);

// invoke the operation
await networkSecurityPerimeterProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
