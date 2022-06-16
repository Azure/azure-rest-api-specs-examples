const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Global parameter
 *
 * @summary Deletes a Global parameter
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Delete.json
 */
async function globalParametersDelete() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const globalParameterName = "default";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.globalParameters.delete(
    resourceGroupName,
    factoryName,
    globalParameterName
  );
  console.log(result);
}

globalParametersDelete().catch(console.error);
