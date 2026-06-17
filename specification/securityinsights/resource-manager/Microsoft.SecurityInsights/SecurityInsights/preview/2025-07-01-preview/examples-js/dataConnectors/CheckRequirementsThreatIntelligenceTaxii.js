const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get requirements state for a data connector type.
 *
 * @summary get requirements state for a data connector type.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CheckRequirementsThreatIntelligenceTaxii.json
 */
async function checkRequirementsForTITaxii() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectorsCheckRequirementsOperations.post(
    "myRg",
    "myWorkspace",
    { kind: "ThreatIntelligenceTaxii", tenantId: "2070ecc9-b4d5-4ae4-adaa-936fa1954fa8" },
  );
  console.log(result);
}
