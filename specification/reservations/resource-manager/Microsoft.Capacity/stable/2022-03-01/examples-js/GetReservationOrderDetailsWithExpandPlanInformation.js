const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the details of the `ReservationOrder`.
 *
 * @summary Get the details of the `ReservationOrder`.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetReservationOrderDetailsWithExpandPlanInformation.json
 */
async function getReservationWithExpandPayments() {
  const reservationOrderId = "a075419f-44cc-497f-b68a-14ee811d48b9";
  const expand = "schedule";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.reservationOrder.get(reservationOrderId, options);
  console.log(result);
}

getReservationWithExpandPayments().catch(console.error);
