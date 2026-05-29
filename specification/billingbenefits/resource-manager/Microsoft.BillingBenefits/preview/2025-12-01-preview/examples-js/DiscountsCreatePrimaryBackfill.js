const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create discount.
 *
 * @summary create discount.
 * x-ms-original-file: 2025-12-01-preview/DiscountsCreatePrimaryBackfill.json
 */
async function discountsCreatePrimaryBackfill() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "30000000-0000-0000-0000-000000000000";
  const client = new BillingBenefitsRP(credential, subscriptionId);
  const result = await client.discounts.create("testrg", "testprimarydiscount", {
    location: "global",
    properties: {
      appliedScopeType: "BillingAccount",
      discountTypeProperties: {
        applyDiscountOn: "Purchase",
        conditions: [{ type: "equalAny", conditionName: "Cloud", value: ["US-Sec"] }],
        discountCombinationRule: "BestOf",
        discountPercentage: 14,
        discountType: "ProductFamily",
        productFamilyName: "Azure",
      },
      displayName: "Virtual Machines D Series",
      endAt: new Date("2024-07-01T23:59:59Z"),
      entityType: "Primary",
      productCode: "0001d726-0000-0160-330f-a0b98cdbbdc4",
      startAt: new Date("2023-07-01T00:00:00Z"),
      systemId: "13810867107109237",
    },
    tags: { key1: "value1", key2: "value2" },
  });
  console.log(result);
}
