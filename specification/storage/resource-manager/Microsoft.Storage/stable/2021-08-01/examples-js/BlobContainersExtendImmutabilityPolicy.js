const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function extendImmutabilityPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6238";
  const accountName = "sto232";
  const containerName = "container5023";
  const ifMatch = '"8d59f830d0c3bf9"';
  const parameters = {
    immutabilityPeriodSinceCreationInDays: 100,
  };
  const options = { parameters: parameters };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.extendImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName,
    ifMatch,
    options
  );
  console.log(result);
}

extendImmutabilityPolicy().catch(console.error);
