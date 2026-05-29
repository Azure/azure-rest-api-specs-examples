const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a conditional credit.
 *
 * @summary update a conditional credit.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditUpdateContributor.json
 */
async function conditionalCreditUpdateContributor() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "20000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.conditionalCredits.update(
    "resource_group_name_02",
    "conditionalCreditContributor_20250801",
    {
      displayName: "Updated Contributor Conditional Credit 20250801",
      tags: { environment: "test", team: "finance" },
    },
  );
  console.log(result);
}
