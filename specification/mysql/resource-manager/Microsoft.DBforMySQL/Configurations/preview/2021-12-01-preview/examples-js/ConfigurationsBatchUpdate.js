const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a list of configurations in a given server.
 *
 * @summary Update a list of configurations in a given server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/Configurations/preview/2021-12-01-preview/examples/ConfigurationsBatchUpdate.json
 */
async function configurationList() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["MYSQL_RESOURCE_GROUP"] || "testrg";
  const serverName = "mysqltestserver";
  const parameters = {
    resetAllToDefault: "False",
    value: [
      { name: "event_scheduler", value: "OFF" },
      { name: "div_precision_increment", value: "8" },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.configurations.beginBatchUpdateAndWait(
    resourceGroupName,
    serverName,
    parameters
  );
  console.log(result);
}
