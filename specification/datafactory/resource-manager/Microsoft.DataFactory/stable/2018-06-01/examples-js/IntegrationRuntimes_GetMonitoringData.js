const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the integration runtime monitoring data, which includes the monitor data for all the nodes under this integration runtime.
 *
 * @summary Get the integration runtime monitoring data, which includes the monitor data for all the nodes under this integration runtime.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_GetMonitoringData.json
 */
async function integrationRuntimesGetMonitoringData() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.getMonitoringData(
    resourceGroupName,
    factoryName,
    integrationRuntimeName
  );
  console.log(result);
}

integrationRuntimesGetMonitoringData().catch(console.error);
