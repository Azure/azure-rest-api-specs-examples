const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the specified account key for the specified Cognitive Services account.
 *
 * @summary Regenerates the specified account key for the specified Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/RegenerateKey.json
 */
async function regenerateKeys() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "myResourceGroup";
  const accountName = "myAccount";
  const keyName = "Key2";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.regenerateKey(resourceGroupName, accountName, keyName);
  console.log(result);
}

regenerateKeys().catch(console.error);
