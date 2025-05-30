using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/preview/2024-06-01-preview/examples/NspLinkReferenceList.json
// this example is just showing the usage of "NetworkSecurityPerimeterLinkReferences_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterResource created on azure
// for more information of creating NetworkSecurityPerimeterResource, please refer to the document of NetworkSecurityPerimeterResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp2";
ResourceIdentifier networkSecurityPerimeterResourceId = NetworkSecurityPerimeterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName);
NetworkSecurityPerimeterResource networkSecurityPerimeter = client.GetNetworkSecurityPerimeterResource(networkSecurityPerimeterResourceId);

// get the collection of this NetworkSecurityPerimeterLinkReferenceResource
NetworkSecurityPerimeterLinkReferenceCollection collection = networkSecurityPerimeter.GetNetworkSecurityPerimeterLinkReferences();

// invoke the operation and iterate over the result
await foreach (NetworkSecurityPerimeterLinkReferenceResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkSecurityPerimeterLinkReferenceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
