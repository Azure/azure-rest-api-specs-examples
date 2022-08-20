const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements resourcePool DELETE method.
 *
 * @summary Implements resourcePool DELETE method.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteResourcePool.json
 */
async function deleteResourcePool() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const resourcePoolName = "HRPool";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.resourcePools.beginDeleteAndWait(resourceGroupName, resourcePoolName);
  console.log(result);
}

deleteResourcePool().catch(console.error);
