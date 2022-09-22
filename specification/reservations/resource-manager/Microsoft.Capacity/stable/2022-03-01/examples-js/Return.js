const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return a reservation.
 *
 * @summary Return a reservation.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/Return.json
 */
async function purchase() {
  const reservationOrderId = "276e7ae4-84d0-4da6-ab4b-d6b94f3557da";
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
  const result = await client.return.post(reservationOrderId, body);
  console.log(result);
}

purchase().catch(console.error);
