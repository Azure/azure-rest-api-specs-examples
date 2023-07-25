const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for developer to access intelligent APIs. It's also the resource type for billing.
 *
 * @summary Create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for developer to access intelligent APIs. It's also the resource type for billing.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/CreateAccount.json
 */
async function createAccount() {
  const subscriptionId =
    process.env["COGNITIVESERVICES_SUBSCRIPTION_ID"] || "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = process.env["COGNITIVESERVICES_RESOURCE_GROUP"] || "myResourceGroup";
  const accountName = "testCreate1";
  const account = {
    identity: { type: "SystemAssigned" },
    kind: "Emotion",
    location: "West US",
    properties: {
      encryption: {
        keySource: "Microsoft.KeyVault",
        keyVaultProperties: {
          keyName: "KeyName",
          keyVaultUri: "https://pltfrmscrts-use-pc-dev.vault.azure.net/",
          keyVersion: "891CF236-D241-4738-9462-D506AF493DFA",
        },
      },
      userOwnedStorage: [
        {
          resourceId:
            "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount",
        },
      ],
    },
    sku: { name: "S0" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginCreateAndWait(resourceGroupName, accountName, account);
  console.log(result);
}
