const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Puts new sql agent configuration to instance.
 *
 * @summary Puts new sql agent configuration to instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SqlAgentConfigurationPut.json
 */
async function putsNewSqlAgentConfigurationToInstance() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const managedInstanceName = "sqlcrudtest-4645";
  const parameters = { state: "Enabled" };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.sqlAgent.createOrUpdate(
    resourceGroupName,
    managedInstanceName,
    parameters
  );
  console.log(result);
}

putsNewSqlAgentConfigurationToInstance().catch(console.error);
