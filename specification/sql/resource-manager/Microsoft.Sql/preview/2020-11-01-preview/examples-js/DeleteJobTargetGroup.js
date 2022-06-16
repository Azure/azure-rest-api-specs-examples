const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a target group.
 *
 * @summary Deletes a target group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeleteJobTargetGroup.json
 */
async function deleteATargetGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const targetGroupName = "targetGroup1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobTargetGroups.delete(
    resourceGroupName,
    serverName,
    jobAgentName,
    targetGroupName
  );
  console.log(result);
}

deleteATargetGroup().catch(console.error);
