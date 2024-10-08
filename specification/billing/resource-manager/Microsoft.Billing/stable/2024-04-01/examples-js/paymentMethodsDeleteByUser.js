const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a payment method owned by the caller.
 *
 * @summary Deletes a payment method owned by the caller.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsDeleteByUser.json
 */
async function deletePaymentMethodOwnedByUser() {
  const paymentMethodName = "ABCDABCDABC0";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.paymentMethods.deleteByUser(paymentMethodName);
  console.log(result);
}
