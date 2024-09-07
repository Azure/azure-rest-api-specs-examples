const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a transfer request by ID. The caller must be the recipient of the transfer request.
 *
 * @summary Gets a transfer request by ID. The caller must be the recipient of the transfer request.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersGet.json
 */
async function recipientTransferGet() {
  const transferName = "aabb123";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.recipientTransfers.get(transferName);
  console.log(result);
}
