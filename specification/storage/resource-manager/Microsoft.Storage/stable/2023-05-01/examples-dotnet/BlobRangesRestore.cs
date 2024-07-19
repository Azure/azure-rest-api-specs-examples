using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/BlobRangesRestore.json
// this example is just showing the usage of "StorageAccounts_RestoreBlobRanges" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res9101";
string accountName = "sto4445";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// invoke the operation
BlobRestoreContent content = new BlobRestoreContent(DateTimeOffset.Parse("placeholder"), new BlobRestoreRange[]
{
new BlobRestoreRange("container/blobpath1","container/blobpath2"),new BlobRestoreRange("container2/blobpath3","")
});
ArmOperation<BlobRestoreStatus> lro = await storageAccount.RestoreBlobRangesAsync(WaitUntil.Completed, content);
BlobRestoreStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
