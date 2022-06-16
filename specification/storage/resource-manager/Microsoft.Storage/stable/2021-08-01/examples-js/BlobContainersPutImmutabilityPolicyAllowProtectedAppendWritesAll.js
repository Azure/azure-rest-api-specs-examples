const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res1782";
  const accountName = "sto7069";
  const containerName = "container6397";
  const parameters = {
    allowProtectedAppendWritesAll: true,
    immutabilityPeriodSinceCreationInDays: 3,
  };
  const options = { parameters: parameters };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.createOrUpdateImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName,
    options
  );
  console.log(result);
}

createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll().catch(console.error);
