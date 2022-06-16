const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Lease Share operation establishes and manages a lock on a share for delete operations. The lock duration can be 15 to 60 seconds, or can be infinite.
 *
 * @summary The Lease Share operation establishes and manages a lock on a share for delete operations. The lock duration can be 15 to 60 seconds, or can be infinite.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/FileSharesLease_Break.json
 */
async function breakALeaseOnAShare() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const shareName = "share12";
  const parameters = {
    action: "Break",
    breakPeriod: undefined,
    leaseDuration: undefined,
    leaseId: "8698f513-fa75-44a1-b8eb-30ba336af27d",
    proposedLeaseId: undefined,
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.lease(resourceGroupName, accountName, shareName, options);
  console.log(result);
}

breakALeaseOnAShare().catch(console.error);
