const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/MaccValidateUpdate.json
 */
async function maccValidateUpdate() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "MACC",
        allowContributors: true,
        billingAccountResourceId: "/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}",
        commitment: { amount: 20000, currencyCode: "USD", grain: "FullTerm" },
        displayName: "Test MACC Benefit",
        endAt: new Date("2030-07-31T23:59:59Z"),
        entityType: "Primary",
        milestones: [
          {
            commitment: { amount: 10000, currencyCode: "USD" },
            endAt: new Date("2024-07-31T23:59:59Z"),
            milestoneId: "11111111-1111-1111-1111-111111111111",
          },
          {
            commitment: { amount: 15000, currencyCode: "USD" },
            endAt: new Date("2027-07-31T23:59:59Z"),
            milestoneId: "22222222-2222-2222-2222-222222222222",
          },
        ],
        resourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.BillingBenefits/benefits/myBenefit",
        systemId: "13810867107109237",
      },
    ],
  });
  console.log(result);
}
