const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Cognitive Services account from the resource group.
 *
 * @summary Deletes a Cognitive Services account from the resource group.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/PurgeDeletedAccount.json
 */
async function deleteAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const location = "westus";
  const resourceGroupName = "myResourceGroup";
  const accountName = "PropTest01";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.deletedAccounts.beginPurgeAndWait(
    location,
    resourceGroupName,
    accountName
  );
  console.log(result);
}

deleteAccount().catch(console.error);
