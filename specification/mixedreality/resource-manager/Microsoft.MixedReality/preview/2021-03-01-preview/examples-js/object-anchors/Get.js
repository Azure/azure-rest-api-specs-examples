const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve an Object Anchors Account.
 *
 * @summary Retrieve an Object Anchors Account.
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/Get.json
 */
async function getObjectAnchorsAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "MyResourceGroup";
  const accountName = "MyAccount";
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const result = await client.objectAnchorsAccounts.get(resourceGroupName, accountName);
  console.log(result);
}

getObjectAnchorsAccount().catch(console.error);
