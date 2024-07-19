using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountGetMigrationFailed.json
// this example is just showing the usage of "StorageAccounts_GetCustomerInitiatedMigration" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "resource-group-name";
string accountName = "accountname";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// get the collection of this StorageAccountMigrationResource
StorageAccountMigrationCollection collection = storageAccount.GetStorageAccountMigrations();

// invoke the operation
StorageAccountMigrationName migrationName = StorageAccountMigrationName.Default;
NullableResponse<StorageAccountMigrationResource> response = await collection.GetIfExistsAsync(migrationName);
StorageAccountMigrationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StorageAccountMigrationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
