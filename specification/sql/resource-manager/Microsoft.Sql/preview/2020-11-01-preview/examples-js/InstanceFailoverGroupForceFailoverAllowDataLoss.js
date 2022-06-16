const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fails over from the current primary managed instance to this managed instance. This operation might result in data loss.
 *
 * @summary Fails over from the current primary managed instance to this managed instance. This operation might result in data loss.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/InstanceFailoverGroupForceFailoverAllowDataLoss.json
 */
async function forcedFailoverOfAFailoverGroupAllowingDataLoss() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const locationName = "Japan West";
  const failoverGroupName = "failover-group-test-3";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.instanceFailoverGroups.beginForceFailoverAllowDataLossAndWait(
    resourceGroupName,
    locationName,
    failoverGroupName
  );
  console.log(result);
}

forcedFailoverOfAFailoverGroupAllowingDataLoss().catch(console.error);
