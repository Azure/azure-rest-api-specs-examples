const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets an entity.
 *
 * @summary gets an entity.
 * x-ms-original-file: 2025-07-01-preview/entities/GetSecurityAlertEntityById.json
 */
async function getASecurityAlertEntity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.entities.get(
    "myRg",
    "myWorkspace",
    "4aa486e0-6f85-41af-99ea-7acdce7be6c8",
  );
  console.log(result);
}
