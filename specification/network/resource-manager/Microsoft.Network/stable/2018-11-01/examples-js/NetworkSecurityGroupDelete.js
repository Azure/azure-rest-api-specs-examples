const { NetworkManagementClient } = require("@azure/arm-network-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified network security group.
 *
 * @summary Deletes the specified network security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2018-11-01/examples/NetworkSecurityGroupDelete.json
 */
async function deleteNetworkSecurityGroup() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkSecurityGroupName = "testnsg";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityGroups.beginDeleteAndWait(
    resourceGroupName,
    networkSecurityGroupName
  );
  console.log(result);
}
