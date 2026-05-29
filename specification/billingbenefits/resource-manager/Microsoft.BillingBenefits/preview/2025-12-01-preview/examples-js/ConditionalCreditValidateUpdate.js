const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditValidateUpdate.json
 */
async function conditionalCreditValidateUpdate() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "ConditionalCredits",
        properties: {
          billingAccountResourceId: "/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}",
          displayName: "Validation Test Update CACO",
          entityType: "Primary",
          milestones: [
            {
              name: "Milestone 1",
              award: { credit: { amount: 5000, currencyCode: "USD" }, duration: "P3M" },
              endAt: new Date("2026-03-31T23:59:59Z"),
              milestoneId: "550e8400-e29b-41d4-a716-446655440001",
              spendTarget: { amount: 50000, currencyCode: "USD" },
            },
            {
              name: "Milestone 2",
              award: { credit: { amount: 5000, currencyCode: "USD" }, duration: "P3M" },
              endAt: new Date("2026-06-30T23:59:59Z"),
              milestoneId: "550e8400-e29b-41d4-a716-446655440002",
              spendTarget: { amount: 100000, currencyCode: "USD" },
            },
          ],
          resourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.BillingBenefits/conditionalCredits/CACO1",
          systemId: "13810867107109237",
        },
      },
    ],
  });
  console.log(result);
}
