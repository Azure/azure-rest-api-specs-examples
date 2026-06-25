const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an image.
 *
 * @summary create or update an image.
 * x-ms-original-file: 2026-03-01/imageExamples/Image_CreateFromABlobWithDiskEncryptionSet.json
 */
async function createAVirtualMachineImageFromABlobWithDiskEncryptionSetResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.images.createOrUpdate("myResourceGroup", "myImage", {
    location: "West US",
    storageProfile: {
      osDisk: {
        osType: "Linux",
        blobUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
        diskEncryptionSet: {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
        },
        osState: "Generalized",
      },
    },
  });
  console.log(result);
}
