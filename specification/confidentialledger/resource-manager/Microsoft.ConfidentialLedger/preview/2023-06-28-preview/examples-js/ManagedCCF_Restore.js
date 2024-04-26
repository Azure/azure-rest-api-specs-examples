const { ConfidentialLedgerClient } = require("@azure/arm-confidentialledger");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restores a Managed CCF Resource.
 *
 * @summary Restores a Managed CCF Resource.
 * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-06-28-preview/examples/ManagedCCF_Restore.json
 */
async function managedCcfRestore() {
  const subscriptionId =
    process.env["CONFIDENTIALLEDGER_SUBSCRIPTION_ID"] || "0000000-0000-0000-0000-000000000001";
  const resourceGroupName =
    process.env["CONFIDENTIALLEDGER_RESOURCE_GROUP"] || "DummyResourceGroupName";
  const appName = "DummyMccfAppName";
  const managedCCF = {
    fileShareName: "DummyFileShareName",
    restoreRegion: "EastUS",
    uri: "DummySASUri",
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfidentialLedgerClient(credential, subscriptionId);
  const result = await client.managedCCFOperations.beginRestoreAndWait(
    resourceGroupName,
    appName,
    managedCCF,
  );
  console.log(result);
}
