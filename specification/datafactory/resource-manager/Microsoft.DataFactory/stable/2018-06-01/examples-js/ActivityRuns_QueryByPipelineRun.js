const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Query activity runs based on input filter conditions.
 *
 * @summary Query activity runs based on input filter conditions.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ActivityRuns_QueryByPipelineRun.json
 */
async function activityRunsQueryByPipelineRun() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const runId = "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b";
  const filterParameters = {
    lastUpdatedAfter: new Date("2018-06-16T00:36:44.3345758Z"),
    lastUpdatedBefore: new Date("2018-06-16T00:49:48.3686473Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.activityRuns.queryByPipelineRun(
    resourceGroupName,
    factoryName,
    runId,
    filterParameters
  );
  console.log(result);
}

activityRunsQueryByPipelineRun().catch(console.error);
