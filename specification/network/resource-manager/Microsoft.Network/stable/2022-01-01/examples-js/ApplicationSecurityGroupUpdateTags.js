const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an application security group's tags.
 *
 * @summary Updates an application security group's tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ApplicationSecurityGroupUpdateTags.json
 */
async function updateApplicationSecurityGroupTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const applicationSecurityGroupName = "test-asg";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationSecurityGroups.updateTags(
    resourceGroupName,
    applicationSecurityGroupName,
    parameters
  );
  console.log(result);
}

updateApplicationSecurityGroupTags().catch(console.error);
