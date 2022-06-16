const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the IP address of self-hosted integration runtime node.
 *
 * @summary Get the IP address of self-hosted integration runtime node.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimeNodes_GetIpAddress.json
 */
async function integrationRuntimeNodesGetIPAddress() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const nodeName = "Node_1";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimeNodes.getIpAddress(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    nodeName
  );
  console.log(result);
}

integrationRuntimeNodesGetIPAddress().catch(console.error);
