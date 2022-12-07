const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Remote Rendering Accounts by Subscription
 *
 * @summary List Remote Rendering Accounts by Subscription
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/GetBySubscription.json
 */
async function listRemoteRenderingAccountsBySubscription() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.remoteRenderingAccounts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRemoteRenderingAccountsBySubscription().catch(console.error);
