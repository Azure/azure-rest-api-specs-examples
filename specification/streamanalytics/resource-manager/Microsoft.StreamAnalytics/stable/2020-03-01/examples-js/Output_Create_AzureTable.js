const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an output or replaces an already existing output under an existing streaming job.
 *
 * @summary Creates an output or replaces an already existing output under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_AzureTable.json
 */
async function createAnAzureTableOutput() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg5176";
  const jobName = "sj2790";
  const outputName = "output958";
  const output = {
    datasource: {
      type: "Microsoft.Storage/Table",
      accountKey: "accountKey==",
      accountName: "someAccountName",
      batchSize: 25,
      columnsToRemove: ["column1", "column2"],
      partitionKey: "partitionKey",
      rowKey: "rowKey",
      table: "samples",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.createOrReplace(
    resourceGroupName,
    jobName,
    outputName,
    output
  );
  console.log(result);
}

createAnAzureTableOutput().catch(console.error);
