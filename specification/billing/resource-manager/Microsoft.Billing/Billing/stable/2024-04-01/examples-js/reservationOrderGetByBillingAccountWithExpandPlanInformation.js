const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get the details of the ReservationOrder in the billing account.
 *
 * @summary get the details of the ReservationOrder in the billing account.
 * x-ms-original-file: 2024-04-01/reservationOrderGetByBillingAccountWithExpandPlanInformation.json
 */
async function reservationOrderGetByBillingAccountWithExpandPlanInformation() {
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.reservationOrders.getByBillingAccount(
    "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
    "20000000-0000-0000-0000-000000000000",
  );
  console.log(result);
}
