const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function blobServicesPutLastAccessTimeBasedTracking() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4410";
  const accountName = "sto8607";
  const parameters = {
    lastAccessTimeTrackingPolicy: {
      name: "AccessTimeTracking",
      blobType: ["blockBlob"],
      enable: true,
      trackingGranularityInDays: 1,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobServices.setServiceProperties(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

blobServicesPutLastAccessTimeBasedTracking().catch(console.error);
