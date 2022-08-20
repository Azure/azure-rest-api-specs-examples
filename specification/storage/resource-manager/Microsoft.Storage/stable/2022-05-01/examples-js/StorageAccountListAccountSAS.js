const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List SAS credentials of a storage account.
 *
 * @summary List SAS credentials of a storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountListAccountSAS.json
 */
async function storageAccountListAccountSas() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7985";
  const accountName = "sto8588";
  const parameters = {
    keyToSign: "key1",
    sharedAccessExpiryTime: new Date("2017-05-24T11:42:03.1567373Z"),
    permissions: "r",
    protocols: "https,http",
    resourceTypes: "s",
    services: "b",
    sharedAccessStartTime: new Date("2017-05-24T10:42:03.1567373Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.listAccountSAS(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

storageAccountListAccountSas().catch(console.error);
