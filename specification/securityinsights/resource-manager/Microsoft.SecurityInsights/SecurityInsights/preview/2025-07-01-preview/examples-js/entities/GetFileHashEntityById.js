const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets an entity.
 *
 * @summary gets an entity.
 * x-ms-original-file: 2025-07-01-preview/entities/GetFileHashEntityById.json
 */
async function getAFileHashEntity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.entities.get(
    "myRg",
    "myWorkspace",
    "ea359fa6-c1e5-f878-e105-6344f3e399a1",
  );
  console.log(result);
}
