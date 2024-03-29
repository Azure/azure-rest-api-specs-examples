const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the configurations in a given server.
 *
 * @summary List all the configurations in a given server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsListByServer.json
 */
async function listAllConfigurationsForAServer() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "mysqltestserver";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurations.listByServer(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
