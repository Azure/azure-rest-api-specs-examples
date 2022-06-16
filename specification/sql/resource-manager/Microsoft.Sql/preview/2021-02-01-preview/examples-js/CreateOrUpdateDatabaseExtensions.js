const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Perform a database extension operation, like polybase import
 *
 * @summary Perform a database extension operation, like polybase import
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/CreateOrUpdateDatabaseExtensions.json
 */
async function createOrUpdateDatabaseExtensions() {
  const subscriptionId = "a1c0814d-3c18-4e1e-a247-c128c12b1677";
  const resourceGroupName = "rg_20cbe0f0-c2d9-4522-9177-5469aad53029";
  const serverName = "srv_1ffd1cf8-9951-47fb-807d-a9c384763849";
  const databaseName = "878e303f-1ea0-4f17-aa3d-a22ac5e9da08";
  const extensionName = "polybaseimport";
  const parameters = {
    operationMode: "PolybaseImport",
    storageKey:
      "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    storageKeyType: "StorageAccessKey",
    storageUri: "https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.databaseExtensionsOperations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    extensionName,
    parameters
  );
  console.log(result);
}

createOrUpdateDatabaseExtensions().catch(console.error);
