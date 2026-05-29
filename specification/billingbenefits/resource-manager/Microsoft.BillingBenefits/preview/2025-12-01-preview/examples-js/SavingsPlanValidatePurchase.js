const { BillingBenefitsRP } = require("@azure/arm-billingbenefits");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to validate savings plan purchase.
 *
 * @summary validate savings plan purchase.
 * x-ms-original-file: 2025-12-01-preview/SavingsPlanValidatePurchase.json
 */
async function savingsPlanValidatePurchase() {
  const credential = new DefaultAzureCredential();
  const client = new BillingBenefitsRP(credential);
  const result = await client.benefit.validate({
    benefits: [
      {
        benefitType: "SavingsPlan",
        appliedScopeProperties: {
          resourceGroupId:
            "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg",
        },
        appliedScopeType: "Single",
        billingScopeId: "/subscriptions/10000000-0000-0000-0000-000000000000",
        commitment: { amount: 15.23, currencyCode: "USD", grain: "Hourly" },
        displayName: "ComputeSavingsPlan_2021-07-01",
        term: "P1Y",
        sku: { name: "Compute_Savings_Plan" },
      },
      {
        benefitType: "SavingsPlan",
        appliedScopeProperties: {
          resourceGroupId: "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/RG",
        },
        appliedScopeType: "Single",
        billingScopeId: "/subscriptions/10000000-0000-0000-0000-000000000000",
        commitment: { amount: 20, currencyCode: "USD", grain: "Hourly" },
        displayName: "ComputeSavingsPlan_2021-07-01",
        term: "P1Y",
        sku: { name: "Compute_Savings_Plan" },
      },
    ],
  });
  console.log(result);
}
