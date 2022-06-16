const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function acquireALeaseOnAShare() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const shareName = "share124";
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
  const result = await client.fileShares.lease(resourceGroupName, accountName, shareName, options);
  console.log(result);
}

acquireALeaseOnAShare().catch(console.error);
