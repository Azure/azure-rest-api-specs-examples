const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available REST API operations.
 *
 * @summary Lists all of the available REST API operations.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/OperationList.json
 */
async function operationList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.operations.list();
  console.log(result);
}

operationList().catch(console.error);
