const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return a reservation and get refund information.
 *
 * @summary Return a reservation and get refund information.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/Return.json
 */
async function returnAReservation() {
  const reservationOrderId = "50000000-aaaa-bbbb-cccc-100000000004";
  const body = {
    properties: {
      reservationToReturn: {
        quantity: 1,
        reservationId:
          "/providers/microsoft.capacity/reservationOrders/50000000-aaaa-bbbb-cccc-100000000004/reservations/40000000-aaaa-bbbb-cccc-100000000000",
      },
      returnReason: "PurchasedWrongProduct",
      scope: "Reservation",
      sessionId: "10000000-aaaa-bbbb-cccc-200000000000",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.return.beginPostAndWait(reservationOrderId, body);
  console.log(result);
}
