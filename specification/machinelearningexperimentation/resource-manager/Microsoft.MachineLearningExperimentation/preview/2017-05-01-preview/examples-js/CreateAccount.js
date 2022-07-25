const { MLTeamAccountManagementClient } = require("@azure/arm-machinelearningexperimentation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a team account with the specified parameters.
 *
 * @summary Creates or updates a team account with the specified parameters.
 * x-ms-original-file: specification/machinelearningexperimentation/resource-manager/Microsoft.MachineLearningExperimentation/preview/2017-05-01-preview/examples/CreateAccount.json
 */
async function accountCreate() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "accountcrud-1234";
  const accountName = "accountcrud5678";
  const parameters = {
    keyVaultId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.KeyVault/vaults/testkv",
    location: "East US",
    storageAccount: {
      accessKey: "key",
      storageAccountId:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/Microsoft.Storage/storageAccounts/testStorageAccount",
    },
    tags: { tagKey1: "TagValue1" },
    vsoAccountId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/accountcrud-1234/providers/microsoft.visualstudio/account/vsotest",
  };
  const credential = new DefaultAzureCredential();
  const client = new MLTeamAccountManagementClient(credential, subscriptionId);
  const result = await client.accounts.createOrUpdate(resourceGroupName, accountName, parameters);
  console.log(result);
}

accountCreate().catch(console.error);
