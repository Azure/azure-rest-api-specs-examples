const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountEnableCmk() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9407";
  const accountName = "sto8596";
  const parameters = {
    encryption: {
      keySource: "Microsoft.Keyvault",
      keyVaultProperties: {
        keyName: "wrappingKey",
        keyVaultUri: "https://myvault8569.vault.azure.net",
        keyVersion: "",
      },
      services: {
        blob: { enabled: true, keyType: "Account" },
        file: { enabled: true, keyType: "Account" },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.update(resourceGroupName, accountName, parameters);
  console.log(result);
}

storageAccountEnableCmk().catch(console.error);
