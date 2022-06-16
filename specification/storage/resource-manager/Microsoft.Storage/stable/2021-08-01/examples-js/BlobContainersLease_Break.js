const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function breakALeaseOnAContainer() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const containerName = "container6185";
  const parameters = {
    action: "Break",
    breakPeriod: undefined,
    leaseDuration: undefined,
    leaseId: "8698f513-fa75-44a1-b8eb-30ba336af27d",
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

breakALeaseOnAContainer().catch(console.error);
