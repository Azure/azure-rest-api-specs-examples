const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists available operations for the Microsoft.Batch provider
 *
 * @summary Lists available operations for the Microsoft.Batch provider
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/OperationsList.json
 */
async function operationsList() {
  const subscriptionId =
    process.env["BATCH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
