const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Express Route Circuit Connection in the specified express route circuits.
 *
 * @summary Creates or updates a Express Route Circuit Connection in the specified express route circuits.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteCircuitConnectionCreate.json
 */
async function expressRouteCircuitConnectionCreate() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const circuitName = "ExpressRouteARMCircuitA";
  const peeringName = "AzurePrivatePeering";
  const connectionName = "circuitConnectionUSAUS";
  const expressRouteCircuitConnectionParameters = {
    addressPrefix: "10.0.0.0/29",
    authorizationKey: "946a1918-b7a2-4917-b43c-8c4cdaee006a",
    expressRouteCircuitPeering: {
      id: "/subscriptions/subid1/resourceGroups/dedharcktinit/providers/Microsoft.Network/expressRouteCircuits/dedharcktlocal/peerings/AzurePrivatePeering",
    },
    ipv6CircuitConnectionConfig: { addressPrefix: "aa:bb::/125" },
    peerExpressRouteCircuitPeering: {
      id: "/subscriptions/subid2/resourceGroups/dedharcktpeer/providers/Microsoft.Network/expressRouteCircuits/dedharcktremote/peerings/AzurePrivatePeering",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuitConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    circuitName,
    peeringName,
    connectionName,
    expressRouteCircuitConnectionParameters
  );
  console.log(result);
}

expressRouteCircuitConnectionCreate().catch(console.error);
