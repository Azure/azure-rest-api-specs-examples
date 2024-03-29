const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an image.
 *
 * @summary Create or update an image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/imageExamples/Image_CreateFromABlobWithDiskEncryptionSet.json
 */
async function createAVirtualMachineImageFromABlobWithDiskEncryptionSetResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageName = "myImage";
  const parameters = {
    location: "West US",
    storageProfile: {
      osDisk: {
        blobUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
        diskEncryptionSet: {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
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

createAVirtualMachineImageFromABlobWithDiskEncryptionSetResource().catch(console.error);
