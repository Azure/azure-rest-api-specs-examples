const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to perform a database extension operation, like database import, database export, or polybase import
 *
 * @summary perform a database extension operation, like database import, database export, or polybase import
 * x-ms-original-file: 2025-01-01/CreateOrUpdateDatabaseExtensions.json
 */
async function createOrUpdateDatabaseExtensions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "a1c0814d-3c18-4e1e-a247-c128c12b1677";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseExtensionsOperations.createOrUpdate(
    "rg_20cbe0f0-c2d9-4522-9177-5469aad53029",
    "srv_1ffd1cf8-9951-47fb-807d-a9c384763849",
    "878e303f-1ea0-4f17-aa3d-a22ac5e9da08",
    "polybaseimport",
    {
      operationMode: "PolybaseImport",
      storageKey:
        "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      storageKeyType: "StorageAccessKey",
      storageUri: "https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml",
    },
  );
  console.log(result);
}
