const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a payment method owned by the caller.
 *
 * @summary Gets a payment method owned by the caller.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsGetByUser.json
 */
async function getPaymentMethodOwnedByUser() {
  const paymentMethodName = "ABCDABCDABC0";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.paymentMethods.getByUser(paymentMethodName);
  console.log(result);
}
