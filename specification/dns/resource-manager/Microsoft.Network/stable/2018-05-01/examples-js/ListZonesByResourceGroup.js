const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the DNS zones within a resource group.
 *
 * @summary Lists the DNS zones within a resource group.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListZonesByResourceGroup.json
 */
async function listZonesByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.zones.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listZonesByResourceGroup().catch(console.error);
