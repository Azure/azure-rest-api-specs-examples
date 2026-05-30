const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update MACC.
 *
 * @summary update MACC.
 * x-ms-original-file: 2025-12-01-preview/MaccAndMilestoneMultiPropertyUpdate.json
 */
async function maccAndMilestoneMultiPropertyUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.maccs.update("resource_group_name_01", "macc_20230614", {
    allowContributors: false,
    commitment: { amount: 50000, currencyCode: "USD", grain: "FullTerm" },
    displayName: "macc 20230614 updated with milestones",
    endAt: new Date("2029-10-31T23:59:59Z"),
    milestones: [
      {
        commitment: { amount: 25000, currencyCode: "USD" },
        endAt: new Date("2027-06-30T23:59:59Z"),
        milestoneId: "11111111-1111-1111-1111-111111111111",
      },
      {
        commitment: { amount: 35000, currencyCode: "USD" },
        endAt: new Date("2028-06-30T23:59:59Z"),
        milestoneId: "22222222-2222-2222-2222-222222222222",
      },
    ],
  });
  console.log(result);
}
