const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate specified Key of a Remote Rendering Account
 *
 * @summary Regenerate specified Key of a Remote Rendering Account
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/RegenerateKey.json
 */
async function regenerateRemoteRenderingAccountKeys() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const regenerate = { serial: 1 };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.remoteRenderingAccounts.regenerateKeys(
    resourceGroupName,
    accountName,
    regenerate
  );
  console.log(result);
}

regenerateRemoteRenderingAccountKeys().catch(console.error);
