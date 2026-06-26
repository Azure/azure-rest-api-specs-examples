const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to perform a database extension operation, like database import, database export, or polybase import
 *
 * @summary perform a database extension operation, like database import, database export, or polybase import
 * x-ms-original-file: 2025-01-01/ExportDatabaseUsingDatabaseExtensions.json
 */
async function exportDatabaseUsingDatabaseExtension() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0ca8cd24-0b47-4ad5-bc7e-d70e35c44adf";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseExtensionsOperations.createOrUpdate(
    "rg_d1ef9eae-044d-4710-ba59-b82e84ad3157",
    "srv_9243d320-ac4e-4f97-8e06-b1167dae5f4c",
    "db_7fe424c8-23cf-4ac3-bdc3-e21f424bdb68",
    "Export",
    {
      administratorLogin: "login",
      administratorLoginPassword: "password",
      authenticationType: "Sql",
      operationMode: "Export",
      storageKey:
        "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      storageKeyType: "StorageAccessKey",
      storageUri: "https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml",
    },
  );
  console.log(result);
}
