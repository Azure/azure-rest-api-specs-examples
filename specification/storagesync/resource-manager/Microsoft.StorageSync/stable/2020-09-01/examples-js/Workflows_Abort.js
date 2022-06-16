const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Abort the given workflow.
 *
 * @summary Abort the given workflow.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/Workflows_Abort.json
 */
async function workflowsAbort() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const workflowId = "7ffd50b3-5574-478d-9ff2-9371bc42ce68";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.workflows.abort(
    resourceGroupName,
    storageSyncServiceName,
    workflowId
  );
  console.log(result);
}

workflowsAbort().catch(console.error);
