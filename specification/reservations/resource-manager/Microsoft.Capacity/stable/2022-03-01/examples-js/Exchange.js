const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns one or more `Reservations` in exchange for one or more `Reservation` purchases.

 *
 * @summary Returns one or more `Reservations` in exchange for one or more `Reservation` purchases.

 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/Exchange.json
 */
async function exchange() {
  const body = {
    properties: { sessionId: "66e2ac8f-439e-4345-8235-6fef07608081" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const result = await client.exchange.beginPostAndWait(body);
  console.log(result);
}

exchange().catch(console.error);
