const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get specific Reservation details in the billing account.
 *
 * @summary Get specific Reservation details in the billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationGetByBillingAccountSingleScope.json
 */
async function reservationGetByBillingAccountSingleScope() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const reservationOrderId = "20000000-0000-0000-0000-000000000000";
  const reservationId = "30000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.reservations.getByReservationOrder(
    billingAccountName,
    reservationOrderId,
    reservationId,
  );
  console.log(result);
}
