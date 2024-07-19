using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountPatchEncryptionScope.json
// this example is just showing the usage of "EncryptionScopes_Patch" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EncryptionScopeResource created on azure
// for more information of creating EncryptionScopeResource, please refer to the document of EncryptionScopeResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "resource-group-name";
string accountName = "accountname";
string encryptionScopeName = "{encryption-scope-name}";
ResourceIdentifier encryptionScopeResourceId = EncryptionScopeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, encryptionScopeName);
EncryptionScopeResource encryptionScope = client.GetEncryptionScopeResource(encryptionScopeResourceId);

// invoke the operation
EncryptionScopeData data = new EncryptionScopeData()
{
    Source = EncryptionScopeSource.KeyVault,
    KeyVaultProperties = new EncryptionScopeKeyVaultProperties()
    {
        KeyUri = new Uri("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
    },
};
EncryptionScopeResource result = await encryptionScope.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EncryptionScopeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
