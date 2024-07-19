using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/BlobContainersList.json
// this example is just showing the usage of "BlobContainers_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlobServiceResource created on azure
// for more information of creating BlobServiceResource, please refer to the document of BlobServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res9290";
string accountName = "sto1590";
ResourceIdentifier blobServiceResourceId = BlobServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
BlobServiceResource blobService = client.GetBlobServiceResource(blobServiceResourceId);

// get the collection of this BlobContainerResource
BlobContainerCollection collection = blobService.GetBlobContainers();

// invoke the operation and iterate over the result
await foreach (BlobContainerResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BlobContainerData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
