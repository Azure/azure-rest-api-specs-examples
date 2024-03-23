const { StorageActionsManagementClient } = require("@azure/arm-storageactions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available Storage Actions Rest API operations.
 *
 * @summary Lists all of the available Storage Actions Rest API operations.
 * x-ms-original-file: specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/misc/OperationsList.json
 */
async function operationsList() {
  const subscriptionId =
    process.env["STORAGEACTIONS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new StorageActionsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
