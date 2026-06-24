const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Kusto database script.
 *
 * @summary creates a Kusto database script.
 * x-ms-original-file: 2025-02-14/KustoScriptsCreateOrUpdateWithManagedIdentity.json
 */
async function kustoScriptsCreateOrUpdateWithManagedIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.scripts.createOrUpdate(
    "kustorptest",
    "kustoCluster",
    "KustoDatabase8",
    "kustoScript",
    {
      scriptUrl: "https://mysa.blob.core.windows.net/container/script.txt",
      forceUpdateTag: "2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe",
      continueOnErrors: true,
      scriptLevel: "Database",
      principalPermissionsAction: "RemovePermissionOnScriptCompletion",
      managedIdentityResourceId:
        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustoprtest/providers/Microsoft.ManagedIdentity/userAssignedIdentities/kustoMi1",
    },
  );
  console.log(result);
}
