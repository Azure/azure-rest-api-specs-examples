const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List workspace managed sql server's minimal tls settings.
 *
 * @summary List workspace managed sql server's minimal tls settings.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings.json
 */
async function listDedicatedSqlMinimalTlsSettingsOfTheWorkspaceManagedSqlServer() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "workspace-6852";
  const workspaceName = "workspace-2080";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaceManagedSqlServerDedicatedSQLMinimalTlsSettings.list(
    resourceGroupName,
    workspaceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
