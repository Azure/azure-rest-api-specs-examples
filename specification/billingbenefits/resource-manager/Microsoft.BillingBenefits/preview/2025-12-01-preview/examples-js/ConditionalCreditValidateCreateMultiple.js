const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/ConditionalCreditValidateCreateMultiple.json
 */
async function conditionalCreditValidateCreateMultiple() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "ConditionalCredits",
        properties: {
          allowContributors: "Disabled",
          displayName: "Standard CACO",
          entityType: "Primary",
          milestones: [
            {
              name: "Milestone 1",
              award: { credit: { amount: 5000, currencyCode: "USD" }, duration: "P3M" },
              endAt: new Date("2026-03-31T00:00:00Z"),
              spendTarget: { amount: 50000, currencyCode: "USD" },
            },
          ],
          productCode: "000187f7-0000-0260-ab43-b8473ce57f1d",
          startAt: new Date("2025-12-01T00:00:00Z"),
        },
      },
      {
        benefitType: "ConditionalCredits",
        properties: {
          allowContributors: "Enabled",
          displayName: "High Value CACO",
          entityType: "Primary",
          milestones: [
            {
              name: "Milestone 1",
              award: { credit: { amount: 500000, currencyCode: "USD" }, duration: "P3M" },
              endAt: new Date("2026-03-31T00:00:00Z"),
              spendTarget: { amount: 2000000, currencyCode: "USD" },
            },
          ],
          productCode: "000187f7-0000-0260-ab43-b8473ce57f1d",
          startAt: new Date("2025-12-01T00:00:00Z"),
        },
      },
      {
        benefitType: "ConditionalCredits",
        properties: {
          allowContributors: "Disabled",
          displayName: "Special Product CACO",
          entityType: "Primary",
          milestones: [
            {
              name: "Milestone 1",
              award: { credit: { amount: 7500, currencyCode: "USD" }, duration: "P3M" },
              endAt: new Date("2026-03-31T00:00:00Z"),
              spendTarget: { amount: 75000, currencyCode: "USD" },
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
