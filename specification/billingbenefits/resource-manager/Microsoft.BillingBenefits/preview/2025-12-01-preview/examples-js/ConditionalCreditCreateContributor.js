const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a conditional credit.
 *
 * @summary create or update a conditional credit.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditCreateContributor.json
 */
async function conditionalCreditCreateContributor() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "20000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.conditionalCredits.createOrUpdate(
    "resource_group_name_02",
    "conditionalCredit_contributor_20250801",
    {
      location: "global",
      properties: {
        displayName: "Contributor Conditional Credit 20250801",
        entityType: "Contributor",
        primaryResourceId:
          "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/conditionalCredits/conditionalCredit_20250801",
        productCode: "000187f7-0000-0260-ab43-b8473ce57f1d",
        startAt: new Date("2025-09-01T00:00:00Z"),
      },
      tags: { environment: "dev", team: "finance" },
    },
  );
  console.log(result);
}
