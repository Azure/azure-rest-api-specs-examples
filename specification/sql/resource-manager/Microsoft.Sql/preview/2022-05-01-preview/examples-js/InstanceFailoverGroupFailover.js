const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fails over from the current primary managed instance to this managed instance.
 *
 * @summary Fails over from the current primary managed instance to this managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/InstanceFailoverGroupFailover.json
 */
async function plannedFailoverOfAFailoverGroup() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "Default";
  const locationName = "Japan West";
  const failoverGroupName = "failover-group-test-3";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.instanceFailoverGroups.beginFailoverAndWait(
    resourceGroupName,
    locationName,
    failoverGroupName,
  );
  console.log(result);
}
