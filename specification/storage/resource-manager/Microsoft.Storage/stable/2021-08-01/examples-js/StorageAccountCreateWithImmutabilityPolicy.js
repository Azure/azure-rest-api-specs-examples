const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountCreateWithImmutabilityPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9101";
  const accountName = "sto4445";
  const parameters = {
    extendedLocation: { name: "losangeles001", type: "EdgeZone" },
    immutableStorageWithVersioning: {
      enabled: true,
      immutabilityPolicy: {
        allowProtectedAppendWrites: true,
        immutabilityPeriodSinceCreationInDays: 15,
        state: "Unlocked",
      },
    },
    kind: "Storage",
    location: "eastus",
    sku: { name: "Standard_GRS" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginCreateAndWait(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

storageAccountCreateWithImmutabilityPolicy().catch(console.error);
