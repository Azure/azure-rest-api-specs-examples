const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a self-hosted integration runtime node.
 *
 * @summary Updates a self-hosted integration runtime node.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimeNodes_Update.json
 */
async function integrationRuntimeNodesUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const nodeName = "Node_1";
  const updateIntegrationRuntimeNodeRequest = {
    concurrentJobsLimit: 2,
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimeNodes.update(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    nodeName,
    updateIntegrationRuntimeNodeRequest
  );
  console.log(result);
}

integrationRuntimeNodesUpdate().catch(console.error);
