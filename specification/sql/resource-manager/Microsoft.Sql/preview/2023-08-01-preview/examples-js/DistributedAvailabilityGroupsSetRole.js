const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the role for managed instance in a distributed availability group.
 *
 * @summary Sets the role for managed instance in a distributed availability group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2023-08-01-preview/examples/DistributedAvailabilityGroupsSetRole.json
 */
async function setDistributedAvailabilityGroupPrimaryReplicaToManagedInstance() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "testrg";
  const managedInstanceName = "testcl";
  const distributedAvailabilityGroupName = "dag";
  const parameters = {
    instanceRole: "Primary",
    roleChangeType: "Forced",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.distributedAvailabilityGroups.beginSetRoleAndWait(
    resourceGroupName,
    managedInstanceName,
    distributedAvailabilityGroupName,
    parameters,
  );
  console.log(result);
}
