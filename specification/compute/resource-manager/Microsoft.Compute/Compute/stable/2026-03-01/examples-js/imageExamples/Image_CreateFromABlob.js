const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an image.
 *
 * @summary create or update an image.
 * x-ms-original-file: 2026-03-01/imageExamples/Image_CreateFromABlob.json
 */
async function createAVirtualMachineImageFromABlob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.images.createOrUpdate("myResourceGroup", "myImage", {
    location: "West US",
    storageProfile: {
      osDisk: {
        osType: "Linux",
        blobUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
        osState: "Generalized",
      },
      zoneResilient: true,
    },
  });
  console.log(result);
}
