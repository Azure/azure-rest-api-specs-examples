const { ConfidentialLedgerClient } = require("@azure/arm-confidentialledger");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Backs up a Confidential Ledger Resource.
 *
 * @summary Backs up a Confidential Ledger Resource.
 * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/examples/ConfidentialLedger_Backup.json
 */
async function confidentialLedgerBackup() {
  const subscriptionId =
    process.env["CONFIDENTIALLEDGER_SUBSCRIPTION_ID"] || "0000000-0000-0000-0000-000000000001";
  const resourceGroupName =
    process.env["CONFIDENTIALLEDGER_RESOURCE_GROUP"] || "DummyResourceGroupName";
  const ledgerName = "DummyLedgerName";
  const confidentialLedger = {
    restoreRegion: "EastUS",
    uri: "DummySASUri",
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfidentialLedgerClient(credential, subscriptionId);
  const result = await client.ledger.beginBackupAndWait(
    resourceGroupName,
    ledgerName,
    confidentialLedger,
  );
  console.log(result);
}
