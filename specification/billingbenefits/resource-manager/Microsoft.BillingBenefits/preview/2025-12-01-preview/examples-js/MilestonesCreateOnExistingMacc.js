const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update MACC.
 *
 * @summary update MACC.
 * x-ms-original-file: 2025-12-01-preview/MilestonesCreateOnExistingMacc.json
 */
async function milestonesCreateOnExistingMacc() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.maccs.update("resource_group_name_01", "macc_20230614", {
    milestones: [
      {
        commitment: { amount: 10000, currencyCode: "USD" },
        endAt: new Date("2026-05-31T23:59:59Z"),
      },
      {
        commitment: { amount: 15000, currencyCode: "USD" },
        endAt: new Date("2027-05-31T23:59:59Z"),
      },
    ],
  });
  console.log(result);
}
