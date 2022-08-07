const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the available bgp service communities.
 *
 * @summary Gets all the available bgp service communities.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ServiceCommunityList.json
 */
async function serviceCommunityList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bgpServiceCommunities.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

serviceCommunityList().catch(console.error);
