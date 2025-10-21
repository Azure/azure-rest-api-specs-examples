using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-01-01/examples/NspLinkDelete.json
// this example is just showing the usage of "NetworkSecurityPerimeterLinks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterLinkResource created on azure
// for more information of creating NetworkSecurityPerimeterLinkResource, please refer to the document of NetworkSecurityPerimeterLinkResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp1";
string linkName = "link1";
ResourceIdentifier networkSecurityPerimeterLinkResourceId = NetworkSecurityPerimeterLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName, linkName);
NetworkSecurityPerimeterLinkResource networkSecurityPerimeterLink = client.GetNetworkSecurityPerimeterLinkResource(networkSecurityPerimeterLinkResourceId);

// invoke the operation
await networkSecurityPerimeterLink.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
