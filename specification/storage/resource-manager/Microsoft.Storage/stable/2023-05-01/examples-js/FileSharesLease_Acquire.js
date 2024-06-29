const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Lease Share operation establishes and manages a lock on a share for delete operations. The lock duration can be 15 to 60 seconds, or can be infinite.
 *
 * @summary The Lease Share operation establishes and manages a lock on a share for delete operations. The lock duration can be 15 to 60 seconds, or can be infinite.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/FileSharesLease_Acquire.json
 */
async function acquireALeaseOnAShare() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res3376";
  const accountName = "sto328";
  const shareName = "share124";
  const parameters = {
    action: "Acquire",
    breakPeriod: undefined,
    leaseDuration: -1,
    leaseId: undefined,
    proposedLeaseId: undefined,
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileShares.lease(resourceGroupName, accountName, shareName, options);
  console.log(result);
}
