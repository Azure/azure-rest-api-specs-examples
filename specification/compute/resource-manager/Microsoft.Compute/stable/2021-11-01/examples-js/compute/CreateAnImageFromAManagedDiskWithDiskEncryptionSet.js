const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAVirtualMachineImageFromAManagedDiskWithDiskEncryptionSetResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageName = "myImage";
  const parameters = {
    location: "West US",
    storageProfile: {
      osDisk: {
        diskEncryptionSet: {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
        },
        managedDisk: {
          id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk",
        },
        osState: "Generalized",
        osType: "Linux",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.images.beginCreateOrUpdateAndWait(
    resourceGroupName,
    imageName,
    parameters
  );
  console.log(result);
}

createAVirtualMachineImageFromAManagedDiskWithDiskEncryptionSetResource().catch(console.error);
