const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists Azure services' emerging issues.
 *
 * @summary Lists Azure services' emerging issues.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/EmergingIssues_List.json
 */
async function getEmergingIssues() {
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const resArray = new Array();
  for await (let item of client.emergingIssues.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
