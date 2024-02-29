const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one or two properties) an input without affecting the rest the job or input definition.
 *
 * @summary Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one or two properties) an input without affecting the rest the job or input definition.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Update_Stream_Blob.json
 */
async function updateAStreamBlobInput() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg8161";
  const jobName = "sj6695";
  const inputName = "input8899";
  const input = {
    properties: {
      type: "Stream",
      datasource: { type: "Microsoft.Storage/Blob", sourcePartitionCount: 32 },
      serialization: { type: "Csv", encoding: "UTF8", fieldDelimiter: "|" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.update(resourceGroupName, jobName, inputName, input);
  console.log(result);
}
