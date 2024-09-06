const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Reservations within a single ReservationOrder in the billing account.
 *
 * @summary List Reservations within a single ReservationOrder in the billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationsGetFromOrderByBillingAccount.json
 */
async function reservationsGetFromOrderByBillingAccount() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const reservationOrderId = "20000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.reservations.listByReservationOrder(
    billingAccountName,
    reservationOrderId,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
