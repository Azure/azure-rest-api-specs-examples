const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the details of a particular failover test performed on the gateway based on the test Guid
 *
 * @summary This operation retrieves the details of a particular failover test performed on the gateway based on the test Guid
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualNetworkGatewayGetFailoverSingleTestDetails.json
 */
async function virtualNetworkGatewayGetFailoverSingleTestDetails() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayName = "ergw";
  const peeringLocation = "Vancouver";
  const failoverTestId = "fe458ae8-d2ae-4520-a104-44bc233bde7e";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginGetFailoverSingleTestDetailsAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
    peeringLocation,
    failoverTestId,
  );
  console.log(result);
}
