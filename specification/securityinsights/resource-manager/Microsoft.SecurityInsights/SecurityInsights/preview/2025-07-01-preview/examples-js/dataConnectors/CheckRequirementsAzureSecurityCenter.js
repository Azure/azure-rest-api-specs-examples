const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get requirements state for a data connector type.
 *
 * @summary get requirements state for a data connector type.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CheckRequirementsAzureSecurityCenter.json
 */
async function checkRequirementsForASC() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectorsCheckRequirementsOperations.post(
    "myRg",
    "myWorkspace",
    { kind: "AzureSecurityCenter", subscriptionId: "c0688291-89d7-4bed-87a2-a7b1bff43f4c" },
  );
  console.log(result);
}
