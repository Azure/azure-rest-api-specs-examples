using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/BlobContainersSetLegalHoldAllowProtectedAppendWritesAll.json
// this example is just showing the usage of "BlobContainers_SetLegalHold" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlobContainerResource created on azure
// for more information of creating BlobContainerResource, please refer to the document of BlobContainerResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res4303";
string accountName = "sto7280";
string containerName = "container8723";
ResourceIdentifier blobContainerResourceId = BlobContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, containerName);
BlobContainerResource blobContainer = client.GetBlobContainerResource(blobContainerResourceId);

// invoke the operation
LegalHold legalHold = new LegalHold(new string[] { "tag1", "tag2", "tag3" })
{
    AllowProtectedAppendWritesAll = true,
};
LegalHold result = await blobContainer.SetLegalHoldAsync(legalHold);

Console.WriteLine($"Succeeded: {result}");
