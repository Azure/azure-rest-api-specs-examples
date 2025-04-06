using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: 2024-12-01/PublicCloudConnectors_Get.json
// this example is just showing the usage of "PublicCloudConnector_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5ACC4579-DB34-4C2F-8F8C-25061168F342";
string resourceGroupName = "rgpublicCloud";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PublicCloudConnectorResource
PublicCloudConnectorCollection collection = resourceGroupResource.GetPublicCloudConnectors();

// invoke the operation
string publicCloudConnector = "rzygvnpsnrdylwzdbsscjazvamyxmh";
NullableResponse<PublicCloudConnectorResource> response = await collection.GetIfExistsAsync(publicCloudConnector);
PublicCloudConnectorResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PublicCloudConnectorData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
