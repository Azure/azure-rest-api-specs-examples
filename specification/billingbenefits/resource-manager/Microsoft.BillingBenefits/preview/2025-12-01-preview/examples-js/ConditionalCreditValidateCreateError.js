const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditValidateCreateError.json
 */
async function conditionalCreditValidateCreateError() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "ConditionalCredits",
        properties: {
          allowContributors: "Enabled",
          displayName: "Invalid CACO Test",
          entityType: "Primary",
          milestones: [
            {
              name: "Milestone 1",
              award: {
                credit: { amount: 100000, currencyCode: "USD" },
                endAt: new Date("2021-12-31T00:00:00Z"),
                startAt: new Date("2022-01-01T00:00:00Z"),
              },
              endAt: new Date("2022-12-31T00:00:00Z"),
              spendTarget: { amount: -1000, currencyCode: "INVALID" },
            },
          ],
          productCode: "b8473ce57f1d",
          startAt: new Date("2023-01-01T00:00:00Z"),
        },
      },
      {
        benefitType: "ConditionalCredits",
        properties: {
          displayName: "Invalid Contributor CACO",
          entityType: "Contributor",
          primaryBillingAccountResourceId:
            "/providers/Microsoft.Billing/billingAccounts/nonexistent",
          primaryResourceId:
            "/subscriptions/invalid/resourceGroups/invalid/providers/Microsoft.BillingBenefits/conditionalCredits/nonexistent",
          productCode: "000187f7-0000-0260-ab43-b8473ce57f1d",
          startAt: new Date("2023-01-01T00:00:00Z"),
        },
      },
    ],
  });
  console.log(result);
}
