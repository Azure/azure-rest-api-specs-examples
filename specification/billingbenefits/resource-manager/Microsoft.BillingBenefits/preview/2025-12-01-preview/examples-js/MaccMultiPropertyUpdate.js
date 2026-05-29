const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update MACC.
 *
 * @summary update MACC.
 * x-ms-original-file: 2025-12-01-preview/MaccMultiPropertyUpdate.json
 */
async function maccMultiPropertyUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.maccs.update("resource_group_name_01", "macc_20230614", {
    allowContributors: false,
    commitment: { amount: 45000, currencyCode: "USD", grain: "FullTerm" },
    displayName: "macc 20230614 updated",
    endAt: new Date("2024-10-01T00:00:00Z"),
  });
  console.log(result);
}
