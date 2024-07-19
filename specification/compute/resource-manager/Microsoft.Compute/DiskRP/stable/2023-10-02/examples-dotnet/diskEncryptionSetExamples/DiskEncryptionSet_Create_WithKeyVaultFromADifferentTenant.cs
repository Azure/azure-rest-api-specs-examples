using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
// this example is just showing the usage of "DiskEncryptionSets_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DiskEncryptionSetResource
DiskEncryptionSetCollection collection = resourceGroupResource.GetDiskEncryptionSets();

// invoke the operation
string diskEncryptionSetName = "myDiskEncryptionSet";
DiskEncryptionSetData data = new DiskEncryptionSetData(new AzureLocation("West US"))
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}")] = new UserAssignedIdentity(),
        },
    },
    EncryptionType = DiskEncryptionSetType.EncryptionAtRestWithCustomerKey,
    ActiveKey = new KeyForDiskEncryptionSet(new Uri("https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}")),
    FederatedClientId = "00000000-0000-0000-0000-000000000000",
};
ArmOperation<DiskEncryptionSetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, diskEncryptionSetName, data);
DiskEncryptionSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiskEncryptionSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
