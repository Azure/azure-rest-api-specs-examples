using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.SecretsStoreExtension.Models;
using Azure.ResourceManager.SecretsStoreExtension;

// Generated from example definition: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "AzureKeyVaultSecretProviderClass_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg-ssc-example";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this KeyVaultSecretProviderClassResource
KeyVaultSecretProviderClassCollection collection = resourceGroupResource.GetKeyVaultSecretProviderClasses();

// invoke the operation
string azureKeyVaultSecretProviderClassName = "akvspc-ssc-example";
KeyVaultSecretProviderClassData data = new KeyVaultSecretProviderClassData(new AzureLocation("eastus"))
{
    Properties = new KeyVaultSecretProviderClassProperties("example-ssc-key-vault", Guid.Parse("00000000-0000-0000-0000-000000000000"), Guid.Parse("00000000-0000-0000-0000-000000000000"))
    {
        Objects = "array: |\n  - |\n    objectName: my-secret-object\n    objectType: secret\n    objectVersionHistory: 1",
    },
    ExtendedLocation = new ExtendedLocation
    {
        Name = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-ssc-example/providers/Microsoft.ExtendedLocation/customLocations/example-custom-location",
    },
    Tags =
    {
    ["example-tag"] = "example-tag-value"
    },
};
ArmOperation<KeyVaultSecretProviderClassResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, azureKeyVaultSecretProviderClassName, data);
KeyVaultSecretProviderClassResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
KeyVaultSecretProviderClassData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
