const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create an integration runtime
 *
 * @summary Create an integration runtime
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Create.json
 */
async function createIntegrationRuntime() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const workspaceName = "exampleWorkspace";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const integrationRuntime = {
    properties: {
      type: "SelfHosted",
      description: "A selfhosted integration runtime",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.beginCreateAndWait(
    resourceGroupName,
    workspaceName,
    integrationRuntimeName,
    integrationRuntime
  );
  console.log(result);
}

createIntegrationRuntime().catch(console.error);
