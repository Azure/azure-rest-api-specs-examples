const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updating a Remote Rendering Account
 *
 * @summary Updating a Remote Rendering Account
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Patch.json
 */
async function updateRemoteRenderingAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const remoteRenderingAccount = {
    identity: { type: "SystemAssigned" },
    location: "eastus2euap",
    tags: { hero: "romeo", heroine: "juliet" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.remoteRenderingAccounts.update(
    resourceGroupName,
    accountName,
    remoteRenderingAccount
  );
  console.log(result);
}

updateRemoteRenderingAccount().catch(console.error);
