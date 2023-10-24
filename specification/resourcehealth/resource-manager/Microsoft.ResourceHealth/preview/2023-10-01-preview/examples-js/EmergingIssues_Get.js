const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets Azure services' emerging issues.
 *
 * @summary Gets Azure services' emerging issues.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/EmergingIssues_Get.json
 */
async function getEmergingIssues() {
  const issueName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const result = await client.emergingIssues.get(issueName);
  console.log(result);
}
