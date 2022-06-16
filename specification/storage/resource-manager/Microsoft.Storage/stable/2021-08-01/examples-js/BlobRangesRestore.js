const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function blobRangesRestore() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9101";
  const accountName = "sto4445";
  const parameters = {
    blobRanges: [
      { endRange: "container/blobpath2", startRange: "container/blobpath1" },
      { endRange: "", startRange: "container2/blobpath3" },
    ],
    timeToRestore: new Date("2019-04-20T15:30:00.0000000Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginRestoreBlobRangesAndWait(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

blobRangesRestore().catch(console.error);
