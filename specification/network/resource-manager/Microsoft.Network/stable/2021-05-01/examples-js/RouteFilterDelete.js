const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified route filter.
 *
 * @summary Deletes the specified route filter.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RouteFilterDelete.json
 */
async function routeFilterDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeFilterName = "filterName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeFilters.beginDeleteAndWait(resourceGroupName, routeFilterName);
  console.log(result);
}

routeFilterDelete().catch(console.error);
