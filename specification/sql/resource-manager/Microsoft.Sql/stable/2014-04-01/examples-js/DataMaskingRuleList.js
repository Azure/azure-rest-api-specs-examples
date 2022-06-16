const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of database data masking rules.
 *
 * @summary Gets a list of database data masking rules.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/DataMaskingRuleList.json
 */
async function listDataMaskingRules() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-6852";
  const serverName = "sqlcrudtest-2080";
  const databaseName = "sqlcrudtest-331";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dataMaskingRules.listByDatabase(
    resourceGroupName,
    serverName,
    databaseName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDataMaskingRules().catch(console.error);
