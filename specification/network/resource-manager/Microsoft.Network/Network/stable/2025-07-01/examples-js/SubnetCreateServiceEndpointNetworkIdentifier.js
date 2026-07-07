const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a subnet in the specified virtual network.
 *
 * @summary creates or updates a subnet in the specified virtual network.
 * x-ms-original-file: 2025-07-01/SubnetCreateServiceEndpointNetworkIdentifier.json
 */
async function createSubnetWithServiceEndpointsWithNetworkIdentifier() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.subnets.createOrUpdate("subnet-test", "vnetname", "subnet1", {
    addressPrefix: "10.0.0.0/16",
    serviceEndpoints: [
      {
        networkIdentifier: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/subnet-test/providers/Microsoft.Network/publicIPAddresses/test-ip",
        },
        service: "Microsoft.Storage",
      },
    ],
  });
  console.log(result);
}
