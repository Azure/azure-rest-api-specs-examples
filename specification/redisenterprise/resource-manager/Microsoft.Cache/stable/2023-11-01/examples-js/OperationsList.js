const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available REST API operations of the Microsoft.Cache provider.
 *
 * @summary Lists all of the available REST API operations of the Microsoft.Cache provider.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/OperationsList.json
 */
async function operationsList() {
  const subscriptionId =
    process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
