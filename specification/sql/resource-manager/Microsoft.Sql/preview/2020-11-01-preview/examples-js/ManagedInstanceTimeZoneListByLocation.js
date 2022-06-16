const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of managed instance time zones by location.
 *
 * @summary Gets a list of managed instance time zones by location.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceTimeZoneListByLocation.json
 */
async function listManagedInstanceTimeZonesByLocation() {
  const subscriptionId = "37d5e605-6142-4d79-b564-28b6dbfeec0f";
  const locationName = "canadaeast";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.timeZones.listByLocation(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedInstanceTimeZonesByLocation().catch(console.error);
