const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function createLocalUser() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
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

createLocalUser().catch(console.error);
