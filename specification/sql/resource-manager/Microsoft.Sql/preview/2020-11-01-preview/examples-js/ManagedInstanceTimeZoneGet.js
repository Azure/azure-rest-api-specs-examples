const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a managed instance time zone.
 *
 * @summary Gets a managed instance time zone.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceTimeZoneGet.json
 */
async function getManagedInstanceTimeZone() {
  const subscriptionId = "37d5e605-6142-4d79-b564-28b6dbfeec0f";
  const locationName = "canadaeast";
  const timeZoneId = "Haiti Standard Time";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.timeZones.get(locationName, timeZoneId);
  console.log(result);
}

getManagedInstanceTimeZone().catch(console.error);
