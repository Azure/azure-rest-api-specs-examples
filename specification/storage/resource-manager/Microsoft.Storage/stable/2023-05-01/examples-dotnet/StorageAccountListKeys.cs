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

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountListKeys.json
// this example is just showing the usage of "StorageAccounts_ListKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res418";
string accountName = "sto2220";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// invoke the operation and iterate over the result
await foreach (StorageAccountKey item in storageAccount.GetKeysAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
