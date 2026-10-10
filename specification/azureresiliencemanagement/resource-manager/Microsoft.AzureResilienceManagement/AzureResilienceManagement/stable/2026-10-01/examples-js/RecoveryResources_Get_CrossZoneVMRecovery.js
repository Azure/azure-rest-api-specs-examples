const { AzureResilienceManagementClient } = require("@azure/arm-resiliencemanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a RecoveryResource
 *
 * @summary get a RecoveryResource
 * x-ms-original-file: 2026-10-01/RecoveryResources_Get_CrossZoneVMRecovery.json
 */
async function recoveryResourcesGetCrossZoneVMRecovery() {
  const credential = new DefaultAzureCredential();
  const client = new AzureResilienceManagementClient(credential);
  const result = await client.recoveryResources.get(
    "sampleServiceGroupName",
    "samplePlanName",
    "12345678-9012-3456-7890-123456789012",
  );
  console.log(result);
}
