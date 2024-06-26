const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the all the children and its current health status for a parent resource. Use the nextLink property in the response to get the next page of children current health
 *
 * @summary Lists the all the children and its current health status for a parent resource. Use the nextLink property in the response to get the next page of children current health
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ChildResources_List.json
 */
async function getCurrentChildHealthByResource() {
  const resourceUri = "resourceUri";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const resArray = new Array();
  for await (let item of client.childResources.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
