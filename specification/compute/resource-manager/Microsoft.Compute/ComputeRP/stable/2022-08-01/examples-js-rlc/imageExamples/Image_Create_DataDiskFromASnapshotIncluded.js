const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update an image.
 *
 * @summary Create or update an image.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/imageExamples/Image_Create_DataDiskFromASnapshotIncluded.json
 */
async function createAVirtualMachineImageThatIncludesADataDiskFromASnapshot() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const imageName = "myImage";
  const options = {
    body: {
      location: "West US",
      properties: {
        storageProfile: {
          dataDisks: [
            {
              lun: 1,
              snapshot: {
                id: "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2",
              },
            },
          ],
          osDisk: {
            osState: "Generalized",
            osType: "Linux",
            snapshot: {
              id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
            },
          },
          zoneResilient: true,
        },
      },
    },
    queryParameters: { "api-version": "2022-08-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}",
      subscriptionId,
      resourceGroupName,
      imageName,
    )
    .put(options);
  const poller = await getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createAVirtualMachineImageThatIncludesADataDiskFromASnapshot().catch(console.error);
