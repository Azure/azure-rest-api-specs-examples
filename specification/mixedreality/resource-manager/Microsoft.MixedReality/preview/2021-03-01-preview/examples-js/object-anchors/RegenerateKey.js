const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate specified Key of an object anchors Account
 *
 * @summary Regenerate specified Key of an object anchors Account
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/RegenerateKey.json
 */
async function regenerateObjectAnchorsAccountKeys() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const regenerate = { serial: 1 };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.objectAnchorsAccounts.regenerateKeys(
    resourceGroupName,
    accountName,
    regenerate
  );
  console.log(result);
}

regenerateObjectAnchorsAccountKeys().catch(console.error);
