const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
 *
 * @summary Creates a VirtualHubRouteTableV2 resource if it doesn't exist else updates the existing VirtualHubRouteTableV2.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VirtualHubRouteTableV2Put.json
 */
async function virtualHubRouteTableV2Put() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualHubName = "virtualHub1";
  const routeTableName = "virtualHubRouteTable1a";
  const virtualHubRouteTableV2Parameters = {
    attachedConnections: ["All_Vnets"],
    routes: [
      {
        destinationType: "CIDR",
        destinations: ["20.10.0.0/16", "20.20.0.0/16"],
        nextHopType: "IPAddress",
        nextHops: ["10.0.0.68"],
      },
      {
        destinationType: "CIDR",
        destinations: ["0.0.0.0/0"],
        nextHopType: "IPAddress",
        nextHops: ["10.0.0.68"],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubRouteTableV2S.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualHubName,
    routeTableName,
    virtualHubRouteTableV2Parameters,
  );
  console.log(result);
}
