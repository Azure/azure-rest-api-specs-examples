const { EdgeClient } = require("@azure/arm-sitemanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get Site at SG scope
 *
 * @summary get Site at SG scope
 * x-ms-original-file: 2025-03-01-preview/SitesByServiceGroup_Get_MaximumSet_Gen.json
 */
async function sitesByServiceGroupGetGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const client = new EdgeClient(credential, subscriptionId);
  const result = await client.sitesByServiceGroup.get("string", "string");
  console.log(result);
}
