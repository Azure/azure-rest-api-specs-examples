const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a subnet in the specified virtual network.
 *
 * @summary Creates or updates a subnet in the specified virtual network.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/SubnetCreateWithSharingScope.json
 */
async function createSubnetWithSharingScope() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "subnet-test";
  const virtualNetworkName = "vnetname";
  const subnetName = "subnet1";
  const subnetParameters = { addressPrefix: "10.0.0.0/16" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.subnets.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualNetworkName,
    subnetName,
    subnetParameters,
  );
  console.log(result);
}
