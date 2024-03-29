const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate specified Key of a Spatial Anchors Account
 *
 * @summary Regenerate specified Key of a Spatial Anchors Account
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/spatial-anchors/RegenerateKey.json
 */
async function regenerateSpatialAnchorsAccountKeys() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const regenerate = { serial: 1 };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.spatialAnchorsAccounts.regenerateKeys(
    resourceGroupName,
    accountName,
    regenerate
  );
  console.log(result);
}

regenerateSpatialAnchorsAccountKeys().catch(console.error);
