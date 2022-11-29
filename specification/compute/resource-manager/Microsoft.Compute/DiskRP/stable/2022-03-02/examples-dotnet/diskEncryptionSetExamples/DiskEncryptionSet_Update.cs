using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Update.json
// this example is just showing the usage of "DiskEncryptionSets_Update" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this DiskEncryptionSetResource created on azure
// for more information of creating DiskEncryptionSetResource, please refer to the document of DiskEncryptionSetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string diskEncryptionSetName = "myDiskEncryptionSet";
ResourceIdentifier diskEncryptionSetResourceId = DiskEncryptionSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskEncryptionSetName);
DiskEncryptionSetResource diskEncryptionSet = client.GetDiskEncryptionSetResource(diskEncryptionSetResourceId);

// invoke the operation
DiskEncryptionSetPatch patch = new DiskEncryptionSetPatch()
{
    Tags =
    {
    ["department"] = "Development",
    ["project"] = "Encryption",
    },
    EncryptionType = DiskEncryptionSetType.EncryptionAtRestWithCustomerKey,
    ActiveKey = new KeyForDiskEncryptionSet(new Uri("https://myvmvault.vault-int.azure-int.net/keys/keyName/keyVersion"))
    {
        SourceVaultId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
    },
};
ArmOperation<DiskEncryptionSetResource> lro = await diskEncryptionSet.UpdateAsync(WaitUntil.Completed, patch);
DiskEncryptionSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiskEncryptionSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
