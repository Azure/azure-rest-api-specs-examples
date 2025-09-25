using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/BlobContainersGet.json
// this example is just showing the usage of "BlobContainers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlobServiceResource created on azure
// for more information of creating BlobServiceResource, please refer to the document of BlobServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res9871";
string accountName = "sto6217";
ResourceIdentifier blobServiceResourceId = BlobServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BlobServiceResource blobService = client.GetBlobServiceResource(blobServiceResourceId);

// get the collection of this BlobContainerResource
BlobContainerCollection collection = blobService.GetBlobContainers();

// invoke the operation
string containerName = "container1634";
NullableResponse<BlobContainerResource> response = await collection.GetIfExistsAsync(containerName);
BlobContainerResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BlobContainerData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
