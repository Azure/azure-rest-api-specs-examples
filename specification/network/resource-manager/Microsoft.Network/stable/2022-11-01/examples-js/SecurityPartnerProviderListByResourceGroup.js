const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Security Partner Providers in a resource group.
 *
 * @summary Lists all Security Partner Providers in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/SecurityPartnerProviderListByResourceGroup.json
 */
async function listAllSecurityPartnerProvidersForAGivenResourceGroup() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.securityPartnerProviders.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
