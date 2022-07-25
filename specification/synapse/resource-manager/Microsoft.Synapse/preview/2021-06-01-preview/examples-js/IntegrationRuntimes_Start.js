const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Start an integration runtime
 *
 * @summary Start an integration runtime
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Start.json
 */
async function startIntegrationRuntime() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const workspaceName = "exampleWorkspace";
  const integrationRuntimeName = "exampleManagedIntegrationRuntime";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.beginStartAndWait(
    resourceGroupName,
    workspaceName,
    integrationRuntimeName
  );
  console.log(result);
}

startIntegrationRuntime().catch(console.error);
