const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the integration account's Key Vault keys.
 *
 * @summary Gets the integration account's Key Vault keys.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_ListKeyVaultKeys.json
 */
async function getIntegrationAccountCallbackUrl() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const integrationAccountName = "testIntegrationAccount";
  const listKeyVaultKeys = {
    keyVault: {
      id: "subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.KeyVault/vaults/testKeyVault",
    },
    skipToken: "testSkipToken",
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.integrationAccounts.listKeyVaultKeys(
    resourceGroupName,
    integrationAccountName,
    listKeyVaultKeys
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
