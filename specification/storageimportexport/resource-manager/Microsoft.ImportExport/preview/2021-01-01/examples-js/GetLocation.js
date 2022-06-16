const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

async function getLocations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const locationName = "West US";
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const result = await client.locations.get(locationName);
  console.log(result);
}

getLocations().catch(console.error);
