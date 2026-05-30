const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a MACC.
 *
 * @summary get a MACC.
 * x-ms-original-file: 2025-12-01-preview/ContributorWithMilestonesAndShortfallGet.json
 */
async function contributorWithMilestonesAndShortfallGetFromPrimary() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.maccs.get("resource_group_name_02", "macc_contributor_20230614");
  console.log(result);
}
