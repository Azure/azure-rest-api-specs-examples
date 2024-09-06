const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the payment methods owned by the caller.
 *
 * @summary Lists the payment methods owned by the caller.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/paymentMethodsListByUser.json
 */
async function listPaymentMethodOwnedByUser() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.paymentMethods.listByUser()) {
    resArray.push(item);
  }
  console.log(resArray);
}
