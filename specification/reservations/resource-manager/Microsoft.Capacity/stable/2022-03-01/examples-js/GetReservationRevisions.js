const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of all the revisions for the `Reservation`.
 *
 * @summary List of all the revisions for the `Reservation`.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetReservationRevisions.json
 */
async function reservationRevisions() {
  const reservationId = "6ef59113-3482-40da-8d79-787f823e34bc";
  const reservationOrderId = "276e7ae4-84d0-4da6-ab4b-d6b94f3557da";
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const resArray = new Array();
  for await (let item of client.reservation.listRevisions(reservationId, reservationOrderId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationRevisions().catch(console.error);
