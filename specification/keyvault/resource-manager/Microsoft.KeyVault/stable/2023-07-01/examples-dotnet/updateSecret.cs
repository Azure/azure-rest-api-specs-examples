using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.KeyVault.Models;
using Azure.ResourceManager.KeyVault;

// Generated from example definition: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/updateSecret.json
// this example is just showing the usage of "Secrets_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KeyVaultSecretResource created on azure
// for more information of creating KeyVaultSecretResource, please refer to the document of KeyVaultSecretResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "sample-group";
string vaultName = "sample-vault";
string secretName = "secret-name";
ResourceIdentifier keyVaultSecretResourceId = KeyVaultSecretResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, secretName);
KeyVaultSecretResource keyVaultSecret = client.GetKeyVaultSecretResource(keyVaultSecretResourceId);

// invoke the operation
KeyVaultSecretPatch patch = new KeyVaultSecretPatch
{
    Properties = new SecretPatchProperties
    {
        Value = "secret-value2",
    },
};
KeyVaultSecretResource result = await keyVaultSecret.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KeyVaultSecretData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
