const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an integration runtime.
 *
 * @summary Creates or updates an integration runtime.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_Create.json
 */
async function integrationRuntimesCreate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const integrationRuntime = {
    properties: {
      type: "SelfHosted",
      description: "A selfhosted integration runtime",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.createOrUpdate(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    integrationRuntime
  );
  console.log(result);
}

integrationRuntimesCreate().catch(console.error);
