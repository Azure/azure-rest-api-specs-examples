const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update storage resource.
 *
 * @summary Create or update storage resource.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Storages_CreateOrUpdate.json
 */
async function storagesCreateOrUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const storageName = "mystorage";
  const storageResource = {
    properties: {
      accountKey: "account-key-of-storage-account",
      accountName: "storage-account-name",
      storageType: "StorageAccount",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.storages.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    storageName,
    storageResource,
  );
  console.log(result);
}
