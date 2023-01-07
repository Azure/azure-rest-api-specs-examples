const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists pipelines.
 *
 * @summary Lists pipelines.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_ListByFactory.json
 */
async function pipelinesListByFactory() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.pipelines.listByFactory(resourceGroupName, factoryName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
