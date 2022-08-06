const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an application security group.
 *
 * @summary Creates or updates an application security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ApplicationSecurityGroupCreate.json
 */
async function createApplicationSecurityGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const applicationSecurityGroupName = "test-asg";
  const parameters = { location: "westus" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationSecurityGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    applicationSecurityGroupName,
    parameters
  );
  console.log(result);
}

createApplicationSecurityGroup().catch(console.error);
