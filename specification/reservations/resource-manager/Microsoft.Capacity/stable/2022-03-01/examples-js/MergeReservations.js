const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Merge the specified `Reservation`s into a new `Reservation`. The two `Reservation`s being merged must have same properties.
 *
 * @summary Merge the specified `Reservation`s into a new `Reservation`. The two `Reservation`s being merged must have same properties.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/MergeReservations.json
 */
async function merge() {
  const reservationOrderId = "276e7ae4-84d0-4da6-ab4b-d6b94f3557da";
  const body = {
    sources: [
      "/providers/Microsoft.Capacity/reservationOrders/c0565a8a-4491-4e77-b07b-5e6d66718e1c/reservations/cea04232-932e-47db-acb5-e29a945ecc73",
      "/providers/Microsoft.Capacity/reservationOrders/c0565a8a-4491-4e77-b07b-5e6d66718e1c/reservations/5bf54dc7-dacd-4f46-a16b-7b78f4a59799",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.reservation.beginMergeAndWait(reservationOrderId, body);
  console.log(result);
}

merge().catch(console.error);
