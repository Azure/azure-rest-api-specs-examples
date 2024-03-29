const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets Azure services' emerging issues.
 *
 * @summary Gets Azure services' emerging issues.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/EmergingIssues_Get.json
 */
async function getEmergingIssues() {
  const subscriptionId =
    process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const issueName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const result = await client.emergingIssues.get(issueName);
  console.log(result);
}
