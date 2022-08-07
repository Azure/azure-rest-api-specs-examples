const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a RouteTable.
 *
 * @summary Deletes a RouteTable.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/HubRouteTableDelete.json
 */
async function routeTableDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const routeTableName = "hubRouteTable1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.hubRouteTables.beginDeleteAndWait(
    resourceGroupName,
    virtualHubName,
    routeTableName
  );
  console.log(result);
}

routeTableDelete().catch(console.error);
