const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified pool.
 *
 * @summary Gets information about the specified pool.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolGet_VirtualMachineConfiguration_Extensions.json
 */
async function getPoolVirtualMachineConfigurationExtensions() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const poolName = "testpool";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.poolOperations.get(resourceGroupName, accountName, poolName);
  console.log(result);
}

getPoolVirtualMachineConfigurationExtensions().catch(console.error);
