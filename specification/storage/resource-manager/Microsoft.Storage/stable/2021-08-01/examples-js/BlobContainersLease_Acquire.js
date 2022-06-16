const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function acquireALeaseOnAContainer() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const containerName = "container6185";
  const parameters = {
    action: "Acquire",
    breakPeriod: undefined,
    leaseDuration: -1,
    leaseId: undefined,
    proposedLeaseId: undefined,
  };
  const options = { parameters: parameters };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.lease(
    resourceGroupName,
    accountName,
    containerName,
    options
  );
  console.log(result);
}

acquireALeaseOnAContainer().catch(console.error);
