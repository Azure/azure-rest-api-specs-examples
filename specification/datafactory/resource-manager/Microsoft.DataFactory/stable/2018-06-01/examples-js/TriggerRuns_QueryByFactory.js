const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Query trigger runs.
 *
 * @summary Query trigger runs.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/TriggerRuns_QueryByFactory.json
 */
async function triggerRunsQueryByFactory() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const filterParameters = {
    filters: [{ operand: "TriggerName", operator: "Equals", values: ["exampleTrigger"] }],
    lastUpdatedAfter: new Date("2018-06-16T00:36:44.3345758Z"),
    lastUpdatedBefore: new Date("2018-06-16T00:49:48.3686473Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.triggerRuns.queryByFactory(
    resourceGroupName,
    factoryName,
    filterParameters
  );
  console.log(result);
}

triggerRunsQueryByFactory().catch(console.error);
