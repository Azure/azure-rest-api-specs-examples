const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list all private scheduled actions.
 *
 * @summary list all private scheduled actions.
 * x-ms-original-file: 2025-03-01/scheduledActions/scheduledActions-list-private.json
 */
async function privateScheduledActionsList() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.scheduledActions.list()) {
    resArray.push(item);
  }

  console.log(resArray);
}
