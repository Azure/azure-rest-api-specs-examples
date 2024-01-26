const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove all linked integration runtimes under specific data factory in a self-hosted integration runtime.
 *
 * @summary Remove all linked integration runtimes under specific data factory in a self-hosted integration runtime.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_RemoveLinks.json
 */
async function integrationRuntimesUpgrade() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const linkedIntegrationRuntimeRequest = {
    linkedFactoryName: "exampleFactoryName-linked",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.removeLinks(
    resourceGroupName,
    factoryName,
    integrationRuntimeName,
    linkedIntegrationRuntimeRequest,
  );
  console.log(result);
}
