const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Query triggers.
 *
 * @summary Query triggers.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_QueryByFactory.json
 */
async function triggersQueryByFactory() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const filterParameters = {
    parentTriggerName: "exampleTrigger",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.triggers.queryByFactory(
    resourceGroupName,
    factoryName,
    filterParameters
  );
  console.log(result);
}

triggersQueryByFactory().catch(console.error);
