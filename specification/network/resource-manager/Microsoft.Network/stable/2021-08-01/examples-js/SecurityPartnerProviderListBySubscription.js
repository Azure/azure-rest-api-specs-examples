const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Security Partner Providers in a subscription.
 *
 * @summary Gets all the Security Partner Providers in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/SecurityPartnerProviderListBySubscription.json
 */
async function listAllSecurityPartnerProvidersForAGivenSubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.securityPartnerProviders.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllSecurityPartnerProvidersForAGivenSubscription().catch(console.error);
