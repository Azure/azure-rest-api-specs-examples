const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an image.
 *
 * @summary create or update an image.
 * x-ms-original-file: 2025-11-01/imageExamples/Image_Create_DataDiskFromABlobIncluded.json
 */
async function createAVirtualMachineImageThatIncludesADataDiskFromABlob() {
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
      dataDisks: [
        {
          lun: 1,
          blobUri: "https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd",
        },
      ],
      zoneResilient: false,
    },
  });
  console.log(result);
}
