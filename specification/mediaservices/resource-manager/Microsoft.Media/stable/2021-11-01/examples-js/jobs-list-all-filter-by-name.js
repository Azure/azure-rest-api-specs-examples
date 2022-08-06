const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the Jobs for the Transform.
 *
 * @summary Lists all of the Jobs for the Transform.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-list-all-filter-by-name.json
 */
async function listsJobsForTheTransformFilterByName() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contosoresources";
  const accountName = "contosomedia";
  const transformName = "exampleTransform";
  const filter = "name eq 'job1' or name eq 'job2'";
  const orderby = "name";
  const options = { filter, orderby };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.list(resourceGroupName, accountName, transformName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsJobsForTheTransformFilterByName().catch(console.error);
