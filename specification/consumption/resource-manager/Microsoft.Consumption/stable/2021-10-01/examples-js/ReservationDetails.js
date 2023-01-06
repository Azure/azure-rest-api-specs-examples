const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations details for provided date range.
 *
 * @summary Lists the reservations details for provided date range.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationDetails.json
 */
async function reservationDetails() {
  const subscriptionId =
    process.env["CONSUMPTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const reservationOrderId = "00000000-0000-0000-0000-000000000000";
  const filter = "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-12-05";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationsDetails.listByReservationOrder(
    reservationOrderId,
    filter
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
