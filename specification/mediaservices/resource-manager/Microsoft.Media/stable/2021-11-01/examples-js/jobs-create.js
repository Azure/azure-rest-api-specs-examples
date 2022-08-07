const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Job.
 *
 * @summary Creates a Job.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-create.json
 */
async function createAJob() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contosoresources";
  const accountName = "contosomedia";
  const transformName = "exampleTransform";
  const jobName = "job1";
  const parameters = {
    correlationData: { key2: "Value 2", key1: "value1" },
    input: {
      odataType: "#Microsoft.Media.JobInputAsset",
      assetName: "job1-InputAsset",
    },
    outputs: [
      {
        odataType: "#Microsoft.Media.JobOutputAsset",
        assetName: "job1-OutputAsset",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.jobs.create(
    resourceGroupName,
    accountName,
    transformName,
    jobName,
    parameters
  );
  console.log(result);
}

createAJob().catch(console.error);
