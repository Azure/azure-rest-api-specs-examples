const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditValidateCreate.json
 */
async function conditionalCreditValidateCreate() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "ConditionalCredits",
        properties: {
          allowContributors: "Enabled",
          billingAccountResourceId:
            "/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111_2025-10-28",
          displayName: "Validation Test CACO",
          entityType: "Primary",
          milestones: [
            {
              name: "Milestone 1",
              award: { credit: { amount: 5000, currencyCode: "USD" }, duration: "P3M" },
              endAt: new Date("2026-03-31T23:59:59Z"),
              spendTarget: { amount: 50000, currencyCode: "USD" },
            },
            {
              name: "Milestone 2",
              award: { credit: { amount: 5000, currencyCode: "USD" }, duration: "P3M" },
              endAt: new Date("2026-06-30T23:59:59Z"),
              spendTarget: { amount: 100000, currencyCode: "USD" },
            },
          ],
          productCode: "000187f7-0000-0260-ab43-b8473ce57f1d",
          startAt: new Date("2025-12-01T00:00:00Z"),
        },
      },
    ],
  });
  console.log(result);
}
