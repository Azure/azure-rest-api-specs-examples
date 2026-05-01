const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an image.
 *
 * @summary create or update an image.
 * x-ms-original-file: 2025-11-01/imageExamples/Image_CreateFromASnapshot.json
 */
async function createAVirtualMachineImageFromASnapshot() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.images.createOrUpdate("myResourceGroup", "myImage", {
    location: "West US",
    storageProfile: {
      osDisk: {
        osType: "Linux",
        snapshot: {
          id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
        },
        osState: "Generalized",
      },
      zoneResilient: false,
    },
  });
  console.log(result);
}
