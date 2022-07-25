const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Split a `Reservation` into two `Reservation`s with specified quantity distribution.
 *
 * @summary Split a `Reservation` into two `Reservation`s with specified quantity distribution.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/SplitReservation.json
 */
async function split() {
  const reservationOrderId = "276e7ae4-84d0-4da6-ab4b-d6b94f3557da";
  const body = {
    quantities: [1, 2],
    reservationId:
      "/providers/Microsoft.Capacity/reservationOrders/276e7ae4-84d0-4da6-ab4b-d6b94f3557da/reservations/bcae77cd-3119-4766-919f-b50d36c75c7a",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.reservation.beginSplitAndWait(reservationOrderId, body);
  console.log(result);
}

split().catch(console.error);
