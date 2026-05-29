const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update MACC.
 *
 * @summary update MACC.
 * x-ms-original-file: 2025-12-01-preview/MaccMilestoneUnsuppressAutomaticShortfall.json
 */
async function maccMilestoneUnsuppressAutomaticShortfall() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.maccs.update("resource_group_name_01", "macc_20230614", {
    milestones: [
      { automaticShortfall: "Enabled", milestoneId: "11111111-1111-1111-1111-111111111111" },
    ],
  });
  console.log(result);
}
