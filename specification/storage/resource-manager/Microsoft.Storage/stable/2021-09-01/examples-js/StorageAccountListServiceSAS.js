const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List service SAS credentials of a specific resource.
 *
 * @summary List service SAS credentials of a specific resource.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountListServiceSAS.json
 */
async function storageAccountListServiceSas() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7439";
  const accountName = "sto1299";
  const parameters = {
    canonicalizedResource: "/blob/sto1299/music",
    sharedAccessExpiryTime: new Date("2017-05-24T11:32:48.8457197Z"),
    permissions: "l",
    resource: "c",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.listServiceSAS(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

storageAccountListServiceSas().catch(console.error);
