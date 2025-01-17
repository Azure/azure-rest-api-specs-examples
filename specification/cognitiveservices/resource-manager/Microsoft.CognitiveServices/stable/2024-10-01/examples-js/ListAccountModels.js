const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List available Models for the requested Cognitive Services account
 *
 * @summary List available Models for the requested Cognitive Services account
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListAccountModels.json
 */
async function listAccountModels() {
  const subscriptionId = process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["COGNITIVESERVICES_RESOURCE_GROUP"] || "resourceGroupName";
  const accountName = "accountName";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.accounts.listModels(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
