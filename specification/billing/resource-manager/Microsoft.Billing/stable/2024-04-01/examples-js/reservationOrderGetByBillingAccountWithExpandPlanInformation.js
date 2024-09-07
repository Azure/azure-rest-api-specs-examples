const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the details of the ReservationOrder in the billing account.
 *
 * @summary Get the details of the ReservationOrder in the billing account.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationOrderGetByBillingAccountWithExpandPlanInformation.json
 */
async function reservationOrderGetByBillingAccountWithExpandPlanInformation() {
  const billingAccountName =
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
  const reservationOrderId = "20000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.reservationOrders.getByBillingAccount(
    billingAccountName,
    reservationOrderId,
  );
  console.log(result);
}
