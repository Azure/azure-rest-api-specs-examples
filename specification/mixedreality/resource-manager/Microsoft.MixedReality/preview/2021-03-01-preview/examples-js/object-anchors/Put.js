const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creating or Updating an object anchors Account.
 *
 * @summary Creating or Updating an object anchors Account.
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/Put.json
 */
async function createObjectAnchorsAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const objectAnchorsAccount = {
    identity: { type: "SystemAssigned" },
    location: "eastus2euap",
  };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.objectAnchorsAccounts.create(
    resourceGroupName,
    accountName,
    objectAnchorsAccount
  );
  console.log(result);
}

createObjectAnchorsAccount().catch(console.error);
