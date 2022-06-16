const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageSyncServicesCreate() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const parameters = {
    incomingTrafficPolicy: "AllowAllTraffic",
    location: "WestUS",
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.storageSyncServices.beginCreateAndWait(
    resourceGroupName,
    storageSyncServiceName,
    parameters
  );
  console.log(result);
}

storageSyncServicesCreate().catch(console.error);
