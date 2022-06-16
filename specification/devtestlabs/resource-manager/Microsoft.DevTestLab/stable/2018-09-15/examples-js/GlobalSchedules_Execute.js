const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Execute a schedule. This operation can take a while to complete.
 *
 * @summary Execute a schedule. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_Execute.json
 */
async function globalSchedulesExecute() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "labvmautostart";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.globalSchedules.beginExecuteAndWait(resourceGroupName, name);
  console.log(result);
}

globalSchedulesExecute().catch(console.error);
