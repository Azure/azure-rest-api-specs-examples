const { ConfidentialLedgerClient } = require("@azure/arm-confidentialledger");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Retrieves the properties of all Managed CCF.
 *
 * @summary Retrieves the properties of all Managed CCF.
 * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/examples/ManagedCCF_ListBySub.json
 */
async function managedCcfListBySub() {
  const subscriptionId =
    process.env["CONFIDENTIALLEDGER_SUBSCRIPTION_ID"] || "0000000-0000-0000-0000-000000000001";
  const credential = new DefaultAzureCredential();
  const client = new ConfidentialLedgerClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.managedCCFOperations.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
