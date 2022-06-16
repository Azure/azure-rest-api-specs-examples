const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function versionLevelWormContainerMigration() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res1782";
  const accountName = "sto7069";
  const containerName = "container6397";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.beginObjectLevelWormAndWait(
    resourceGroupName,
    accountName,
    containerName
  );
  console.log(result);
}

versionLevelWormContainerMigration().catch(console.error);
