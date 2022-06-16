const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified application security group.
 *
 * @summary Deletes the specified application security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ApplicationSecurityGroupDelete.json
 */
async function deleteApplicationSecurityGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const applicationSecurityGroupName = "test-asg";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationSecurityGroups.beginDeleteAndWait(
    resourceGroupName,
    applicationSecurityGroupName
  );
  console.log(result);
}

deleteApplicationSecurityGroup().catch(console.error);
