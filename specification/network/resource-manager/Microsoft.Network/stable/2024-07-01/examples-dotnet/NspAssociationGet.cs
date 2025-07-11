using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspAssociationGet.json
// this example is just showing the usage of "NetworkSecurityPerimeterAssociations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterResource created on azure
// for more information of creating NetworkSecurityPerimeterResource, please refer to the document of NetworkSecurityPerimeterResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string networkSecurityPerimeterName = "nsp1";
ResourceIdentifier networkSecurityPerimeterResourceId = NetworkSecurityPerimeterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityPerimeterName);
NetworkSecurityPerimeterResource networkSecurityPerimeter = client.GetNetworkSecurityPerimeterResource(networkSecurityPerimeterResourceId);

// get the collection of this NetworkSecurityPerimeterAssociationResource
NetworkSecurityPerimeterAssociationCollection collection = networkSecurityPerimeter.GetNetworkSecurityPerimeterAssociations();

// invoke the operation
string associationName = "association1";
NullableResponse<NetworkSecurityPerimeterAssociationResource> response = await collection.GetIfExistsAsync(associationName);
NetworkSecurityPerimeterAssociationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkSecurityPerimeterAssociationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
