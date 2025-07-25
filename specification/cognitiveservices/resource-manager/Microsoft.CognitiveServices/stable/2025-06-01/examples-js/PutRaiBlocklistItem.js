const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update the state of specified blocklist item associated with the Azure OpenAI account.
 *
 * @summary Update the state of specified blocklist item associated with the Azure OpenAI account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/PutRaiBlocklistItem.json
 */
async function putRaiBlocklistItem() {
  const subscriptionId =
    process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["COGNITIVESERVICES_RESOURCE_GROUP"] || "resourceGroupName";
  const accountName = "accountName";
  const raiBlocklistName = "raiBlocklistName";
  const raiBlocklistItemName = "raiBlocklistItemName";
  const raiBlocklistItem = {
    properties: { isRegex: false, pattern: "Pattern To Block" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.raiBlocklistItems.createOrUpdate(
    resourceGroupName,
    accountName,
    raiBlocklistName,
    raiBlocklistItemName,
    raiBlocklistItem,
  );
  console.log(result);
}
