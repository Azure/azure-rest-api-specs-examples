const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to perform a database extension operation, like database import, database export, or polybase import
 *
 * @summary perform a database extension operation, like database import, database export, or polybase import
 * x-ms-original-file: 2025-01-01/ImportDatabaseUsingDatabaseExtensions.json
 */
async function importDatabaseUsingDatabaseExtension() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "17ca4d13-bf7d-4c33-a60e-b87a2820a325";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseExtensionsOperations.createOrUpdate(
    "rg_062866bf-c4f4-41f9-abf0-b59132ca7924",
    "srv_2d6be2d2-26c8-4930-8fb6-82a5e95e0e82",
    "db_2a47e946-e414-4c00-94ac-ed886bb78302",
    "Import",
    {
      administratorLogin: "login",
      administratorLoginPassword: "password",
      authenticationType: "Sql",
      operationMode: "Import",
      storageKey:
        "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      storageKeyType: "StorageAccessKey",
      storageUri: "https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml",
    },
  );
  console.log(result);
}
