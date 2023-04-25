const { ConfidentialLedgerClient } = require("@azure/arm-confidentialledger");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the properties of a Confidential Ledger.
 *
 * @summary Retrieves the properties of a Confidential Ledger.
 * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/ConfidentialLedger_Get.json
 */
async function confidentialLedgerGet() {
  const subscriptionId =
    process.env["CONFIDENTIALLEDGER_SUBSCRIPTION_ID"] || "0000000-0000-0000-0000-000000000001";
  const resourceGroupName =
    process.env["CONFIDENTIALLEDGER_RESOURCE_GROUP"] || "DummyResourceGroupName";
  const ledgerName = "DummyLedgerName";
  const credential = new DefaultAzureCredential();
  const client = new ConfidentialLedgerClient(credential, subscriptionId);
  const result = await client.ledger.get(resourceGroupName, ledgerName);
  console.log(result);
}
