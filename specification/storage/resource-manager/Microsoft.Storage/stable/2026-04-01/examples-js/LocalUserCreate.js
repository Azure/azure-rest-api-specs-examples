const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the properties of a local user associated with the storage account. Properties for NFSv3 enablement and extended groups cannot be set with other properties.
 *
 * @summary create or update the properties of a local user associated with the storage account. Properties for NFSv3 enablement and extended groups cannot be set with other properties.
 * x-ms-original-file: 2026-04-01/LocalUserCreate.json
 */
async function createLocalUser() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.localUsers.createOrUpdate("res6977", "sto2527", "user1", {
    allowAclAuthorization: true,
    groupId: 2000,
    hasSshPassword: true,
    homeDirectory: "homedirectory",
    permissionScopes: [
      { permissions: "rwd", resourceName: "share1", service: "file" },
      { permissions: "rw", resourceName: "share2", service: "file" },
    ],
    sshAuthorizedKeys: [{ description: "key name", key: "ssh-rsa keykeykeykeykey=" }],
  });
  console.log(result);
}
