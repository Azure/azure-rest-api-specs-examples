const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an integration runtime.
 *
 * @summary Updates an integration runtime.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Update.json
 */
async function integrationRuntimesUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const updateIntegrationRuntimeRequest = {
    autoUpdate: "Off",
    updateDelayOffset: '"PT3H"',
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.update(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    updateIntegrationRuntimeRequest
  );
  console.log(result);
}

integrationRuntimesUpdate().catch(console.error);
