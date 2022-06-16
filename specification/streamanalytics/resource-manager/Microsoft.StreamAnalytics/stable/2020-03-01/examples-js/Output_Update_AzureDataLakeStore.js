const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing output under an existing streaming job. This can be used to partially update (ie. update one or two properties) an output without affecting the rest the job or output definition.
 *
 * @summary Updates an existing output under an existing streaming job. This can be used to partially update (ie. update one or two properties) an output without affecting the rest the job or output definition.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_AzureDataLakeStore.json
 */
async function updateAnAzureDataLakeStoreOutputWithJsonSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6912";
  const jobName = "sj3310";
  const outputName = "output5195";
  const output = {
    datasource: {
      type: "Microsoft.DataLake/Accounts",
      accountName: "differentaccount",
    },
    serialization: { type: "Json", format: "LineSeparated", encoding: "UTF8" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.update(resourceGroupName, jobName, outputName, output);
  console.log(result);
}

updateAnAzureDataLakeStoreOutputWithJsonSerialization().catch(console.error);
