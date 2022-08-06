const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the details of a Virtual Hub Ip configuration.
 *
 * @summary Retrieves the details of a Virtual Hub Ip configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualHubIpConfigurationGet.json
 */
async function virtualHubVirtualHubRouteTableV2Get() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "hub1";
  const ipConfigName = "ipconfig1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubIpConfiguration.get(
    resourceGroupName,
    virtualHubName,
    ipConfigName
  );
  console.log(result);
}

virtualHubVirtualHubRouteTableV2Get().catch(console.error);
