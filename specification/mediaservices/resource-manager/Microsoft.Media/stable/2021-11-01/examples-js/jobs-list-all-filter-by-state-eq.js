const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the Jobs for the Transform.
 *
 * @summary Lists all of the Jobs for the Transform.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-list-all-filter-by-state-eq.json
 */
async function listsJobsForTheTransformFilterByStateEqual() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contosoresources";
  const accountName = "contosomedia";
  const transformName = "exampleTransform";
  const filter = "properties/state eq Microsoft.Media.JobState'Processing'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.list(resourceGroupName, accountName, transformName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsJobsForTheTransformFilterByStateEqual().catch(console.error);
