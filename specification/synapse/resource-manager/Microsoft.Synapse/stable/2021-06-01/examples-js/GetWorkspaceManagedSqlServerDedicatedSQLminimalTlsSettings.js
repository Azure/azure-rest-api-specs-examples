const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get workspace managed sql server's minimal tls settings.
 *
 * @summary Get workspace managed sql server's minimal tls settings.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings.json
 */
async function getWorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workspace-6852";
  const workspaceName = "workspace-2080";
  const dedicatedSQLminimalTlsSettingsName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceManagedSqlServerDedicatedSQLMinimalTlsSettings.get(
    resourceGroupName,
    workspaceName,
    dedicatedSQLminimalTlsSettingsName
  );
  console.log(result);
}

getWorkspaceManagedSqlServerDedicatedSqlMinimalTlsSettings().catch(console.error);
