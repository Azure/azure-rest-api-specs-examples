const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the transfer requests received by the caller.
 *
 * @summary Lists the transfer requests received by the caller.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersList.json
 */
async function recipientTransfersList() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.recipientTransfers.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
