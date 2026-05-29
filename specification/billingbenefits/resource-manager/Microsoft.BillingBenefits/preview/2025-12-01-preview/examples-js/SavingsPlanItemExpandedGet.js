const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get savings plan.
 *
 * @summary get savings plan.
 * x-ms-original-file: 2025-12-01-preview/SavingsPlanItemExpandedGet.json
 */
async function savingsPlanItemWithExpandedRenewPropertiesGet() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.savingsPlan.get(
    "20000000-0000-0000-0000-000000000000",
    "30000000-0000-0000-0000-000000000000",
    { expand: "renewProperties" },
  );
  console.log(result);
}
