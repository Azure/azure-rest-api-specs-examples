const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the authentication keys for an integration runtime.
 *
 * @summary Retrieves the authentication keys for an integration runtime.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_ListAuthKeys.json
 */
async function integrationRuntimesListAuthKeys() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.listAuthKeys(
    resourceGroupName,
    factoryName,
    integrationRuntimeName
  );
  console.log(result);
}

integrationRuntimesListAuthKeys().catch(console.error);
