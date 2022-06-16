const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all historical availability transitions and impacting events for a single resource. Use the nextLink property in the response to get the next page of availability status
 *
 * @summary Lists all historical availability transitions and impacting events for a single resource. Use the nextLink property in the response to get the next page of availability status
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2017-07-01/examples/AvailabilityStatuses_List.json
 */
async function getHealthHistoryByResource() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri = "resourceUri";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilityStatuses.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getHealthHistoryByResource().catch(console.error);
