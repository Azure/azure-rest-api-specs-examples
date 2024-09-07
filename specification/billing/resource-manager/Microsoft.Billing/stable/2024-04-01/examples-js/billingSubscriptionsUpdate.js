const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of a billing subscription.
 *
 * @summary Updates the properties of a billing subscription.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsUpdate.json
 */
async function billingSubscriptionsUpdate() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const billingSubscriptionName = "11111111-1111-1111-1111-111111111111";
  const parameters = {
    consumptionCostCenter: "ABC1234",
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.billingSubscriptions.beginUpdateAndWait(
    billingAccountName,
    billingSubscriptionName,
    parameters,
  );
  console.log(result);
}
