const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns activity on databases inside of an elastic pool.
 *
 * @summary Returns activity on databases inside of an elastic pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01-legacy/examples/ElasticPoolDatabaseActivityList.json
 */
async function listElasticPoolDatabaseActivity() {
  const subscriptionId = "9d4e2ad0-e20b-4464-9219-353bded52513";
  const resourceGroupName = "sqlcrudtest-4673";
  const serverName = "sqlcrudtest-603";
  const elasticPoolName = "7537";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.elasticPoolDatabaseActivities.listByElasticPool(
    resourceGroupName,
    serverName,
    elasticPoolName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listElasticPoolDatabaseActivity().catch(console.error);
