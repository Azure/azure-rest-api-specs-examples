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

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountEnableCMK.json
// this example is just showing the usage of "StorageAccounts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res9407";
string accountName = "sto8596";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// invoke the operation
StorageAccountPatch patch = new StorageAccountPatch()
{
    Encryption = new StorageAccountEncryption()
    {
        Services = new StorageAccountEncryptionServices()
        {
            Blob = new StorageEncryptionService()
            {
                IsEnabled = true,
                KeyType = StorageEncryptionKeyType.Account,
            },
            File = new StorageEncryptionService()
            {
                IsEnabled = true,
                KeyType = StorageEncryptionKeyType.Account,
            },
        },
        KeySource = StorageAccountKeySource.KeyVault,
        KeyVaultProperties = new StorageAccountKeyVaultProperties()
        {
            KeyName = "wrappingKey",
            KeyVersion = "",
            KeyVaultUri = new Uri("https://myvault8569.vault.azure.net"),
        },
    },
};
StorageAccountResource result = await storageAccount.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
