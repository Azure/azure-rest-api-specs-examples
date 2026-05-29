const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/MaccValidateCreateNewMaccWithMilestone.json
 */
async function maccValidateCreateNewMaccWithMilestone() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "MACC",
        allowContributors: true,
        billingAccountResourceId: "/providers/Microsoft.Billing/billingAccounts/{acctId:orgId}",
        commitment: { amount: 20000, currencyCode: "USD", grain: "FullTerm" },
        endAt: new Date("2026-07-31T23:59:59Z"),
        entityType: "Primary",
        milestones: [
          {
            commitment: { amount: 5000, currencyCode: "USD" },
            endAt: new Date("2024-07-31T23:59:59Z"),
          },
          {
            commitment: { amount: 15000, currencyCode: "USD" },
            endAt: new Date("2025-07-31T23:59:59Z"),
          },
        ],
        productCode: "0001d726-0000-0160-330f-a0b98cdbbdc4",
        resourceId:
          "/subscriptions/{primaryCloudSubId}/resourceGroups/resource_group_name_01/providers/Microsoft.BillingBenefits/maccs/macc_20230614",
        startAt: new Date("2023-07-01T00:00:00Z"),
      },
    ],
  });
  console.log(result);
}
