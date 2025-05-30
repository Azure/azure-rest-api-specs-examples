using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.KeyVault.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.KeyVault;

// Generated from example definition: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/createVault.json
// this example is just showing the usage of "Vaults_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "sample-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this KeyVaultResource
KeyVaultCollection collection = resourceGroupResource.GetKeyVaults();

// invoke the operation
string vaultName = "sample-vault";
KeyVaultCreateOrUpdateContent content = new KeyVaultCreateOrUpdateContent(new AzureLocation("westus"), new KeyVaultProperties(Guid.Parse("00000000-0000-0000-0000-000000000000"), new KeyVaultSku(KeyVaultSkuFamily.A, KeyVaultSkuName.Standard))
{
    AccessPolicies = {new KeyVaultAccessPolicy(Guid.Parse("00000000-0000-0000-0000-000000000000"), "00000000-0000-0000-0000-000000000000", new IdentityAccessPermissions
    {
    Keys = {IdentityAccessKeyPermission.Encrypt, IdentityAccessKeyPermission.Decrypt, IdentityAccessKeyPermission.WrapKey, IdentityAccessKeyPermission.UnwrapKey, IdentityAccessKeyPermission.Sign, IdentityAccessKeyPermission.Verify, IdentityAccessKeyPermission.Get, IdentityAccessKeyPermission.List, IdentityAccessKeyPermission.Create, IdentityAccessKeyPermission.Update, IdentityAccessKeyPermission.Import, IdentityAccessKeyPermission.Delete, IdentityAccessKeyPermission.Backup, IdentityAccessKeyPermission.Restore, IdentityAccessKeyPermission.Recover, IdentityAccessKeyPermission.Purge},
    Secrets = {IdentityAccessSecretPermission.Get, IdentityAccessSecretPermission.List, IdentityAccessSecretPermission.Set, IdentityAccessSecretPermission.Delete, IdentityAccessSecretPermission.Backup, IdentityAccessSecretPermission.Restore, IdentityAccessSecretPermission.Recover, IdentityAccessSecretPermission.Purge},
    Certificates = {IdentityAccessCertificatePermission.Get, IdentityAccessCertificatePermission.List, IdentityAccessCertificatePermission.Delete, IdentityAccessCertificatePermission.Create, IdentityAccessCertificatePermission.Import, IdentityAccessCertificatePermission.Update, IdentityAccessCertificatePermission.ManageContacts, IdentityAccessCertificatePermission.GetIssuers, IdentityAccessCertificatePermission.ListIssuers, IdentityAccessCertificatePermission.SetIssuers, IdentityAccessCertificatePermission.DeleteIssuers, IdentityAccessCertificatePermission.ManageIssuers, IdentityAccessCertificatePermission.Recover, IdentityAccessCertificatePermission.Purge},
    })},
    EnabledForDeployment = true,
    EnabledForDiskEncryption = true,
    EnabledForTemplateDeployment = true,
    PublicNetworkAccess = "Enabled",
});
ArmOperation<KeyVaultResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vaultName, content);
KeyVaultResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KeyVaultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
