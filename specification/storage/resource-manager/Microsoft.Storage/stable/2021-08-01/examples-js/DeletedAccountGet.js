const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function deletedAccountGet() {
  const subscriptionId = "{subscription-id}";
  const deletedAccountName = "sto1125";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.deletedAccounts.get(deletedAccountName, location);
  console.log(result);
}

deletedAccountGet().catch(console.error);
