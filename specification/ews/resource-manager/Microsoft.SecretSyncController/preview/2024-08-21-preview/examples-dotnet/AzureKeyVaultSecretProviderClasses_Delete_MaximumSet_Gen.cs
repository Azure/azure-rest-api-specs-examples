using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecretsStoreExtension.Models;
using Azure.ResourceManager.SecretsStoreExtension;

// Generated from example definition: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "AzureKeyVaultSecretProviderClass_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KeyVaultSecretProviderClassResource created on azure
// for more information of creating KeyVaultSecretProviderClassResource, please refer to the document of KeyVaultSecretProviderClassResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg-ssc-example";
string azureKeyVaultSecretProviderClassName = "akvspc-ssc-example";
ResourceIdentifier keyVaultSecretProviderClassResourceId = KeyVaultSecretProviderClassResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureKeyVaultSecretProviderClassName);
KeyVaultSecretProviderClassResource keyVaultSecretProviderClass = client.GetKeyVaultSecretProviderClassResource(keyVaultSecretProviderClassResourceId);

// invoke the operation
await keyVaultSecretProviderClass.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
