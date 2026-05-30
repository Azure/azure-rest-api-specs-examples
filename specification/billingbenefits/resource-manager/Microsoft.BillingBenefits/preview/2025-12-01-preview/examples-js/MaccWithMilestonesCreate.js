const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create MACC.
 *
 * @summary create MACC.
 * x-ms-original-file: 2025-12-01-preview/MaccWithMilestonesCreate.json
 */
async function maccWithMilestonesCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "10000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.maccs.create("resource_group_name_01", "macc_20230614", {
    location: "global",
    allowContributors: true,
    commitment: { amount: 20000, currencyCode: "USD", grain: "FullTerm" },
    displayName: "macc 20230614",
    endAt: new Date("2028-05-01T23:59:59Z"),
    entityType: "Primary",
    milestones: [
      {
        commitment: { amount: 10000, currencyCode: "USD" },
        endAt: new Date("2026-05-31T23:59:59Z"),
      },
      {
        commitment: { amount: 15000, currencyCode: "USD" },
        endAt: new Date("2027-05-31T23:59:59Z"),
      },
    ],
    productCode: "0001d726-0000-0160-330f-a0b98cdbbdc4",
    startAt: new Date("2025-05-01T00:00:00Z"),
    systemId: "13810867107109237",
    tags: { key1: "value1", key2: "value2" },
  });
  console.log(result);
}
