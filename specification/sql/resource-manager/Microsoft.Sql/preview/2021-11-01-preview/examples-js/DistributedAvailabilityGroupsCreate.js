const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a distributed availability group between Sql On-Prem and Sql Managed Instance.
 *
 * @summary Creates a distributed availability group between Sql On-Prem and Sql Managed Instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/DistributedAvailabilityGroupsCreate.json
 */
async function createADistributedAvailabilityGroup() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "testrg";
  const managedInstanceName = "testcl";
  const distributedAvailabilityGroupName = "dag";
  const parameters = {
    primaryAvailabilityGroupName: "BoxLocalAg1",
    secondaryAvailabilityGroupName: "testcl",
    sourceEndpoint: "TCP://SERVER:7022",
    targetDatabase: "testdb",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.distributedAvailabilityGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    distributedAvailabilityGroupName,
    parameters,
  );
  console.log(result);
}
