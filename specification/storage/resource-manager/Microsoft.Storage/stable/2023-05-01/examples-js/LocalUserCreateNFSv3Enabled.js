const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the properties of a local user associated with the storage account. Properties for NFSv3 enablement and extended groups cannot be set with other properties.
 *
 * @summary Create or update the properties of a local user associated with the storage account. Properties for NFSv3 enablement and extended groups cannot be set with other properties.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/LocalUserCreateNFSv3Enabled.json
 */
async function createNfSv3EnabledLocalUser() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res6977";
  const accountName = "sto2527";
  const username = "user1";
  const properties = {
    extendedGroups: [1001, 1005, 2005],
    isNFSv3Enabled: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.localUsersOperations.createOrUpdate(
    resourceGroupName,
    accountName,
    username,
    properties,
  );
  console.log(result);
}
