const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update reservation by billing account.
 *
 * @summary Update reservation by billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationUpdateByBillingAccount.json
 */
async function reservationUpdate() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const reservationOrderId = "20000000-0000-0000-0000-000000000000";
  const reservationId = "30000000-0000-0000-0000-000000000000";
  const body = { displayName: "NewName" };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.reservations.beginUpdateByBillingAccountAndWait(
    billingAccountName,
    reservationOrderId,
    reservationId,
    body,
  );
  console.log(result);
}
