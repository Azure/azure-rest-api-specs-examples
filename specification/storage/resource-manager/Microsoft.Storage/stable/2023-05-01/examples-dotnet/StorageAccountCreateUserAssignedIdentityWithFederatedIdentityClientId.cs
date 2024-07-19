using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountCreateUserAssignedIdentityWithFederatedIdentityClientId.json
// this example is just showing the usage of "StorageAccounts_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res131918";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this StorageAccountResource
StorageAccountCollection collection = resourceGroupResource.GetStorageAccounts();

// invoke the operation
string accountName = "sto131918";
StorageAccountCreateOrUpdateContent content = new StorageAccountCreateOrUpdateContent(new StorageSku(StorageSkuName.StandardLrs), StorageKind.Storage, new AzureLocation("eastus"))
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}")] = new UserAssignedIdentity(),
        },
    },
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
        EncryptionIdentity = new StorageAccountEncryptionIdentity()
        {
            EncryptionUserAssignedIdentity = "/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{managed-identity-name}",
            EncryptionFederatedIdentityClientId = "f83c6b1b-4d34-47e4-bb34-9d83df58b540",
        },
    },
};
ArmOperation<StorageAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, content);
StorageAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StorageAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
