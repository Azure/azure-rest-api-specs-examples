const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Available Scopes for `Reservation`.

 *
 * @summary Get Available Scopes for `Reservation`.

 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetAvailableScope.json
 */
async function availableScopes() {
  const reservationOrderId = "276e7ae4-84d0-4da6-ab4b-d6b94f3557da";
  const reservationId = "356e7ae4-84d0-4da6-ab4b-d6b94f3557da";
  const body = {
    properties: {
      scopes: ["/subscriptions/efc7c997-7700-4a74-b731-55aec16c15e9"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.reservation.beginAvailableScopesAndWait(
    reservationOrderId,
    reservationId,
    body
  );
  console.log(result);
}

availableScopes().catch(console.error);
