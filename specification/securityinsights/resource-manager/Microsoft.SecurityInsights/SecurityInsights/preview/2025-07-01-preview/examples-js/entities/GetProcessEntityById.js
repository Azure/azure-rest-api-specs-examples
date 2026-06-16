const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets an entity.
 *
 * @summary gets an entity.
 * x-ms-original-file: 2025-07-01-preview/entities/GetProcessEntityById.json
 */
async function getAProcessEntity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.entities.get(
    "myRg",
    "myWorkspace",
    "7264685c-038c-42c6-948c-38e14ef1fb98",
  );
  console.log(result);
}
