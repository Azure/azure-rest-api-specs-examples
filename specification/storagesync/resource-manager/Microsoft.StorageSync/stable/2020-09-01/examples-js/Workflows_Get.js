const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Workflows resource
 *
 * @summary Get Workflows resource
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/Workflows_Get.json
 */
async function workflowsGet() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const workflowId = "828219ea-083e-48b5-89ea-8fd9991b2e75";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.workflows.get(resourceGroupName, storageSyncServiceName, workflowId);
  console.log(result);
}

workflowsGet().catch(console.error);
