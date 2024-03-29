const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Both of the 2 Keys of a Remote Rendering Account
 *
 * @summary List Both of the 2 Keys of a Remote Rendering Account
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/ListKeys.json
 */
async function listRemoteRenderingAccountKey() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.remoteRenderingAccounts.listKeys(resourceGroupName, accountName);
  console.log(result);
}

listRemoteRenderingAccountKey().catch(console.error);
