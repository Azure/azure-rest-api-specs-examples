const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers export of jobs and returns an OperationID to track.
 *
 * @summary Triggers export of jobs and returns an OperationID to track.
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/JobCRUD/TriggerExportJobs.json
 */
async function triggerExportJobs() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "SwaggerTestRg";
  const vaultName = "NetSDKTestRsVault";
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.exportJobs.beginTriggerAndWait(resourceGroupName, vaultName);
  console.log(result);
}
