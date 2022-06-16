const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations summaries for daily or monthly grain.
 *
 * @summary Lists the reservations summaries for daily or monthly grain.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesDaily.json
 */
async function reservationSummariesDaily() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const reservationOrderId = "00000000-0000-0000-0000-000000000000";
  const grain = "daily";
  const filter = "properties/usageDate ge 2017-10-01 AND properties/usageDate le 2017-11-20";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationsSummaries.listByReservationOrder(
    reservationOrderId,
    grain,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationSummariesDaily().catch(console.error);
