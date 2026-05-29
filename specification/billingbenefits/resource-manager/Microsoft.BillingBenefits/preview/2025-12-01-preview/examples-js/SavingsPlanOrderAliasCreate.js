const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a savings plan. Learn more about permissions needed at https://go.microsoft.com/fwlink/?linkid=2215851
 *
 * @summary create a savings plan. Learn more about permissions needed at https://go.microsoft.com/fwlink/?linkid=2215851
 * x-ms-original-file: 2025-12-01-preview/SavingsPlanOrderAliasCreate.json
 */
async function savingsPlanOrderAliasCreate() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.savingsPlanOrderAlias.create("spAlias123", {
    appliedScopeType: "Shared",
    billingPlan: "P1M",
    billingScopeId: "/subscriptions/30000000-0000-0000-0000-000000000000",
    commitment: { amount: 0.001, currencyCode: "USD", grain: "Hourly" },
    displayName: "Compute_SavingsPlan_10-28-2022_16-38",
    term: "P3Y",
    sku: { name: "Compute_Savings_Plan" },
  });
  console.log(result);
}
