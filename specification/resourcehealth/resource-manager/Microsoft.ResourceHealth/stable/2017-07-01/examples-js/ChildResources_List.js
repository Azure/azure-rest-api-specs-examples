const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the all the children and its current health status for a parent resource. Use the nextLink property in the response to get the next page of children current health
 *
 * @summary Lists the all the children and its current health status for a parent resource. Use the nextLink property in the response to get the next page of children current health
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2017-07-01/examples/ChildResources_List.json
 */
async function getHealthHistoryByResource() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri = "resourceUri";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.childResources.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getHealthHistoryByResource().catch(console.error);
