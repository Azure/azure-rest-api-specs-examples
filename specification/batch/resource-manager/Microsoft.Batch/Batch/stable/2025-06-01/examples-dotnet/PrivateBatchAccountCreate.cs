using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/PrivateBatchAccountCreate.json
// this example is just showing the usage of "BatchAccount_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
string resourceGroupName = "default-azurebatch-japaneast";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this BatchAccountResource
BatchAccountCollection collection = resourceGroupResource.GetBatchAccounts();

// invoke the operation
string accountName = "sampleacct";
BatchAccountCreateOrUpdateContent content = new BatchAccountCreateOrUpdateContent(new AzureLocation("japaneast"))
{
    AutoStorage = new BatchAccountAutoStorageBaseConfiguration(new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage")),
    KeyVaultReference = new BatchKeyVaultReference(new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.KeyVault/vaults/sample"), new Uri("http://sample.vault.azure.net/")),
    PublicNetworkAccess = BatchPublicNetworkAccess.Disabled,
};
ArmOperation<BatchAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, content);
BatchAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BatchAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
