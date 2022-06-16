const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new pool inside the specified account.
 *
 * @summary Creates a new pool inside the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolCreate_NoPublicIPAddresses.json
 */
async function createPoolNoPublicIP() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const parameters = {
    deploymentConfiguration: {
      virtualMachineConfiguration: {
        imageReference: {
          id: "/subscriptions/subid/resourceGroups/networking-group/providers/Microsoft.Compute/galleries/testgallery/images/testimagedef/versions/0.0.1",
        },
        nodeAgentSkuId: "batch.node.ubuntu 18.04",
      },
    },
    networkConfiguration: {
      publicIPAddressConfiguration: { provision: "NoPublicIPAddresses" },
      subnetId:
        "/subscriptions/subid/resourceGroups/rg1234/providers/Microsoft.Network/virtualNetworks/network1234/subnets/subnet123",
    },
    vmSize: "STANDARD_D4",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.create(
    resourceGroupName,
    accountName,
    poolName,
    parameters
  );
  console.log(result);
}

createPoolNoPublicIP().catch(console.error);
