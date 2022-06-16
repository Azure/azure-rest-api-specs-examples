const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables uploading ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
 *
 * @summary Disables uploading ledger digests to an Azure Storage account or an Azure Confidential Ledger instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/LedgerDigestUploadsDisable.json
 */
async function disablesUploadingLedgerDigestsForADatabase() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "ledgertestrg";
  const serverName = "ledgertestserver";
  const databaseName = "testdb";
  const ledgerDigestUploads = "current";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.ledgerDigestUploadsOperations.beginDisableAndWait(
    resourceGroupName,
    serverName,
    databaseName,
    ledgerDigestUploads
  );
  console.log(result);
}

disablesUploadingLedgerDigestsForADatabase().catch(console.error);
