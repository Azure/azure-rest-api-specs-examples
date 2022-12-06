const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creating or Updating a Spatial Anchors Account.
 *
 * @summary Creating or Updating a Spatial Anchors Account.
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/spatial-anchors/Put.json
 */
async function createSpatialAnchorAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const spatialAnchorsAccount = {
    location: "eastus2euap",
  };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.spatialAnchorsAccounts.create(
    resourceGroupName,
    accountName,
    spatialAnchorsAccount
  );
  console.log(result);
}

createSpatialAnchorAccount().catch(console.error);
