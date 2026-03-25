const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new pool inside the specified account.
 *
 * @summary creates a new pool inside the specified account.
 * x-ms-original-file: 2025-06-01/PoolCreate_SharedImageGallery.json
 */
async function createPoolCustomImage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.pool.create(
    "default-azurebatch-japaneast",
    "sampleacct",
    "testpool",
    {
      deploymentConfiguration: {
        virtualMachineConfiguration: {
          imageReference: {
            id: "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/networking-group/providers/Microsoft.Compute/galleries/testgallery/images/testimagedef/versions/0.0.1",
          },
          nodeAgentSkuId: "batch.node.ubuntu 18.04",
        },
      },
      vmSize: "STANDARD_D4",
    },
  );
  console.log(result);
}
