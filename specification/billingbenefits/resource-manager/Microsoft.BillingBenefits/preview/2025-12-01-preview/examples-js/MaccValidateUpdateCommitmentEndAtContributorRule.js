const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/MaccValidateUpdateCommitmentEndAtContributorRule.json
 */
async function maccValidateUpdateCommitmentEndAtContributorRule() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "MACC",
        allowContributors: true,
        billingAccountResourceId: "/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}",
        commitment: { amount: 20000, currencyCode: "USD", grain: "FullTerm" },
        endAt: new Date("2024-07-31T23:59:59Z"),
        entityType: "Primary",
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.BillingBenefits/benefits/myBenefit",
        systemId: "13810867107109237",
      },
    ],
  });
  console.log(result);
}
