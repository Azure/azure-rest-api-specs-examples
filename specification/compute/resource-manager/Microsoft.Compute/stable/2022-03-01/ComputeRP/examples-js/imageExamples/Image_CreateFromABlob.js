const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an image.
 *
 * @summary Create or update an image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/imageExamples/Image_CreateFromABlob.json
 */
async function createAVirtualMachineImageFromABlob() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const imageName = "myImage";
  const parameters = {
    location: "West US",
    storageProfile: {
      osDisk: {
        blobUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
        osState: "Generalized",
        osType: "Linux",
      },
      zoneResilient: true,
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

createAVirtualMachineImageFromABlob().catch(console.error);
