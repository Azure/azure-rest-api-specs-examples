const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for developer to access intelligent APIs. It's also the resource type for billing.
 *
 * @summary Create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for developer to access intelligent APIs. It's also the resource type for billing.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/CreateAccountMin.json
 */
async function createAccountMin() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "myResourceGroup";
  const accountName = "testCreate1";
  const account = {
    identity: { type: "SystemAssigned" },
    kind: "CognitiveServices",
    location: "West US",
    properties: {},
    sku: { name: "S0" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginCreateAndWait(resourceGroupName, accountName, account);
  console.log(result);
}

createAccountMin().catch(console.error);
