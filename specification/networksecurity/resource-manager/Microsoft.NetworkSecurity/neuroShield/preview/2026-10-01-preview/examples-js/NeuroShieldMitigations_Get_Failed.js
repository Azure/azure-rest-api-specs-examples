const { NetworkSecurityManagementClient } = require("@azure/arm-networksecurity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a mitigation under a protected resource.
 *
 * @summary gets a mitigation under a protected resource.
 * x-ms-original-file: 2026-10-01-preview/NeuroShieldMitigations_Get_Failed.json
 */
async function getMitigationFailedStatus() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkSecurityManagementClient(credential, subscriptionId);
  const result = await client.neuroShieldMitigations.get(
    "myResourceGroup",
    "myNeuroShield",
    "vm-01",
    "mitigation-01",
  );
  console.log(result);
}
