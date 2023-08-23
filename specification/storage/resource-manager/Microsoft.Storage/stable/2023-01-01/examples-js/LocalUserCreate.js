const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the properties of a local user associated with the storage account
 *
 * @summary Create or update the properties of a local user associated with the storage account
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/LocalUserCreate.json
 */
async function createLocalUser() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res6977";
  const accountName = "sto2527";
  const username = "user1";
  const properties = {
    hasSshPassword: true,
    homeDirectory: "homedirectory",
    permissionScopes: [
      { permissions: "rwd", resourceName: "share1", service: "file" },
      { permissions: "rw", resourceName: "share2", service: "file" },
    ],
    sshAuthorizedKeys: [{ description: "key name", key: "ssh-rsa keykeykeykeykey=" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.localUsersOperations.createOrUpdate(
    resourceGroupName,
    accountName,
    username,
    properties
  );
  console.log(result);
}
