const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all historical availability transitions and impacting events for a single resource.
 *
 * @summary Lists all historical availability transitions and impacting events for a single resource.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/AvailabilityStatuses_List.json
 */
async function getHealthHistoryByResource() {
  const resourceUri = "resourceUri";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const resArray = new Array();
  for await (let item of client.availabilityStatuses.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
