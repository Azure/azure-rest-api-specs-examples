const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified network security group.
 *
 * @summary Deletes the specified network security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkSecurityGroupDelete.json
 */
async function deleteNetworkSecurityGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkSecurityGroupName = "testnsg";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityGroups.beginDeleteAndWait(
    resourceGroupName,
    networkSecurityGroupName
  );
  console.log(result);
}

deleteNetworkSecurityGroup().catch(console.error);
