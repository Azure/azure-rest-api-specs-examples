const { ConfidentialLedgerClient } = require("@azure/arm-confidentialledger");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the properties of all Confidential Ledgers.
 *
 * @summary Retrieves the properties of all Confidential Ledgers.
 * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/examples/ConfidentialLedger_List.json
 */
async function confidentialLedgerList() {
  const subscriptionId =
    process.env["CONFIDENTIALLEDGER_SUBSCRIPTION_ID"] || "0000000-0000-0000-0000-000000000001";
  const resourceGroupName =
    process.env["CONFIDENTIALLEDGER_RESOURCE_GROUP"] || "DummyResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new ConfidentialLedgerClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.ledger.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
