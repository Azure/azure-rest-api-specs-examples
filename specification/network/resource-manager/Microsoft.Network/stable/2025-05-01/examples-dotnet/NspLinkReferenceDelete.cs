using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLinkReferenceDelete.json
// this example is just showing the usage of "NetworkSecurityPerimeterLinkReferences_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterLinkReferenceResource created on azure
// for more information of creating NetworkSecurityPerimeterLinkReferenceResource, please refer to the document of NetworkSecurityPerimeterLinkReferenceResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp2";
string linkReferenceName = "link1-guid";
ResourceIdentifier networkSecurityPerimeterLinkReferenceResourceId = NetworkSecurityPerimeterLinkReferenceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName, linkReferenceName);
NetworkSecurityPerimeterLinkReferenceResource networkSecurityPerimeterLinkReference = client.GetNetworkSecurityPerimeterLinkReferenceResource(networkSecurityPerimeterLinkReferenceResourceId);

// invoke the operation
await networkSecurityPerimeterLinkReference.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
