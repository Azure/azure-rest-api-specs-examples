const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of all the `ReservationOrder`s that the user has access to in the current tenant.
 *
 * @summary List of all the `ReservationOrder`s that the user has access to in the current tenant.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetReservationOrders.json
 */
async function reservationOrderList() {
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const resArray = new Array();
  for await (let item of client.reservationOrder.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationOrderList().catch(console.error);
