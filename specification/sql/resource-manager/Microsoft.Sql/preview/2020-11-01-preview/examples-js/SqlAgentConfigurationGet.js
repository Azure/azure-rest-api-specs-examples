const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets current instance sql agent configuration.
 *
 * @summary Gets current instance sql agent configuration.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SqlAgentConfigurationGet.json
 */
async function getsCurrentInstanceSqlAgentConfiguration() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.sqlAgent.get(resourceGroupName, managedInstanceName);
  console.log(result);
}

getsCurrentInstanceSqlAgentConfiguration().catch(console.error);
