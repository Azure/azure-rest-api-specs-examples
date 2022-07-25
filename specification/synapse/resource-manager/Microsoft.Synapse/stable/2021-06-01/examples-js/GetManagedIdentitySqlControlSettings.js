const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Managed Identity Sql Control Settings
 *
 * @summary Get Managed Identity Sql Control Settings
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetManagedIdentitySqlControlSettings.json
 */
async function getManagedIdentitySqlControlSettings() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "resourceGroup1";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceManagedIdentitySqlControlSettings.get(
    resourceGroupName,
    workspaceName
  );
  console.log(result);
}

getManagedIdentitySqlControlSettings().catch(console.error);
