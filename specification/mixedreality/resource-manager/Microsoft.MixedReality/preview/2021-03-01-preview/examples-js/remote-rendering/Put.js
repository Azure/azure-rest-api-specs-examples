const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creating or Updating a Remote Rendering Account.
 *
 * @summary Creating or Updating a Remote Rendering Account.
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Put.json
 */
async function createRemoteRenderingAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const remoteRenderingAccount = {
    identity: { type: "SystemAssigned" },
    location: "eastus2euap",
  };
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.remoteRenderingAccounts.create(
    resourceGroupName,
    accountName,
    remoteRenderingAccount
  );
  console.log(result);
}

createRemoteRenderingAccount().catch(console.error);
