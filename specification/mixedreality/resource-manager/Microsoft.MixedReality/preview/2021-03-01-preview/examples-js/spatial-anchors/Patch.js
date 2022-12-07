const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updating a Spatial Anchors Account
 *
 * @summary Updating a Spatial Anchors Account
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/spatial-anchors/Patch.json
 */
async function updateSpatialAnchorsAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const spatialAnchorsAccount = {
    location: "eastus2euap",
    tags: { hero: "romeo", heroine: "juliet" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.spatialAnchorsAccounts.update(
    resourceGroupName,
    accountName,
    spatialAnchorsAccount
  );
  console.log(result);
}

updateSpatialAnchorsAccount().catch(console.error);
