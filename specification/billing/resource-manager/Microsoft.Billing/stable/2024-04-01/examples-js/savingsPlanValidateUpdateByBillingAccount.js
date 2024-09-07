const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validate savings plan patch by billing account.
 *
 * @summary Validate savings plan patch by billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlanValidateUpdateByBillingAccount.json
 */
async function savingsPlanValidateUpdate() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const savingsPlanOrderId = "20000000-0000-0000-0000-000000000000";
  const savingsPlanId = "30000000-0000-0000-0000-000000000000";
  const body = {
    benefits: [
      {
        appliedScopeProperties: {
          subscriptionId: "/subscriptions/50000000-0000-0000-0000-000000000000",
        },
        appliedScopeType: "Single",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.savingsPlans.validateUpdateByBillingAccount(
    billingAccountName,
    savingsPlanOrderId,
    savingsPlanId,
    body,
  );
  console.log(result);
}
