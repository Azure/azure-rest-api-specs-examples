const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Upgrade an integration runtime
 *
 * @summary Upgrade an integration runtime
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Upgrade.json
 */
async function upgradeIntegrationRuntime() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const workspaceName = "exampleWorkspace";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.upgrade(
    resourceGroupName,
    workspaceName,
    integrationRuntimeName
  );
  console.log(result);
}

upgradeIntegrationRuntime().catch(console.error);
