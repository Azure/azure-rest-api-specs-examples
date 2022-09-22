const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the reservations and the roll up counts of reservations group by provisioning states that the user has access to in the current tenant.
 *
 * @summary List the reservations and the roll up counts of reservations group by provisioning states that the user has access to in the current tenant.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/GetReservations.json
 */
async function catalog() {
  const filter = "(properties%2farchived+eq+false)";
  const orderby = "properties/displayName asc";
  const skiptoken = 50;
  const take = 1;
  const options = {
    filter,
    orderby,
    skiptoken,
    take,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const resArray = new Array();
  for await (let item of client.reservation.listAll(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

catalog().catch(console.error);
