const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of operations performed on the elastic pool.
 *
 * @summary Gets a list of operations performed on the elastic pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListElasticPoolOperations.json
 */
async function listTheElasticPoolManagementOperations() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtestgroup";
  const serverName = "sqlcrudtestserver";
  const elasticPoolName = "testpool";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.elasticPoolOperations.listByElasticPool(
    resourceGroupName,
    serverName,
    elasticPoolName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTheElasticPoolManagementOperations().catch(console.error);
