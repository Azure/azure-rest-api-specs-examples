const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get capabilities at specified location in a given subscription.
 *
 * @summary Get capabilities at specified location in a given subscription.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/stable/2023-12-30/examples/CapabilitySetByLocation.json
 */
async function capabilityResult() {
  const subscriptionId =
    process.env["MYSQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const locationName = "WestUS";
  const capabilitySetName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.locationBasedCapabilitySet.get(locationName, capabilitySetName);
  console.log(result);
}
