const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Declines a transfer request.
 *
 * @summary Declines a transfer request.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersDecline.json
 */
async function declineTransfer() {
  const transferName = "aabb123";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.recipientTransfers.decline(transferName);
  console.log(result);
}
