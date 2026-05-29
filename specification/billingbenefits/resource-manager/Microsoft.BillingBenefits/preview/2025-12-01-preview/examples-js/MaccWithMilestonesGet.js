const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a MACC.
 *
 * @summary get a MACC.
 * x-ms-original-file: 2025-12-01-preview/MaccWithMilestonesGet.json
 */
async function maccWithMilestonesGet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.maccs.get("resource_group_name_01", "macc_20230614");
  console.log(result);
}
