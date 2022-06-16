const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get usages for the requested Cognitive Services account
 *
 * @summary Get usages for the requested Cognitive Services account
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/GetUsages.json
 */
async function getUsages() {
  const subscriptionId = "5a4f5c2e-6983-4ccb-bd34-2196d5b5bbd3";
  const resourceGroupName = "myResourceGroup";
  const accountName = "TestUsage02";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.listUsages(resourceGroupName, accountName);
  console.log(result);
}

getUsages().catch(console.error);
