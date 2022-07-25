const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the integration runtime status
 *
 * @summary Get the integration runtime status
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_GetStatus.json
 */
async function getStatus() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const workspaceName = "exampleWorkspace";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimeStatusOperations.get(
    resourceGroupName,
    workspaceName,
    integrationRuntimeName
  );
  console.log(result);
}

getStatus().catch(console.error);
