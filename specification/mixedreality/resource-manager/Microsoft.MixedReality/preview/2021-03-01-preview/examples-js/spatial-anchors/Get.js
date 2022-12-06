const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a Spatial Anchors Account.
 *
 * @summary Retrieve a Spatial Anchors Account.
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/spatial-anchors/Get.json
 */
async function getSpatialAnchorsAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.spatialAnchorsAccounts.get(resourceGroupName, accountName);
  console.log(result);
}

getSpatialAnchorsAccount().catch(console.error);
