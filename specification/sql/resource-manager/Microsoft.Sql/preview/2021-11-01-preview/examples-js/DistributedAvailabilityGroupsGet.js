const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a distributed availability group info.
 *
 * @summary Gets a distributed availability group info.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DistributedAvailabilityGroupsGet.json
 */
async function getsTheDistributedAvailabilityGroupInfo() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "f2669dff-5f08-45dd-b857-b2a60b72cdc9";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "testrg";
  const managedInstanceName = "testcl";
  const distributedAvailabilityGroupName = "dag";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.distributedAvailabilityGroups.get(
    resourceGroupName,
    managedInstanceName,
    distributedAvailabilityGroupName,
  );
  console.log(result);
}
